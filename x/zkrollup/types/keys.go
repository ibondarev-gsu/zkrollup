package types

const (
	// ModuleName defines the module name
	ModuleName = "zkrollup"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_zkrollup"
)

var (
	ParamsKey = []byte("p_zkrollup")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
