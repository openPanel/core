package shared

import (
	"context"
	"time"

	"github.com/openPanel/core/app/db/db"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/generated/db/shared/kv"
)

type kvRepo struct{}

var KVRepo = new(kvRepo)

func (r *kvRepo) Get(ctx context.Context, key string) (string, error) {
	v, err := db.GetSharedDb().KV.Query().
		Where(kv.Key(key)).
		Only(ctx)
	if err != nil {
		if shared.IsNotFound(err) {
			return "", nil
		} else {
			return "", err
		}
	}

	if v.ExpiresAt.IsZero() && v.ExpiresAt.Before(time.Now()) {
		_, _ = db.GetSharedDb().KV.Delete().Where(kv.Key(key)).Exec(ctx)
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
	_, err := db.GetSharedDb().KV.
		Create().
		SetKey(key).
		SetValue(value).
		SetExpiresAt(expiresAt).Save(ctx)
	return err
}
