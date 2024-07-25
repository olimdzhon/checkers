package types

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

var (
	ParamsKey = []byte("p_checkers")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
