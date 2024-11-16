package auth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func (s *service) CreateWallet() (Wallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return Wallet{}, fmt.Errorf("crypto.GenerateKey: %w", err)
	}

	return Wallet{
		privateKey: privateKey,
	}, nil
}
