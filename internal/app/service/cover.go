package service

import (
	"backend-go/internal/app/model/request"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image"
	"log"
	"net/http"
)

const px = 1024
const maskHeight = px * 0.23
const textHeight = 64
const textPad = 60
const textMaxWidth = px - (textPad * 2)

func mediaToImage(uploadJSONNFT request.UploadJSONNFT) {
	// 获取IPFS中的图片
	// 发起 HTTP 请求获取图片数据
	response, err := http.Get(uploadJSONNFT.Properties.Media)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// 检查 HTTP 响应状态码
	if response.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status code: %d", response.StatusCode)
	}
	// 解码图片数据
	img, _, err := image.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 画布大小
	width := 1024
	height := 1024
	// 缩放图片
	dstImage := imaging.Resize(img, width, height, imaging.Lanczos)
	// 创建画布
	dc := gg.NewContext(width, height)
	// 加载字体
	if err := dc.LoadFontFace("msyh", 64); err != nil {
		fmt.Println("Error loading font:", err)
		return
	}
}
