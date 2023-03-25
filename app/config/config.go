// Config load from local sqlite database, ensure all call to this package is after bootstrap LocalDb

package config

import (
	"context"

	"github.com/openPanel/core/app/db/db"
	"github.com/openPanel/core/app/generated/db/local"
	"github.com/openPanel/core/app/generated/db/local/kv"
)

func Load(key Key) (string, error) {
	v, err := db.GetLocalDb().KV.Query().
		Where(kv.Key(string(key))).
		Only(context.Background())
	if err != nil {
		if local.IsNotFound(err) {
			return "", nil
		} else {
			return "", err
		}
	}
	return v.Value, nil
}

func Save(key Key, value string) error {
	err := db.GetLocalDb().KV.
		Create().
		SetKey(string(key)).
		SetValue(value).
		OnConflictColumns(kv.FieldKey).
		UpdateValue().
		Exec(context.Background())
	return err
}
