package util

import (
	"encoding/hex"
	"generate_btc/constant"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

var net = &chaincfg.MainNetParams

// GenerateHDWalletSegWitAddress
// 按照指定路径生成HD钱包公钥已经隔离见证的比特币地址
// seed 种子
// index 层级
func GenerateHDWalletSegWitAddress(seed []byte, index uint32) (pubKeyS string, address string, err error) {
	master, err := hdkeychain.NewMaster(seed, net)
	if err != nil {
		return "", "", err
	}
	var pubKey *btcec.PublicKey

	if err != nil {
		return "", "", err
	}

	if 0 == index {
		// 母公钥
		pubKey, err = master.ECPubKey()
	} else {
		//子私钥 层级生成
		child, err := master.Child(index)
		if err != nil {
			return "", "", err
		}
		//子公钥
		pubKey, err = child.ECPubKey()
	}
	compressed := pubKey.SerializeCompressed()

	//生成隔离见证比特币地址 P2WPSH
	hash160 := btcutil.Hash160(compressed)
	pubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(hash160, net)
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(compressed), pubKeyHash.EncodeAddress(), nil
}

// GenerateMulAddress
// 生成多签比特币地址
// n 需要签名的数量 m/n ex 3/2
func GenerateMulAddress(n int, pubKeys ...string) (address string, err error) {
	// 验证
	pubLen := len(pubKeys)
	if pubLen == 0 {
		return "", constant.MUL_SCRIPT_NUMBER_ERROR
	}
	if pubLen < n {
		return "", constant.MUL_ADDRESS_NUMBER_VERIFY_ERROR
	}

	// convert
	var addressPublicKeyList []*btcutil.AddressPubKey
	for _, key := range pubKeys {
		pubKey, err := convertToPubKey(key)
		if err != nil {
			return "", err
		}
		addressPublicKeyList = append(addressPublicKeyList, pubKey)
	}

	//生成多签脚本
	script, err := txscript.MultiSigScript(addressPublicKeyList, n)
	if err != nil {
		return "", err
	}
	//生成多签比特币地址
	hash, err := btcutil.NewAddressScriptHash(script, net)
	if err != nil {
		return "", err
	}
	return hash.EncodeAddress(), nil
}

func convertToPubKey(data string) (pubKey *btcutil.AddressPubKey, err error) {
	decode, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return btcutil.NewAddressPubKey(decode, net)
}

func SetNet(env string) {
	net = chooseNet(env)
}

func chooseNet(env string) *chaincfg.Params {
	if "test" == env {
		return &chaincfg.TestNet3Params
	} else {
		return &chaincfg.MainNetParams
	}
}

func NewFromSeed(mnemonic string) []byte {
	return bip39.NewSeed(mnemonic, "")
}

func NewMnemonic() (mnemonic string, err error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}
