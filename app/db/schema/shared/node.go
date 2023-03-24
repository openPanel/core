package shared

import (
	"entgo.io/ent"
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
		field.UUID("id", uuid.New()).Immutable(),
		field.String("name"),
		field.String("ip"),
		field.Int("port").Default(constant.DefaultListenPort),
		field.String("comment").Optional(),
	}
}

func (Node) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeCreateMixin{},
		mixin.TimeUpdateMixin{},
	}
}
