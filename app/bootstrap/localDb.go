package bootstrap

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"

	"github.com/openPanel/core/app/bootstrap/clean"
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/db/local"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/utils/fileUtils"
)

func getInitLocalDatabase() *local.Client {
	dir, err := fileUtils.RequireDataDir("")
	if err != nil {
		log.Fatalf("Failed to create local database directory: %s", err)
	}

	path, err := filepath.Abs(dir + string(os.PathSeparator) + constant.DefaultLocalSqliteFilename)
	if err != nil {
		log.Fatalf("Failed to get absolute path of local database file: %s", err)
	}

	DSN := fmt.Sprintf("file:%s?cache=shared&mode=rwc&_journal_mode=WAL&_fk=1&_timeout=5000", path)

	client, err := local.Open(dialect.SQLite, DSN)
	if err != nil {
		log.Fatalf("Failed to open local database: %s", err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to create local database schema: %s", err)
	}

	log.Infof("Local database initialized at %s", path)

	clean.RegisterCleanup(func() {
		err := client.Close()
		if err != nil {
			log.Warn("Failed to close local database: %v", err)
		}
		log.Infof("Local database closed")
	})

	return client
}

func initLocalDatabase() {
	global.App.DbLocal = getInitLocalDatabase()
}
