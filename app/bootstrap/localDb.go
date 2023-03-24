package bootstrap

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"

	"github.com/openPanel/core/app/generated/db/local"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/utils"
)

func getInitLocalDatabase() *local.Client {
	const filename = "core.local.db"
	file, err := utils.RequireDataFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open local database file: %s", err)
	}
	_ = file.Close()

	path, err := filepath.Abs(file.Name())
	if err != nil {
		log.Fatalf("Failed to get absolute path of local database file: %s", err)
	}

	dbString := fmt.Sprintf("file:%s?cache=shared&mode=rwc&_journal_mode=WAL&_fk=1&_timeout=5000", path)

	client, err := local.Open(dialect.SQLite, dbString)
	if err != nil {
		log.Fatalf("Failed to open local database: %s", err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to create local database schema: %s", err)
	}

	return client
}
