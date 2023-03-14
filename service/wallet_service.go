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
		//无助记词 自动生成
		mnemonic, err = util.NewMnemonic()
		if err != nil {
			log.Println(err)
			return models.HDWalletResp{}, constant.HD_WALLET_GENERATE_ERROR
		}
	}
	//从助记词生成种子
	seed, err := util.NewFromSeed(mnemonic)
	if err != nil {
		log.Println(err)
		return models.HDWalletResp{}, constant.HD_WALLET_GENERATE_ERROR
	}

	path := util.NewPath(req.Path)
	if path == nil {
		// 指定路径错误
		return models.HDWalletResp{}, constant.HD_WALLET_GENERATE_ERROR
	}

	pubKeyS, address, err := util.GenerateHDWalletSegWitAddress(seed, path)
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
