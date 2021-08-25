package types

import (
	transfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"

	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterLegacyAminoCodec registers the necessary ibc-go interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// Transfer
	cdc.RegisterConcrete(&transfertypes.MsgTransfer{}, "ibc/MsgTransfer", nil)
}
