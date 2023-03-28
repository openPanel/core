// Config load from local sqlite database, ensure all call to this package is after bootstrap LocalDb

package config

import (
	"context"
	"encoding/json"

	"github.com/openPanel/core/app/db/repo/local"
)

func Load(key Key, value any) error {
	v, err := local.KVRepo.Get(context.Background(), string(key))
	if err != nil {
		value = nil
		return err
	}
	err = json.Unmarshal([]byte(v), &value)
	return err
}

func Save(key Key, value any) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return local.KVRepo.Set(context.Background(), string(key), string(v))
}
