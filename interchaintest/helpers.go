package pessimisticinterchaintest

import (
	"context"
	"cosmossdk.io/math"
	"encoding/json"
	"fmt"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testreporter"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func (s *E2ETestSuite) SendTx(node *cosmos.ChainNode, keyName string, command ...string) error {
	command = append(command, "--keyring-dir", "/var/cosmos-chain/hub") // Something off with the hubd binary
	txHash, err := node.ExecTx(s.ctx, keyName, command...)
	if err != nil {
		return err
	}

	txRespBz, _, err := node.ExecQuery(s.ctx, "tx", txHash)
	if err != nil {
		return err
	}
	fmt.Println("tx resp", string(txRespBz))
	var txResp TxResponse
	if err := json.Unmarshal(txRespBz, &txResp); err != nil {
		return err
	}
	if txResp.Code != 0 {
		return fmt.Errorf("tx failed with code %d: %s", txResp.Code, txResp.RawLog)
	}
	return nil
}

func IBCTransferWorksTest(
	t *testing.T,
	ctx context.Context,
	r ibc.Relayer,
	eRep *testreporter.RelayerExecReporter,
	path string,
	srcChain *cosmos.CosmosChain,
	dstChain *cosmos.CosmosChain,
	srcUser ibc.Wallet,
	dstUser ibc.Wallet,
	srcChannel string,
	dstChannel string) {
	// Wait a few blocks for relayer to start and for user accounts to be created
	err := testutil.WaitForBlocks(ctx, 5, srcChain, dstChain)
	require.NoError(t, err)

	srcUserAddr := srcUser.FormattedAddress()
	dstUserAddr := dstUser.FormattedAddress()

	// Get original account balances
	srcOrigBal, err := srcChain.GetBalance(ctx, srcUserAddr, srcChain.Config().Denom)
	require.NoError(t, err)

	// Compose an IBC transfer and send from srcChain -> dstChain
	var transferAmount = math.NewInt(1_000)
	transfer := ibc.WalletAmount{
		Address: dstUserAddr,
		Denom:   srcChain.Config().Denom,
		Amount:  transferAmount,
	}

	_, err = srcChain.SendIBCTransfer(ctx, srcChannel, srcUserAddr, transfer, ibc.TransferOptions{})
	require.NoError(t, err)

	require.NoError(t, testutil.WaitForBlocks(ctx, 5, srcChain))
	require.NoError(t, r.Flush(ctx, eRep, path, srcChannel))
	require.NoError(t, r.Flush(ctx, eRep, path, dstChannel))
	require.NoError(t, testutil.WaitForBlocks(ctx, 5, srcChain))

	// Get the IBC denom for srcChain on dstChain
	srcTokenDenom := transfertypes.GetPrefixedDenom("transfer", dstChannel, srcChain.Config().Denom)
	srcIBCDenom := transfertypes.ParseDenomTrace(srcTokenDenom).IBCDenom()

	// Assert that the funds are no longer present in user acc on srcChain and are in the user acc on dstChain
	srcUpdateBal, err := srcChain.GetBalance(ctx, srcUserAddr, srcChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, srcOrigBal.Sub(transferAmount), srcUpdateBal)

	dstUpdateBal, err := dstChain.GetBalance(ctx, dstUserAddr, srcIBCDenom)
	require.NoError(t, err)
	require.Equal(t, transferAmount, dstUpdateBal)

	// Compose an IBC transfer and send from dstChain -> srcChain
	transfer = ibc.WalletAmount{
		Address: srcUserAddr,
		Denom:   srcIBCDenom,
		Amount:  transferAmount,
	}

	_, err = dstChain.SendIBCTransfer(ctx, dstChannel, dstUserAddr, transfer, ibc.TransferOptions{})
	require.NoError(t, err)

	require.NoError(t, testutil.WaitForBlocks(ctx, 5, srcChain))
	require.NoError(t, r.Flush(ctx, eRep, path, srcChannel))
	require.NoError(t, r.Flush(ctx, eRep, path, dstChannel))
	require.NoError(t, testutil.WaitForBlocks(ctx, 5, srcChain))

	// Assert that the funds are now back on srcChain and not on dstChain
	srcUpdateBal, err = srcChain.GetBalance(ctx, srcUserAddr, srcChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, srcOrigBal, srcUpdateBal)

	dstUpdateBal, err = dstChain.GetBalance(ctx, dstUserAddr, srcIBCDenom)
	require.NoError(t, err)
	require.Equal(t, int64(0), dstUpdateBal.Int64())
}

