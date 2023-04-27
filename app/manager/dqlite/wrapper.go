package dqlite

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
	dqliteApp "github.com/canonical/go-dqlite/app"

	"entgo.io/ent/dialect/sql"

	"github.com/openPanel/core/app/bootstrap/clean"
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/utils/fileUtils"
)

var createOnce = sync.Once{}
var sharedClient *shared.Client

func createSharedDatabase(clusterAddrs *[]string) (*shared.Client, error) {
	options := []dqliteApp.Option{
		dqliteApp.WithAddress(global.App.NodeInfo.ServerId),
		dqliteApp.WithLogFunc(getDqliteLogger()),
		dqliteApp.WithExternalConn(DialFunction, AcceptChan),
	}
	if clusterAddrs != nil {
		options = append(options, dqliteApp.WithCluster(*clusterAddrs))
	}

	dqliteDir, err := fileUtils.RequireDataDir(constant.DqliteDataDir)
	if err != nil {
		return nil, err
	}

	app, err := dqliteApp.New(dqliteDir, options...)
	if err != nil {
		return nil, err
	}

	err = app.Ready(context.Background())
	if err != nil {
		return nil, err
	}

	db, err := app.Open(context.Background(), constant.SharedDatabaseDSN)
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
		err = app.Handover(context.Background())
		if err != nil {
			log.Warn("Failed to handover dqlite: %v", err)
		}
		err = app.Close()
		if err != nil {
			log.Warn("Failed to close dqlite: %v", err)
		}
		log.Infof("Shared database closed")
	})

	return client, nil
}

func CreateSharedDatabase(serverAddr *[]string) *shared.Client {
	createOnce.Do(func() {
		var err error
		sharedClient, err = createSharedDatabase(serverAddr)
		if err != nil {
			log.Panicf("Failed to create shared database: %v", err)
		}
	})
	return sharedClient
}
