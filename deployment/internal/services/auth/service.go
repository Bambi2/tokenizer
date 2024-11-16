package auth

type Service interface {
	CreateWallet() (Wallet, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}
