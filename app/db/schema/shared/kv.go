package shared

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/openPanel/core/app/db/mixin"
)

type KV struct {
	ent.Schema
}

func (KV) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").
			Unique().
			Annotations(entproto.Field(2)),
		field.String("value").
			Annotations(entproto.Field(3)),
		field.Time("expires_at").
			Optional().
			Annotations(entproto.Field(4)),
	}
}

func (KV) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeCreateMixin{},
		mixin.TimeUpdateMixin{},
	}
}

func (KV) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
