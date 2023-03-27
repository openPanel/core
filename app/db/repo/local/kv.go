package local

import (
	"context"
	"time"

	"github.com/openPanel/core/app/db/db"
	"github.com/openPanel/core/app/generated/db/local"
	"github.com/openPanel/core/app/generated/db/local/kv"
)

type kvRepo struct{}

var KVRepo = new(kvRepo)

func (r *kvRepo) Get(ctx context.Context, key string) (string, error) {
	v, err := db.GetLocalDb().KV.Query().
		Where(kv.Key(key)).
		Only(ctx)
	if err != nil {
		if local.IsNotFound(err) {
			return "", nil
		} else {
			return "", err
		}
	}

	if v.ExpiresAt.IsZero() && v.ExpiresAt.Before(time.Now()) {
		_, _ = db.GetLocalDb().KV.Delete().Where(kv.Key(key)).Exec(ctx)
		return "", nil
	}

	return v.Value, nil
}

func (r *kvRepo) BatchSet(ctx context.Context, values map[string]string) error {
	for k, v := range values {
		err := r.Set(ctx, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *kvRepo) Set(ctx context.Context, key string, value string) error {
	return r.SetExpire(ctx, key, value, time.Time{})
}

func (r *kvRepo) SetExpire(ctx context.Context, key string, value string, expiresAt time.Time) error {
	return db.GetLocalDb().KV.
		Create().
		SetKey(key).
		SetValue(value).
		SetExpiresAt(expiresAt).
		OnConflictColumns(kv.FieldKey).
		UpdateValue().
		Exec(ctx)
}
