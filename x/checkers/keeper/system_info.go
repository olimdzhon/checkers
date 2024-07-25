package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/olimdzhon/checkers/x/checkers/types"
)

// SetSystemInfo set systemInfo in the store
func (k Keeper) SetSystemInfo(ctx context.Context, systemInfo types.SystemInfo) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SystemInfoKey))
	b := k.cdc.MustMarshal(&systemInfo)
	store.Set([]byte{0}, b)
}

// GetSystemInfo returns systemInfo
func (k Keeper) GetSystemInfo(ctx context.Context) (val types.SystemInfo, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SystemInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSystemInfo removes systemInfo from the store
func (k Keeper) RemoveSystemInfo(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SystemInfoKey))
	store.Delete([]byte{0})
}
