package keeper

import (
	"github.com/incubus-network/fury/x/bet/types"
)

var _ types.QueryServer = Keeper{}
