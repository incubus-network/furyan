package keeper

import (
	"github.com/incubus-network/fury/x/dvm/types"
)

var _ types.QueryServer = Keeper{}
