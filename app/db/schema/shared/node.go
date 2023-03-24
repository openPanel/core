package shared

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/db/mixin"
)

type Node struct {
	ent.Schema
}

func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Immutable().
			Annotations(entproto.Field(1)),
		field.String("name").
			Annotations(entproto.Field(2)),
		field.String("ip").
			Annotations(entproto.Field(3)),
		field.Int("port").
			Default(constant.DefaultListenPort).
			Annotations(entproto.Field(4)),
		field.String("comment").
			Optional().
			Annotations(entproto.Field(5)),
	}
}

func (Node) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeCreateMixin{},
		mixin.TimeUpdateMixin{},
	}
}

func (Node) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
