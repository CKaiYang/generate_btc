package util

import (
	"encoding/hex"
	"generate_btc/constant"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/foxnut/go-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"strconv"
	"strings"
)

var net = &chaincfg.MainNetParams

const (
	ZeroQuote uint32 = 0x80000000
)

type Path struct {
	// 44 BIP44 49 BIP49
	Purpose uint32
	// 0 BTC 60 ETC
	CoinType uint32
	//帐户 从0开始递增
	Account uint32
	//0 对外 1 找零
	Change uint32
	//派生地址下标 从0开始递增
	AddressIndex uint32
}

// GenerateHDWalletSegWitAddress
// 按照指定路径生成HD钱包公钥已经隔离见证的比特币地址 使用
// seed 种子
// path 指定路径
func GenerateHDWalletSegWitAddress(seed []byte, path *Path) (pubKeyS string, pkwsh string, err error) {
	master, err := hdkeychain.NewMaster(seed, net)
	if err != nil {
		return "", "", err
	}
	var pubKey *btcec.PublicKey

	key := master
	for _, i := range path.GetPath() {
		key, err = key.Child(i)
		if err != nil {
			return "", "", err
		}
	}

	//子公钥
	pubKey, err = key.ECPubKey()
	compressed := pubKey.SerializeCompressed()

	p2wpkhAddress, err := GenerateP2WSHAddress(pubKey)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(compressed), p2wpkhAddress, nil
}

// GenerateHDWalletSegWitAddress2
// 按照指定路径生成HD钱包公钥已经隔离见证的比特币地址 使用
// mnemonic 助记词
// path 指定路径
func GenerateHDWalletSegWitAddress2(mnemonic string, path string) (pubKeyS string, pkwsh string, err error) {
	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
		hdwallet.Params(net),
	)
	if err != nil {
		return "", "", err
	}

	wallet, err := master.GetWallet(
		hdwallet.CoinType(hdwallet.BTC),
		hdwallet.Path(path),
	)
	if err != nil {
		return "", "", err
	}
	publicKey := wallet.GetKey().Public.SerializeCompressed()

	p2wsh, err := wallet.GetKey().AddressP2WPKHInP2SH()
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(publicKey), p2wsh, nil
}

// GenerateP2WPKHAddress 生成隔离见证地址
func GenerateP2WPKHAddress(pubKey *btcec.PublicKey) (address string, err error) {
	compressed := pubKey.SerializeCompressed()
	//生成隔离见证比特币地址 P2WPKH
	hash160 := btcutil.Hash160(compressed)
	addressPubKey, err := btcutil.NewAddressWitnessPubKeyHash(hash160, net)
	if err != nil {
		return "", err
	}
	return addressPubKey.EncodeAddress(), nil
}

// GenerateP2WSHAddress 生成隔离见证(兼容)地址
func GenerateP2WSHAddress(pubKey *btcec.PublicKey) (address string, err error) {
	compressed := pubKey.SerializeCompressed()
	//生成隔离见证比特币地址 P2WPKH
	witnessProg := btcutil.Hash160(compressed)
	addr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, net)
	if err != nil {
		return "", err
	}
	script, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return "", err
	}
	addrScr, err := btcutil.NewAddressScriptHash(script, &chaincfg.MainNetParams)
	return addrScr.EncodeAddress(), nil
}

// GenerateP2PKHAddress 生成 普通子地址
func GenerateP2PKHAddress(pubKey *btcec.PublicKey) (address string, err error) {
	compressed := pubKey.SerializeCompressed()
	addressPubKey, err := btcutil.NewAddressPubKey(compressed, net)
	if err != nil {
		return "", err
	}
	return addressPubKey.EncodeAddress(), nil
}

// GenerateMulAddress
// 生成多签比特币地址
// n 需要签名的数量 m/n ex 3/2
// pubKeys 公钥地址 注意不是比特币地址 而是公钥地址
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

func NewFromSeed(mnemonic string) (seed []byte, err error) {
	return bip39.NewSeedWithErrorChecking(mnemonic, "")
}

func NewMnemonic() (mnemonic string, err error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func NewPath(path string) *Path {
	path = strings.TrimPrefix(path, "m/")
	paths := strings.Split(path, "/")
	if len(paths) != 5 {
		return nil
	}
	p := new(Path)
	p.Purpose = PathNumber(paths[0])
	p.CoinType = PathNumber(paths[1])
	p.Account = PathNumber(paths[2])
	p.Change = PathNumber(paths[3])
	p.AddressIndex = PathNumber(paths[4])
	return p
}

func PathNumber(str string) uint32 {
	num64, _ := strconv.ParseInt(strings.TrimSuffix(str, "'"), 10, 64)
	num := uint32(num64)
	if strings.HasSuffix(str, "'") {
		num += ZeroQuote
	}
	return num
}

func (p *Path) GetPath() []uint32 {
	return []uint32{
		p.Purpose,
		p.CoinType,
		p.Account,
		p.Change,
		p.AddressIndex,
	}
}
