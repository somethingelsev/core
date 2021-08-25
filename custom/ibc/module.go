package ibc

import (
	ibc "github.com/cosmos/ibc-go/modules/core"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"

	customtypes "github.com/terra-money/core/custom/ibc/types"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the distribution module.
type AppModuleBasic struct {
	ibc.AppModuleBasic
}

// RegisterLegacyAminoCodec registers the bank module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	customtypes.RegisterLegacyAminoCodec(cdc)
}
