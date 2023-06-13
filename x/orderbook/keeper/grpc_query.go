package keeper

import (
	"github.com/incubus-network/fury/x/orderbook/types"
)

var _ types.QueryServer = Keeper{}
