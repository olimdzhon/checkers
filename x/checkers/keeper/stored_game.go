package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/olimdzhon/checkers/x/checkers/types"
)

// SetStoredGame set a specific storedGame in the store from its index
func (k Keeper) SetStoredGame(ctx context.Context, storedGame types.StoredGame) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoredGameKeyPrefix))
	b := k.cdc.MustMarshal(&storedGame)
	store.Set(types.StoredGameKey(
		storedGame.Index,
	), b)
}

// GetStoredGame returns a storedGame from its index
func (k Keeper) GetStoredGame(
	ctx context.Context,
	index string,

) (val types.StoredGame, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoredGameKeyPrefix))

	b := store.Get(types.StoredGameKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredGame removes a storedGame from the store
func (k Keeper) RemoveStoredGame(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoredGameKeyPrefix))
	store.Delete(types.StoredGameKey(
		index,
	))
}

// GetAllStoredGame returns all storedGame
func (k Keeper) GetAllStoredGame(ctx context.Context) (list []types.StoredGame) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoredGameKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredGame
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
