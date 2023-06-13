package keeper

import (
	"github.com/incubus-network/fury/x/sportevent/types"
)

var _ types.QueryServer = Keeper{}
