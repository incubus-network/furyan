package keeper

import (
	"github.com/incubus-network/fury/x/house/types"
)

var _ types.QueryServer = Keeper{}
