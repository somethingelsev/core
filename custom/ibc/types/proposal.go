package types

import (
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func init() {
	govtypes.RegisterProposalTypeCodec(&clienttypes.ClientUpdateProposal{}, "ibc/ClientUpdateProposal")
	govtypes.RegisterProposalTypeCodec(&clienttypes.UpgradeProposal{},
		"ibc/UpgradeProposal")
}
