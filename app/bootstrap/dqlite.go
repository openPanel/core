package bootstrap

import (
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/manager/dqlite"
)

func createDqlite() *shared.Client {
	return dqlite.CreateSharedDatabase(nil)
}
