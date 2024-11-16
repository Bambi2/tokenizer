package usecases

import (
	"context"
	"fmt"
	"math/big"

	"github.com/bambi2/tokenizer/deployment/internal/services/chainactions"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type DeployRequest struct {
	DeployerPrivateKey                    string
	MultiSigNumberOfRequiredConfirmations uint
	MultiSigOwners                        []string
}

func (u *useCases) Deploy(ctx context.Context, req DeployRequest) error {
	parsedReq, err := u.validateAndParseDeployReq(req)
	if err != nil {
		return err
	}

	if err := u.chainActionsService.Deploy(ctx, parsedReq); err != nil {
		return fmt.Errorf("chainActionsService.Deploy: %w", err)
	}

	return nil
}

func (u *useCases) validateAndParseDeployReq(req DeployRequest) (chainactions.DeployRequest, error) {
	if req.MultiSigNumberOfRequiredConfirmations < 2 {
		return chainactions.DeployRequest{}, fmt.Errorf("multisig number of required confirmations should be at least 2")
	}

	if uint(len(req.MultiSigOwners)) < req.MultiSigNumberOfRequiredConfirmations {
		return chainactions.DeployRequest{}, fmt.Errorf("number of owners is less than the number required confirmations")
	}

	multiSigOwners := make([]common.Address, 0, len(req.MultiSigOwners))
	for _, multiSigOwner := range req.MultiSigOwners {
		if !common.IsHexAddress(multiSigOwner) {
			return chainactions.DeployRequest{}, fmt.Errorf("%s is not a valid address", multiSigOwner)
		}

		multiSigOwners = append(multiSigOwners, common.HexToAddress(multiSigOwner))
	}

	if req.DeployerPrivateKey == "" {
		return chainactions.DeployRequest{}, fmt.Errorf("empty private key")
	}

	privateKey, err := crypto.ToECDSA(common.FromHex(req.DeployerPrivateKey))
	if err != nil {
		return chainactions.DeployRequest{}, fmt.Errorf("private key: %w", err)
	}

	return chainactions.DeployRequest{
		DeployerPrivateKey:                    privateKey,
		MultiSigNumberOfRequiredConfirmations: big.NewInt(int64(req.MultiSigNumberOfRequiredConfirmations)),
		MultiSigOwners:                        multiSigOwners,
	}, nil
}
