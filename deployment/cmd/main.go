package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bambi2/tokenizer/deployment/cli"
	"github.com/bambi2/tokenizer/deployment/internal/services/auth"
	"github.com/bambi2/tokenizer/deployment/internal/services/chainactions"
	"github.com/bambi2/tokenizer/deployment/internal/services/chainactions/holesky"
	"github.com/bambi2/tokenizer/deployment/internal/usecases"
	"github.com/oklog/oklog/pkg/group"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	holeskyClient, err := holesky.DialClient(ctx)
	if err != nil {
		return fmt.Errorf("holesky.DialClient: %w", err)
	}

	chainActionsService := chainactions.NewService(holeskyClient)
	authService := auth.NewService()

	useCases := usecases.New(chainActionsService, authService)

	cliServer := cli.NewServer(useCases)

	g := group.Group{}

	{
		term := make(chan os.Signal, 1)
		signal.Notify(term, os.Interrupt, syscall.SIGTERM)
		g.Add(
			func() error {
				select {
				case <-term:
					fmt.Println("Stopping all the processes...")
				case <-ctx.Done():
					break
				}
				return nil
			},
			func(error) {
				cancelCtx()
			},
		)
	}

	g.Add(
		func() error {
			if err := cliServer.Execute(ctx); err != nil {
				return fmt.Errorf("cliServer.Execute: %w", err)
			}

			return nil
		},
		func(err error) {
			cancelCtx()
		},
	)

	if err := g.Run(); err != nil {
		return fmt.Errorf("g.Run(): %w", err)
	}

	return nil
}
