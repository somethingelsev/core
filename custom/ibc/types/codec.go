package types

import (
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	commitmenttypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	solomachinetypes "github.com/cosmos/ibc-go/modules/light-clients/06-solomachine/types"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	localhosttypes "github.com/cosmos/ibc-go/modules/light-clients/09-localhost/types"

	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterLegacyAminoCodec registers the necessary ibc-go interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {

	// Client
	cdc.RegisterInterface((*exported.ClientState)(nil), nil)
	cdc.RegisterInterface((*exported.ConsensusState)(nil), nil)
	cdc.RegisterInterface((*exported.Header)(nil), nil)
	cdc.RegisterInterface((*exported.Height)(nil), nil)
	cdc.RegisterInterface((*exported.Misbehaviour)(nil), nil)

	cdc.RegisterConcrete(&clienttypes.ClientUpdateProposal{}, "ibc/ClientUpdateProposal", nil)
	cdc.RegisterConcrete(&clienttypes.UpgradeProposal{}, "ibc/UpgradeProposal", nil)
	cdc.RegisterConcrete(&clienttypes.MsgCreateClient{}, "ibc/MsgCreateClient", nil)
	cdc.RegisterConcrete(&clienttypes.MsgUpdateClient{}, "ibc/MsgUpdateClient", nil)
	cdc.RegisterConcrete(&clienttypes.MsgUpgradeClient{}, "ibc/MsgUpgradeClient", nil)
	cdc.RegisterConcrete(&clienttypes.MsgSubmitMisbehaviour{}, "ibc/MsgSubmitMisbehaviour", nil)

	// Connection
	cdc.RegisterInterface((*exported.ConnectionI)(nil), nil)
	cdc.RegisterInterface((*exported.CounterpartyConnectionI)(nil), nil)
	cdc.RegisterInterface((*exported.Version)(nil), nil)

	cdc.RegisterConcrete(&connectiontypes.MsgConnectionOpenInit{}, "ibc/MsgConnectionOpenInit", nil)
	cdc.RegisterConcrete(&connectiontypes.MsgConnectionOpenTry{}, "ibc/MsgConnectionOpenTry", nil)
	cdc.RegisterConcrete(&connectiontypes.MsgConnectionOpenAck{}, "ibc/MsgConnectionOpenAck", nil)
	cdc.RegisterConcrete(&connectiontypes.MsgConnectionOpenConfirm{}, "ibc/MsgConnectionOpenConfirm", nil)

	// Channel
	cdc.RegisterInterface((*exported.ChannelI)(nil), nil)
	cdc.RegisterInterface((*exported.CounterpartyChannelI)(nil), nil)
	cdc.RegisterInterface((*exported.PacketI)(nil), nil)

	cdc.RegisterConcrete(&channeltypes.MsgChannelOpenInit{}, "ibc/MsgChannelOpenInit", nil)
	cdc.RegisterConcrete(&channeltypes.MsgChannelOpenTry{}, "ibc/MsgChannelOpenTry", nil)
	cdc.RegisterConcrete(&channeltypes.MsgChannelOpenAck{}, "ibc/MsgChannelOpenAck", nil)
	cdc.RegisterConcrete(&channeltypes.MsgChannelOpenConfirm{}, "ibc/MsgChannelOpenConfirm", nil)
	cdc.RegisterConcrete(&channeltypes.MsgChannelCloseInit{}, "ibc/MsgChannelCloseInit", nil)
	cdc.RegisterConcrete(&channeltypes.MsgChannelCloseConfirm{}, "ibc/MsgChannelCloseConfirm", nil)
	cdc.RegisterConcrete(&channeltypes.MsgRecvPacket{}, "ibc/MsgRecvPacket", nil)
	cdc.RegisterConcrete(&channeltypes.MsgAcknowledgement{}, "ibc/MsgAcknowledgement", nil)
	cdc.RegisterConcrete(&channeltypes.MsgTimeout{}, "ibc/MsgTimeout", nil)
	cdc.RegisterConcrete(&channeltypes.MsgTimeoutOnClose{}, "ibc/MsgTimeoutOnClose", nil)

	// Solomachine
	cdc.RegisterConcrete(&solomachinetypes.ClientState{}, "ibc-solomachine/ClientState", nil)
	cdc.RegisterConcrete(&solomachinetypes.ConsensusState{}, "ibc-solomachine/ConsensusState", nil)
	cdc.RegisterConcrete(&solomachinetypes.Header{}, "ibc-solomachine/Header", nil)
	cdc.RegisterConcrete(&solomachinetypes.Misbehaviour{}, "ibc-solomachine/Misbehaviour", nil)

	// Tendermint client
	cdc.RegisterConcrete(&ibctmtypes.ClientState{}, "ibc-tendermint/ClientState", nil)
	cdc.RegisterConcrete(&ibctmtypes.ConsensusState{}, "ibc-tendermint/ConsensusState", nil)
	cdc.RegisterConcrete(&ibctmtypes.Header{}, "ibc-tendermint/Header", nil)
	cdc.RegisterConcrete(&ibctmtypes.Misbehaviour{}, "ibc-tendermint/Misbehaviour", nil)

	// Localhost
	cdc.RegisterConcrete(&localhosttypes.ClientState{}, "ibc-localhost/ClientState", nil)

	// Commitment
	cdc.RegisterInterface((*exported.Root)(nil), nil)
	cdc.RegisterInterface((*exported.Prefix)(nil), nil)
	cdc.RegisterInterface((*exported.Path)(nil), nil)
	cdc.RegisterInterface((*exported.Proof)(nil), nil)

	cdc.RegisterConcrete(&commitmenttypes.MerkleRoot{}, "ibc/MerkleRoot", nil)
	cdc.RegisterConcrete(&commitmenttypes.MerklePrefix{}, "ibc/MerklePrefix", nil)
	cdc.RegisterConcrete(&commitmenttypes.MerklePath{}, "ibc/MerklePath", nil)
	cdc.RegisterConcrete(&commitmenttypes.MerkleProof{}, "ibc/MerkleProof", nil)
}
