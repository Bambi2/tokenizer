package usecases

import (
	"context"

	"github.com/bambi2/tokenizer/deployment/internal/services/auth"
	"github.com/bambi2/tokenizer/deployment/internal/services/chainactions"
)

type UseCases interface {
	Deploy(ctx context.Context, req DeployRequest) error
	CreateWallet() (CreateWalletResponse, error)
}

type useCases struct {
	chainActionsService chainactions.Service
	authService         auth.Service
}

func New(
	chainActionsService chainactions.Service,
	authService auth.Service,
) UseCases {
	return &useCases{
		chainActionsService: chainActionsService,
		authService:         authService,
	}
}
