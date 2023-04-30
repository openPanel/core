// Config load from local sqlite database, ensure all call to this package is after bootstrap LocalDb

package config

import (
	"context"
	"encoding/json"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/db/repo/local"
	"github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/router"
)

func init() {
	// just prevent import cycle
	router.NodePersistence = func(nodes []router.Node) {
		err := UpdateNodesCache(nodes)
		if err != nil {
			log.Warnf("Failed to update nodes cache: %v", err)
		}
	}
}

func Load(key constant.ConfigKey, value any, store constant.Store) error {
	var v string
	var err error

	if store == constant.SharedStore {
		v, err = shared.KVRepo.Get(context.Background(), string(key))
	} else {
		v, err = local.KVRepo.Get(context.Background(), string(key))
	}

	if err != nil {
		value = nil
		return err
	}
	err = json.Unmarshal([]byte(v), value)
	if err != nil {
		log.Debugf("Error unmarshalling config key %s: %s", key, err.Error())
		log.Debugf("Value: %s", v)
	}

	return err
}

func Save(key constant.ConfigKey, value any, store constant.Store) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if store == constant.SharedStore {
		return shared.KVRepo.Set(context.Background(), string(key), string(v))
	} else {
		return local.KVRepo.Set(context.Background(), string(key), string(v))
	}
}
