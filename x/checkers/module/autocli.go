package checkers

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/olimdzhon/checkers/api/checkers/checkers"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "SystemInfo",
					Use:       "show-system-info",
					Short:     "show systemInfo",
				},
				{
					RpcMethod: "StoredGameAll",
					Use:       "list-stored-game",
					Short:     "List all storedGame",
				},
				{
					RpcMethod:      "StoredGame",
					Use:            "show-stored-game [id]",
					Short:          "Shows a storedGame",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateGame",
					Use:            "create-game [black] [red]",
					Short:          "Send a createGame tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "black"}, {ProtoField: "red"}},
				},
				{
					RpcMethod:      "PlayMove",
					Use:            "play-move [game-index] [from-x] [from-y] [to-x] [to-y]",
					Short:          "Send a playMove tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "gameIndex"}, {ProtoField: "fromX"}, {ProtoField: "fromY"}, {ProtoField: "toX"}, {ProtoField: "toY"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
