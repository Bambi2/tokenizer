package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (s *Server) newCreateWalletCmd() *cobra.Command {
	createWalletCmd := &cobra.Command{
		Use:   "createWallet",
		Short: "Create a new wallet",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := s.useCases.CreateWallet()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			fmt.Printf("Private Key: %s\n", resp.PrivateKeyHex)
			fmt.Printf("Public Key: %s\n", resp.PublicKeyHex)
			fmt.Printf("Address: %s\n", resp.AddressHex)
		},
	}

	return createWalletCmd
}
