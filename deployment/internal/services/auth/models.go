package auth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
}

func (w Wallet) GetPrivateKeyHex() string {
	if w.privateKey == nil {
		return ""
	}

	return hexutil.Encode(crypto.FromECDSA(w.privateKey))
}

func (w Wallet) GetPublicKeyHex() string {
	if w.privateKey == nil {
		return ""
	}

	return hexutil.Encode(crypto.FromECDSAPub(&w.privateKey.PublicKey))
}

func (w Wallet) GetAddress() common.Address {
	if w.privateKey == nil {
		return common.Address{}
	}

	return crypto.PubkeyToAddress(w.privateKey.PublicKey)
}
