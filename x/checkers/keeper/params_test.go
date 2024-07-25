package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/olimdzhon/checkers/testutil/keeper"
	"github.com/olimdzhon/checkers/x/checkers/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CheckersKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
