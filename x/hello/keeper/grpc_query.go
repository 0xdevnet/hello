package keeper

import (
	"github.com/0xkdavid/hello/x/hello/types"
)

var _ types.QueryServer = Keeper{}
