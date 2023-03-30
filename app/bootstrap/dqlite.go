package bootstrap

import (
	"github.com/openPanel/core/app/manager/dqlite"
)

func createDqlite() {
	dqlite.CreateSharedDatabase(nil)
}
