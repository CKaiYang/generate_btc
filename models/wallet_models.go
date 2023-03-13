package models

type HDWalletReq struct {
	// 助记词
	Mnemonic string `json:"mnemonic"`
	// 指定层级
	Index uint32 `json:"index"`
}

type HDWalletResp struct {
	// 公钥
	PublicKey string `json:"public_key"`
	// 比特币地址
	Address string `json:"address"`
	// 助记词
	Mnemonic string `json:"mnemonic"`
}

type MulAddressReq struct {
	// 公钥列表
	PublicKeyList []string `json:"public_key_list"`
	// m/n
	N int `json:"n"`
}

type MulAddressResp struct {
	// 地址
	Address string `json:"address"`
}
