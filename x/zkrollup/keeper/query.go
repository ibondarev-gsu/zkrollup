package keeper

import (
	"zkrollup/x/zkrollup/types"
)

var _ types.QueryServer = Keeper{}
