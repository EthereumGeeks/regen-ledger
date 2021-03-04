package server

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/simulation"
)

func (s serverImpl) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {

	queryClient := group.NewQueryClient(s.key)
	return simulation.WeightedOperations(
		simState.AppParams, simState.Cdc,
		s.accKeeper, s.bankKeeper, queryClient,
	)
}