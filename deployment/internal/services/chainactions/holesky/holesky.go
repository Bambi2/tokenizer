package holesky

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	possibleHoleskyEndpoints = [...]string{
		"https://eth-holesky.public.blastapi.io",
	}
)

func DialClient(ctx context.Context) (*ethclient.Client, error) {
	for _, endpoint := range possibleHoleskyEndpoints {
		client, err := ethclient.DialContext(ctx, endpoint)
		if err == nil {
			return client, nil
		}

		log.Printf("ethclient.DialContext: %s", err)
	}

	return nil, fmt.Errorf("couldn't connect to any known sepolia endpoint")
}
