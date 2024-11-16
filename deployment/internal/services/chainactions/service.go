package chainactions

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Service interface {
	Deploy(ctx context.Context, req DeployRequest) error
}

type service struct {
	// Could be an abstraction for not ethereum chains
	chainClient *ethclient.Client
}

// NewService new chain-actions client
func NewService(chainClient *ethclient.Client) Service {
	return &service{
		chainClient: chainClient,
	}
}
