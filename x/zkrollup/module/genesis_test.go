package zkrollup_test

import (
	"testing"

	keepertest "zkrollup/testutil/keeper"
	"zkrollup/testutil/nullify"
	zkrollup "zkrollup/x/zkrollup/module"
	"zkrollup/x/zkrollup/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ZkrollupKeeper(t)
	zkrollup.InitGenesis(ctx, k, genesisState)
	got := zkrollup.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
