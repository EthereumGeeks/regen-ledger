package module

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/client"
	"github.com/regen-network/regen-ledger/x/group/server"

	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/simulation"
)

type Module struct {
	Registry      types.InterfaceRegistry
	BankKeeper    server.BankKeeper
	AccountKeeper server.AccountKeeper
	Configurator  servermodule.Configurator
}

var _ module.AppModuleBasic = Module{}
var _ module.AppModuleSimulation = Module{}
var _ servermodule.Module = Module{}
var _ climodule.Module = Module{}
var _ servermodule.LegacyRouteModule = Module{}

func (a Module) Name() string {
	return group.ModuleName
}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	group.RegisterTypes(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	server.RegisterServices(configurator, a.AccountKeeper)
}

func (a Module) DefaultGenesis(marshaler codec.JSONMarshaler) json.RawMessage {
	return marshaler.MustMarshalJSON(group.NewGenesisState())
}

func (a Module) ValidateGenesis(cdc codec.JSONMarshaler, config sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	var data group.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", group.ModuleName, err)
	}
	return data.Validate()
}

func (a Module) RegisterGRPCGatewayRoutes(sdkclient.Context, *runtime.ServeMux) {}

func (a Module) GetTxCmd() *cobra.Command {
	return client.TxCmd(a.Name())
}

func (a Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd(a.Name())
}

func (a Module) GenerateGenesisState(input *module.SimulationState) {
}

func (a Module) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

func (a Module) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return nil
}

func (a Module) RegisterStoreDecoder(registry sdk.StoreDecoderRegistry) {
}

func (a Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	protoCdc := codec.NewProtoCodec(a.Registry)
	querySrvr := group.NewQueryClient(a.Configurator.ModuleKey())
	return simulation.WeightedOperations(
		simState.AppParams, simState.Cdc,
		a.AccountKeeper, a.BankKeeper, protoCdc, querySrvr,
	)
}

/**** DEPRECATED ****/
func (a Module) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}
func (a Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	group.RegisterLegacyAminoCodec(cdc)
}

func (a Module) Route(configurator servermodule.Configurator) sdk.Route {
	return sdk.NewRoute(group.RouterKey, server.NewHandler(configurator, a.AccountKeeper))
}
