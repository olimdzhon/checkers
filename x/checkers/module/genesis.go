package checkers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/olimdzhon/checkers/x/checkers/keeper"
	"github.com/olimdzhon/checkers/x/checkers/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.SystemInfo != nil {
		k.SetSystemInfo(ctx, *genState.SystemInfo)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = &systemInfo
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
