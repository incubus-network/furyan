package keeper

import (
	"github.com/incubus-network/fury/x/mint/types"
)

var _ types.QueryServer = Keeper{}
