package db

import (
	"github.com/openPanel/core/app/generated/db/local"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/global"
)

func GetSharedDb() *shared.Client {
	return global.App.DbShared
}

func GetLocalDb() *local.Client {
	return global.App.DbLocal
}
