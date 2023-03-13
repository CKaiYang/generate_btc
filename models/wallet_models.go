package models

type HDWalletReq struct {
	Mnemonic string `json:"mnemonic"`
	Index    uint32 `json:"index"`
}

type HDWalletResp struct {
	PublicKey string `json:"public_key"`
	Address   string `json:"address"`
	Mnemonic  string `json:"mnemonic"`
}

type MulAddressReq struct {
	PublicKeyList []string `json:"public_key_list"`
	N             int      `json:"n"`
}

type MulAddressResp struct {
	Address string `json:"address"`
}
