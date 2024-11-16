package chainactions

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/bambi2/tokenizer/deployment/g42"
	"github.com/bambi2/tokenizer/deployment/multisig"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type DeployRequest struct {
	DeployerPrivateKey                    *ecdsa.PrivateKey
	MultiSigNumberOfRequiredConfirmations *big.Int
	MultiSigOwners                        []common.Address
}

func (s *service) Deploy(ctx context.Context, req DeployRequest) error {
	nonce, err := s.chainClient.PendingNonceAt(ctx, crypto.PubkeyToAddress(req.DeployerPrivateKey.PublicKey))
	if err != nil {
		return fmt.Errorf("s.chainClient.PendingNonceAt: %w", err)
	}

	gasPrice, err := s.chainClient.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("s.chainClient.SuggestGasPrice: %w", err)
	}

	chainID, err := s.chainClient.NetworkID(ctx)
	if err != nil {
		return fmt.Errorf("s.chainClient.NetworkID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(req.DeployerPrivateKey, chainID)
	if err != nil {
		return fmt.Errorf("bind.NewKeyedTransactorWithChainID: %w", err)
	}

	auth.GasLimit = 3_000_000
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))

	multiSigAddress, multiSigTx, _, err := multisig.DeployMultisig(
		auth,
		s.chainClient,
		req.MultiSigOwners,
		req.MultiSigNumberOfRequiredConfirmations,
	)
	if err != nil {
		return fmt.Errorf("multisig.DeployMultisig: %w", err)
	}

	auth.GasLimit = 3_000_000
	auth.GasPrice = gasPrice
	auth.Nonce.SetUint64(nonce + 1)

	g42Address, g42Tx, _, err := g42.DeployG42(
		auth,
		s.chainClient,
		multiSigAddress,
	)
	if err != nil {
		return fmt.Errorf("g42.DeployG42: %w", err)
	}

	fmt.Printf("MultiSig Address: %s\n", multiSigAddress.Hex())
	fmt.Printf("MultiSig Creation Transaction Hash: %s\n", multiSigTx.Hash().Hex())

	fmt.Printf("G42 Address: %s\n", g42Address.Hex())
	fmt.Printf("G42 Creation Transaction Hash: %s\n", g42Tx.Hash().Hex())

	return nil
}
