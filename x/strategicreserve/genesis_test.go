package strategicreserve_test

import (
	"testing"

	"github.com/incubus-network/fury/testutil/nullify"
	simappUtil "github.com/incubus-network/fury/testutil/simapp"
	"github.com/incubus-network/fury/x/strategicreserve"
	"github.com/incubus-network/fury/x/strategicreserve/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	strategicreserve.InitGenesis(ctx, tApp.StrategicReserveKeeper, genesisState)
	got := strategicreserve.ExportGenesis(ctx, tApp.StrategicReserveKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
