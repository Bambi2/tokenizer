package cli

import (
	"context"
	"fmt"

	"github.com/bambi2/tokenizer/deployment/internal/usecases"
	"github.com/spf13/cobra"
)

type Server struct {
	useCases usecases.UseCases
}

func NewServer(useCases usecases.UseCases) *Server {
	return &Server{
		useCases: useCases,
	}
}

// Execute initiates cli commands of the binary and reads all the args
func (s *Server) Execute(ctx context.Context) error {
	rootCmd := &cobra.Command{
		Use:   "g42",
		Short: "G42 is a command line wallet to interact with different blockchains and contracts",
	}

	rootCmd.AddCommand(s.newDeployCmd(ctx))
	rootCmd.AddCommand(s.newCreateWalletCmd())

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("rootCmd.Execute: %w", err)
	}

	return nil
}
