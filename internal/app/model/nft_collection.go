package model

import (
	"time"
)

type NftMODEL struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"id"` // 主键
	CreatedAt time.Time `json:"-"`                                                         // 创建时间
	UpdatedAt time.Time `json:"-"`                                                         // 更新时间
}

type NftCollection struct {
	NftMODEL
	Chain          string `gorm:"column:chain;index:chain_address_contract_token,unique" json:"chain" form:"chain"`                                             // 区块链的简称（eth, bnb, polygon, moonbeam, arbitrum, optimism, platon, avalanche, cronos）
	AccountAddress string `gorm:"column:account_address;type:char(44);index:chain_address_contract_token,unique" json:"account_address" form:"account_address"` // 资产持有者的地址
	Status         uint8  `gorm:"default:0;" json:"status" form:"status"`                                                                                       // 显示状态(0:初始状态 1:隐藏 2:显示)
	NFTScanOwn
}

type NFTScanOwn struct {
	ContractAddress     string  `gorm:"index:chain_address_contract_token,unique" json:"contract_address" form:"contract_address"` // 合约地址
	ContractName        string  `json:"contract_name" form:"contract_name"`                                                        // 合约名称
	ContractTokenID     string  `json:"contract_token_id" form:"contract_token_id"`
	TokenID             string  `gorm:"index:chain_address_contract_token,unique" json:"token_id" form:"token_id"`
	ErcType             string  `gorm:"column:erc_type" json:"erc_type" form:"erc_type"` // NFT 的 erc 标准类型（erc721 或 erc1155）
	Amount              string  `gorm:"column:amount" json:"amount" form:"amount"`       // 持有数量
	Minter              string  `json:"minter" form:"minter"`
	Owner               string  `json:"owner" form:"owner"`
	OwnTimestamp        int64   `json:"own_timestamp"`
	MintTimestamp       int64   `json:"mint_timestamp"`
	MintTransactionHash string  `json:"mint_transaction_hash"`
	MintPrice           float64 `json:"mint_price"`
	TokenURI            string  `json:"token_uri"`
	MetadataJSON        string  `json:"metadata_json"`
	Name                string  `json:"name"`
	ContentType         string  `json:"content_type"`
	ContentURI          string  `json:"-"`
	ImageURI            string  `json:"image_uri"`
	ExternalLink        string  `gorm:"column:external_link" json:"external_link" form:"external_link"` // NFT对应的网站链接
	//LatestTradePrice     interface{}   `json:"latest_trade_price"`
	//LatestTradeSymbol    interface{}   `json:"latest_trade_symbol"`
	//LatestTradeTimestamp interface{}   `json:"latest_trade_timestamp"`
	NftscanID  string `json:"nftscan_id"`
	NftscanURI string `json:"nftscan_uri"`
	//Attributes           []interface{} `json:"attributes"`
	RarityScore float64 `json:"rarity_score"`
	RarityRank  int     `json:"rarity_rank"`
}
