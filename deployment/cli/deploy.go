package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/bambi2/tokenizer/deployment/internal/usecases"
	"github.com/spf13/cobra"
)

func (s *Server) newDeployCmd(ctx context.Context) *cobra.Command {
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy G42 and MultiSig contracts to helesky blockchain",
		Run: func(cmd *cobra.Command, args []string) {
			pkey, err := cmd.Flags().GetString("pkey")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			numberOfRequiredConfirmations, err := cmd.Flags().GetUint("norc")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			multiSigOwners, err := cmd.Flags().GetStringSlice("owners")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			if err := s.useCases.Deploy(ctx, usecases.DeployRequest{
				DeployerPrivateKey:                    pkey,
				MultiSigNumberOfRequiredConfirmations: numberOfRequiredConfirmations,
				MultiSigOwners:                        multiSigOwners,
			}); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		},
	}

	deployCmd.Flags().String("pkey", "", "private key of the sender")
	deployCmd.Flags().Uint("norc", 0, "number of required confirmations for multi sig")
	deployCmd.Flags().StringSlice("owners", nil, "owners, who can confirm the multi sig transactions")

	return deployCmd
}
