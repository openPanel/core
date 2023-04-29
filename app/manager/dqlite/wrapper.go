package dqlite

import (
	"context"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/canonical/go-dqlite/app"

	"github.com/openPanel/core/app/bootstrap/clean"
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/utils/fileUtils"
)

var sharedClient *shared.Client

func createSharedDatabase(clusterAddrs *[]string) (*shared.Client, error) {
	options := []app.Option{
		app.WithAddress(global.App.NodeInfo.ServerId),
		app.WithLogFunc(getDqliteLogger()),
		app.WithExternalConn(DialFunction, AcceptChan),
	}
	if clusterAddrs != nil {
		options = append(options, app.WithCluster(*clusterAddrs))
	}

	dqliteDir, err := fileUtils.RequireDataDir(constant.DqliteDataDir)
	if err != nil {
		return nil, err
	}

	dqliteApp, err := app.New(dqliteDir, options...)
	if err != nil {
		return nil, err
	}

	err = dqliteApp.Ready(context.Background())
	if err != nil {
		return nil, err
	}

	db, err := dqliteApp.Open(context.Background(), constant.SharedDatabaseDSN)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	client := shared.NewClient(shared.Driver(sql.OpenDB(dialect.SQLite, db)))

	err = client.Schema.Create(context.Background())
	if err != nil {
		return nil, err
	}

	clean.RegisterCleanup(func() {
		err := db.Close()
		if err != nil {
			log.Warn("Failed to close shared database: %v", err)
		}
		err = dqliteApp.Handover(context.Background())
		if err != nil {
			log.Warn("Failed to handover dqlite: %v", err)
		}
		err = dqliteApp.Close()
		if err != nil {
			log.Warn("Failed to close dqlite: %v", err)
		}
		log.Infof("Shared database closed")
	})

	return client, nil
}

func CreateSharedDatabase(serverAddr *[]string) *shared.Client {
	var err error
	sharedClient, err = createSharedDatabase(serverAddr)
	if err != nil {
		log.Panicf("Failed to create shared database: %v", err)
	}

	return sharedClient
}
