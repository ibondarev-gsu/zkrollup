package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "zkrollup/testutil/keeper"
	"zkrollup/x/ltwotx/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LtwotxKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
