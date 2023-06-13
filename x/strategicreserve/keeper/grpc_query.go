package keeper

import (
	"github.com/incubus-network/fury/x/strategicreserve/types"
)

var _ types.QueryServer = Keeper{}
