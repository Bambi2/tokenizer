package usecases

import (
	"fmt"
)

type CreateWalletResponse struct {
	PrivateKeyHex string
	PublicKeyHex  string
	AddressHex    string
}

func (u *useCases) CreateWallet() (CreateWalletResponse, error) {
	wallet, err := u.authService.CreateWallet()
	if err != nil {
		return CreateWalletResponse{}, fmt.Errorf("authService.CreateWallet: %w", err)
	}

	return CreateWalletResponse{
		PrivateKeyHex: wallet.GetPrivateKeyHex(),
		PublicKeyHex:  wallet.GetPublicKeyHex(),
		AddressHex:    wallet.GetAddress().Hex(),
	}, nil
}
