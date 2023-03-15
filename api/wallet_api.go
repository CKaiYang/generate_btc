package api

import (
	"generate_btc/constant"
	"generate_btc/models"
	"generate_btc/service"
	"generate_btc/util"
)

type WalletApi struct {
	BaseApi
	Serv *service.WalletService
}

// PostNewHdWallet
// POST /wallet/new/hd/wallet
func (api WalletApi) PostNewHdWallet(req models.HDWalletReq) {
	resp, err := api.Serv.NewHDWallet(req)
	if err != nil {
		util.CommonResultFailure(api.Ctx, constant.HD_WALLET_GENERATE_ERROR_STATUS, err.Error())
		return
	}
	util.CommonResultSuccess(api.Ctx, &resp)
}

// PostGenerateMulAddress
// POST /wallet/generate/mul/address
func (api WalletApi) PostGenerateMulAddress(req models.MulAddressReq) {
	resp, err := api.Serv.MulAddress(req)
	if err != nil {
		util.CommonResultFailure(api.Ctx, constant.MUL_ADDRESS_GENERATE_ERROR_STATUS, err.Error())
		return
	}
	util.CommonResultSuccess(api.Ctx, &resp)
}
