package types

const (
	// ModuleName defines the module name
	ModuleName = "ltwotx"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ltwotx"
)

var (
	ParamsKey = []byte("p_ltwotx")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
