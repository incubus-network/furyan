package mint_test

import (
	"testing"

	"github.com/incubus-network/fury/testutil/nullify"
	simappUtil "github.com/incubus-network/fury/testutil/simapp"
	"github.com/incubus-network/fury/x/mint"
	"github.com/incubus-network/fury/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	mint.InitGenesis(ctx, tApp.MintKeeper, genesisState)
	got := mint.ExportGenesis(ctx, tApp.MintKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
