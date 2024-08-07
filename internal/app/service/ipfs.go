package service

import (
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type ipfsRPC struct {
	index int
	lock  atomic.Bool
}

var ipfsPoint ipfsRPC

func (s *Service) GetIPFSUploadAPI() string {
	return s.c.IPFS[ipfsPoint.index].UploadAPI
}

func (s *Service) GetIPFSGateway(hash string) string {
	if hash == "" {
		return s.c.IPFS[ipfsPoint.index].API
	}
	url := fmt.Sprintf("%s/%s", s.c.IPFS[ipfsPoint.index].API, hash)
	return url
}

func (s *Service) BalanceIPFS() {
	if ipfsPoint.lock.Load() {
		return
	}
	ipfsPoint.lock.Store(true)
	defer ipfsPoint.lock.Store(false)

	IPFS := s.c.IPFS
	indexList := make([]int64, len(IPFS))
	for i, v := range IPFS {
		if v.API == "" || v.UploadAPI == "" {
			return
		}
		spent, err := ipfsRequest(v.API, v.UploadAPI)
		if err != nil {
			fmt.Println(err)
		}
		indexList[i] = spent
		time.Sleep(time.Second * 1)
	}
	ipfsPoint.index, _ = utils.SliceMin[int64](indexList)
	log.Warnv("IPFS 切换: " + strconv.Itoa(ipfsPoint.index))
}

func ipfsRequest(api string, uploadAPI string) (spent int64, err error) {
	max := int64(9999999999999)
	defer func() {
		if err := recover(); err != nil {
			spent = max
			return
		}
	}()
	client := req.C().SetTimeout(15 * time.Second)
	startTime := time.Now()
	// 上传JSON
	// 组成请求体
	jsonReq := make(map[string]interface{})
	jsonReq["body"] = "{\"foo\":\"bar\"}"
	// 发送请求
	url := fmt.Sprintf("%s/upload/json", uploadAPI)
	res, err := client.R().SetBody(jsonReq).Post(url)
	if err != nil {
		return max, err
	}
	// 解析返回结果
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Hash    string `gorm:"column:hash" json:"hash" form:"hash"`
	}
	var resJson Response
	err = json.Unmarshal(res.Bytes(), &resJson)
	if err != nil {
		return max, err
	}
	if resJson.Status != "1" {
		return max, err
	}
	// 请求JSON
	urlReq := fmt.Sprintf("%s/%s", api, resJson.Hash)
	content, err := client.R().Get(urlReq)
	if err != nil || !gjson.Valid(content.String()) {
		return max, err
	}
	return time.Since(startTime).Milliseconds(), nil
}

// IPFSUploadFile
// @description: 上传文件
// @param: header *multipart.FileHeader
// @return: err error, list interface{}, total int64
func (s *Service) IPFSUploadFile(header *multipart.FileHeader) (err error, hash string) {
	file, err := header.Open()
	if err != nil {
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return
	}
	// Convert the byte slice to an io.Reader
	reader := bytes.NewReader(content)
	// 发送请求
	url := fmt.Sprintf("%s/upload/image", s.GetIPFSUploadAPI())
	client := req.C().SetTimeout(120 * time.Second)
	res, err := client.R().SetFileUpload(req.FileUpload{
		ParamName: "file",
		FileName:  header.Filename,
		GetFileContent: func() (io.ReadCloser, error) {
			return io.NopCloser(reader), nil
		},
		ContentType: header.Header.Get("Content-Type"),
	}).Post(url)
	if err != nil {
		go s.BalanceIPFS()
		return err, hash
	}
	// 解析返回结果
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Hash    string `gorm:"column:hash" json:"hash" form:"hash"`
	}
	var resJson Response
	err = json.Unmarshal(res.Bytes(), &resJson)
	if err != nil {
		return err, hash
	}
	if resJson.Status != "1" {
		go s.BalanceIPFS()
		log.Errorv("upload file failed", zap.Error(err))
		return err, hash
	}
	// 保存 IPFS
	f, openError := header.Open() // 读取文件
	if openError != nil {
		log.Errorv("file.Open() Filed", zap.Error(err))
		return err, hash
	}
	// 目录
	director := s.c.Local.IPFS + "/"
	mkdirErr := os.MkdirAll(director, os.ModePerm)
	if mkdirErr != nil {
		log.Errorv("os.MkdirAll() Filed", zap.Error(err))
	}
	p := director + resJson.Hash
	out, createErr := os.Create(p)
	if createErr != nil {
		log.Errorv("os.Create() Filed", zap.Error(err))
		return err, hash
	}
	defer out.Close()             // 创建文件 defer 关闭
	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		log.Errorv("io.Copy() Filed", zap.Error(err))
		return err, hash
	}

	return err, resJson.Hash
}

// IPFSUploadJSON
// @description: 上传JSON
// @param: header *multipart.FileHeader
// @return: err error, list interface{}, total int64
func (s *Service) IPFSUploadJSON(uploadJSON interface{}) (err error, hash string) {
	data, err := json.Marshal(uploadJSON)
	if err != nil {
		return err, hash
	}
	// 组成请求体
	jsonReq := make(map[string]interface{})
	jsonReq["body"] = string(data)
	// 发送请求
	url := fmt.Sprintf("%s/upload/json", s.GetIPFSUploadAPI())
	client := req.C().SetTimeout(120 * time.Second)
	res, err := client.R().SetBody(jsonReq).Post(url)
	if err != nil {
		go s.BalanceIPFS()
		return err, hash
	}

	// 解析返回结果
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Hash    string `gorm:"column:hash" json:"hash" form:"hash"`
	}
	var resJson Response
	err = json.Unmarshal(res.Bytes(), &resJson)
	if err != nil {
		return err, hash
	}
	if resJson.Status != "1" {
		go s.BalanceIPFS()
		return err, hash
	}
	director := s.c.Local.IPFS + "/"
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		log.Errorv("os.MkdirAll() Filed", zap.Error(err))
		return err, hash
	}
	p := director + resJson.Hash
	err = os.WriteFile(p, []byte(data), 0664)
	if err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return err, hash
	}
	return err, resJson.Hash
}

func getCIDFromIPFSURL(ipfsURL string) (string, error) {
	u, err := url.Parse(ipfsURL)
	if err != nil {
		return strings.TrimPrefix(ipfsURL, "ipfs://"), err
	}
	if u.Path == "" {
		return strings.TrimPrefix(ipfsURL, "ipfs://"), nil
	}
	return strings.TrimPrefix(u.Path, "/"), nil
}

// GetDataFromCid
// @description: 获取IPFS内容
// @param: cid string
// @return: string, error
func (s *Service) GetDataFromCid(cid string) (result string, err error) {
	// 去除前缀
	cid, _ = getCIDFromIPFSURL(cid)
	// 读取文件内容
	path := s.c.Local.IPFS
	filePath := path + "/" + cid
	// 判断文件是否存在
	if _, errExist := os.Stat(filePath); os.IsNotExist(errExist) {
		url := s.GetIPFSGateway(cid)
		client := req.C().SetTimeout(30 * time.Second)
		res, err := client.R().Get(url)
		if err != nil {
			return result, err
		}
		if !gjson.Valid(res.String()) {
			return res.String(), errors.New("invalid json")
		}

		director := path + "/"
		err = os.MkdirAll(director, os.ModePerm)
		if err != nil {
			log.Errorv("os.MkdirAll() Filed", zap.Error(err))
			return result, err
		}
		err = os.WriteFile(filePath, res.Bytes(), 0664)
		if err != nil {
			log.Errorv("os.WriteFile() Filed", zap.Error(err))
			return result, err
		}
		return res.String(), nil
	}
	data, err := os.ReadFile(path + "/" + cid)
	if err != nil {
		log.Errorv("读取文件出错: ", zap.Error(err))
		return
	}
	return string(data), err
}

func (s *Service) GetDataFromUri(uri string) (string, error) {
	return s.GetDataFromCid(strings.Replace(uri, "ipfs://", "", 1))
}
