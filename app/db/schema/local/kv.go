package local

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/openPanel/core/app/db/mixin"
)

type KV struct {
	ent.Schema
}

func (KV) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Unique(),
		field.String("value"),
		field.Time("expires_at").Optional(),
	}
}

func (KV) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeCreateMixin{},
		mixin.TimeUpdateMixin{},
	}
}

func (KV) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").Unique(),
	}
}
