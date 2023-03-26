package shared

import (
	"context"
	"time"

	"github.com/openPanel/core/app/db/db"
	"github.com/openPanel/core/app/generated/db/shared/kv"
)

type kvRepo struct{}

var KVRepo = new(kvRepo)

func (r *kvRepo) Get(ctx context.Context, key string) string {
	v, err := db.GetSharedDb().KV.Query().
		Where(kv.Key(key)).
		Only(ctx)
	if err != nil {
		return ""
	}

	if v.ExpiresAt.IsZero() && v.ExpiresAt.Before(time.Now()) {
		_, _ = db.GetSharedDb().KV.Delete().Where(kv.Key(key)).Exec(ctx)
		return ""
	}

	return v.Value
}

func (r *kvRepo) Set(ctx context.Context, key string, value string) error {
	return r.SetExpire(ctx, key, value, time.Time{})
}

func (r *kvRepo) SetExpire(ctx context.Context, key string, value string, expiresAt time.Time) error {
	return db.GetSharedDb().KV.
		Create().
		SetKey(key).
		SetValue(value).
		SetExpiresAt(expiresAt).
		OnConflictColumns(kv.FieldKey).
		UpdateValue().
		Exec(ctx)
}
