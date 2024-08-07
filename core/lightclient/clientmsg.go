package lightclient

import (
	"crypto/sha256"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/gjermundgaraba/pessimistic-validation/core/types"
)

var _ exported.ClientMessage = (*AttestationClaim)(nil)

func NewPessimisticClaims(attestation []types.Attestation) *AttestationClaim {
	return &AttestationClaim{
		Attestations: attestation,
	}
}

func (m *AttestationClaim) ClientType() string {
	return ModuleName
}

func (m *AttestationClaim) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func GetSignableBytes(cdc codec.BinaryCodec, dataToAttestTo types.IBCData) []byte {
	packetBytes := cdc.MustMarshal(&dataToAttestTo)
	hash := sha256.Sum256(packetBytes)
	return hash[:]
}