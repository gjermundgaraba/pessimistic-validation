package lightclient_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/gogoproto/proto"
	"github.com/gjermundgaraba/pessimistic-validation/lightclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCodec(t *testing.T) {
	encodingCfg := moduletestutil.MakeTestEncodingConfig(lightclient.AppModuleBasic{})
	attestators := generateAttestators(10)
	validClientMsg := generateClientMsg(encodingCfg.Codec, attestators, 5)

	testCases := []struct {
		name            string
		toMarshalFrom   proto.Message
		toUnmarshalInto proto.Message
	} {
		{
			"ClientState",
			initialClientState,
			&lightclient.ClientState{},
		},
		{
			"ConsensusState",
			initialConsensusState,
			&lightclient.ConsensusState{},
		},
		{
			"PessimisticClaims",
			validClientMsg,
			&lightclient.PessimisticClaims{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			bz, err := encodingCfg.Codec.Marshal(tc.toMarshalFrom)
			require.NoError(t, err)
			require.NotNil(t, bz)

			err = encodingCfg.Codec.Unmarshal(bz, tc.toUnmarshalInto)
			require.NoError(t, err)

			msg, err := encodingCfg.Codec.InterfaceRegistry().Resolve(sdk.MsgTypeURL(tc.toUnmarshalInto))
			require.NoError(t, err)
			require.NotNil(t, msg)
		})
	}
}
