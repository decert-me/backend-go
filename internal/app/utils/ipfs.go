package utils

// GetDataFromCid
// @description: 获取IPFS内容
// @param: cid string
// @return: string, error
func GetDataFromCid(cid string) (string, error) {
	baseURL := "https://ipfs.learnblockchain.cn/"
	url := baseURL + cid
	client := GetReqClient()
	req, err := client.R().Get(url)
	return req.String(), err
}
