package test

import (
	"generate_btc/util"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
	"log"
	"testing"
)

var mnemonic = "range sheriff try enroll deer over ten level bring display stamp recycle"

func TestGeneratePublicKeyAndAddress(t *testing.T) {
	//从助记词生成种子
	seed, err := util.NewFromSeed(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	path := util.NewPath("m/49'/0'/0'/0/0")
	pubKeyS, pkwsh, err := util.GenerateHDWalletSegWitAddress(seed, path)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("pubKey:", pubKeyS)
	log.Println("p2wpsh:", pkwsh)
}
func TestGeneratePublicKeyAndAddress2(t *testing.T) {
	path := "m/49'/0'/0'/0/0"
	pubKeyS, pkwsh, err := util.GenerateHDWalletSegWitAddress2(mnemonic, path)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("pubKey:", pubKeyS)
	log.Println("p2wpsh:", pkwsh)
}

func TestMulAddress(t *testing.T) {
	pubKeyList := []string{
		"038ea00f7eb51072243de550ae3b641d755e714f42e726116f69631405d88a9f1e",
		"03b6cd4d40fcaf70bb3f721651baf9b7a73bffba60b5c6cd260975b9988f0e163b",
		"023f77c6729a74ac8b402d169e526d1cf3aa315bf6db4e89bc0c92b6cc8cb38987",
	}
	address, err := util.GenerateMulAddress(2, pubKeyList...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("multi address:", address)
}

func TestGenerateMasterAddress(t *testing.T) {
	//从助记词生成种子
	seed, err := util.NewFromSeed(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	master, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		log.Panic(err)
	}
	var pubKey *btcec.PublicKey
	pubKey, err = master.ECPubKey()

	print(pubKey)
}

func print(pubKey *btcec.PublicKey) {
	p2wpkhAddress, err := util.GenerateP2WPKHAddress(pubKey)
	if err != nil {
		log.Panic(err)
	}
	pkhAddress, err := util.GenerateP2PKHAddress(pubKey)
	if err != nil {
		log.Panic(err)
	}
	pkwsh, err := util.GenerateP2WSHAddress(pubKey)

	log.Println("p2pkh:", pkhAddress)
	log.Println("p2wpkh:", p2wpkhAddress)
	log.Println("p2wpsh:", pkwsh)
}

func TestGenerateChildAddress44(t *testing.T) {
	// 从助记词生成种子
	seed := bip39.NewSeed(mnemonic, "")

	// 从种子生成主私钥
	master, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		log.Panic(err)
	}

	// 生成子私钥
	p, _ := master.Child(hdkeychain.HardenedKeyStart + 44)
	c, _ := p.Child(hdkeychain.HardenedKeyStart + 0)
	a, _ := c.Child(hdkeychain.HardenedKeyStart + 0)
	cg, _ := a.Child(0)
	ai, _ := cg.Child(0)
	key, err := ai.ECPubKey()
	print(key)
}

func TestGenerateChildAddress49(t *testing.T) {
	// 从助记词生成种子
	seed := bip39.NewSeed(mnemonic, "")

	// 从种子生成主私钥
	master, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		log.Panic(err)
	}

	// 生成子私钥
	p, _ := master.Child(hdkeychain.HardenedKeyStart + 49)
	c, _ := p.Child(hdkeychain.HardenedKeyStart + 0)
	a, _ := c.Child(hdkeychain.HardenedKeyStart + 0)
	cg, _ := a.Child(0)
	ai, _ := cg.Child(0)
	key, err := ai.ECPubKey()
	print(key)
}
