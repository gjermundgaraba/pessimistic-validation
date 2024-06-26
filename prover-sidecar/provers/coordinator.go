package provers

import (
	"context"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"proversidecar/config"
	"proversidecar/provers/chainprover"
	"proversidecar/provers/cosmos"
	"time"
)

const (
	defaultMinQueryLoopDuration      = 1 * time.Second
)

type Coordinator struct {
	logger *zap.Logger

	chainProvers map[string]chainprover.ChainProver
}

func NewCoordinator(logger *zap.Logger, sidecarConfig config.Config) (*Coordinator, error) {
	chainProvers := make(map[string]chainprover.ChainProver)
	for _, cosmosConfig := range sidecarConfig.CosmosChains {
		prover, err := cosmos.NewCosmosProver(logger, cosmosConfig.ChainID, cosmosConfig.RPC, cosmosConfig.ClientID)
		if err != nil {
			return nil, err
		}
		chainProvers[cosmosConfig.ChainID] = prover
	}

	return &Coordinator{
		logger:       logger,
		chainProvers: chainProvers,
	}, nil
}

func (c *Coordinator) GetChainProver(chainID string) chainprover.ChainProver {
	return c.chainProvers[chainID]
}

func (c *Coordinator) Run(ctx context.Context) error {
	c.logger.Debug("Coordinator.Run")

	var eg errgroup.Group
	runCtx, runCtxCancel := context.WithCancel(ctx)
	for _, chainProver := range c.chainProvers {
		c.logger.Info("Starting chain prover loop", zap.String("chain_id", chainProver.ChainID()))

		chainProver := chainProver
		eg.Go(func() error {
			err := c.chainProverLoop(runCtx, chainProver)
			runCtxCancel() // Signal the other chain processors to exit.
			return err
		})
	}

	err := eg.Wait()
	runCtxCancel()
	return err
}

func (c *Coordinator) chainProverLoop(ctx context.Context, chainProver chainprover.ChainProver) error {
	ticker := time.NewTicker(defaultMinQueryLoopDuration) // TODO: Make this configurable per chain
	defer ticker.Stop()

	for {
		// TODO: Add retry logic
		c.logger.Info("Collecting proofs", zap.String("chain_id", chainProver.ChainID()))
		if err := chainProver.CollectProofs(ctx); err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			ticker.Reset(defaultMinQueryLoopDuration)
		}
	}
}
