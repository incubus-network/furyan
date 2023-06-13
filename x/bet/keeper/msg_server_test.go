package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	simappUtil "github.com/incubus-network/fury/testutil/simapp"
	"github.com/incubus-network/fury/x/bet/keeper"
	"github.com/incubus-network/fury/x/bet/types"
)

func setupMsgServerAndApp(t testing.TB) (*simappUtil.TestApp, *keeper.KeeperTest, types.MsgServer, sdk.Context, context.Context) {
	tApp, k, ctx := setupKeeperAndApp(t)
	return tApp, k, keeper.NewMsgServerImpl(*k), ctx, sdk.WrapSDKContext(ctx)
}
