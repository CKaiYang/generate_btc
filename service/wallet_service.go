package service

import (
	"generate_btc/constant"
	"generate_btc/models"
	"generate_btc/util"
	"log"
)

type WalletService struct {
}

func (serv WalletService) NewHDWallet(req models.HDWalletReq) (resp models.HDWalletResp, err error) {
	mnemonic := req.Mnemonic
	if mnemonic == "" || len(mnemonic) <= 0 {
		mnemonic, err = util.NewMnemonic()
		if err != nil {
			log.Println(err)
			return models.HDWalletResp{}, constant.HD_WALLET_GENERATE_ERROR
		}
	}
	seed := util.NewFromSeed(mnemonic)
	pubKeyS, address, err := util.GenerateHDWalletSegWitAddress(seed, req.Index)
	if err != nil {
		log.Println(err)
		return models.HDWalletResp{}, constant.HD_WALLET_GENERATE_ERROR
	}
	return models.HDWalletResp{
		PublicKey: pubKeyS,
		Address:   address,
		Mnemonic:  mnemonic,
	}, nil
}

func (serv WalletService) MulAddress(req models.MulAddressReq) (resp models.MulAddressResp, err error) {
	mulAddress, err := util.GenerateMulAddress(req.N, req.PublicKeyList...)
	if err != nil {
		log.Println(err)
		return models.MulAddressResp{}, constant.MUL_ADDRESS_GENERATE_ERROR
	}
	return models.MulAddressResp{
		Address: mulAddress,
	}, nil
}
