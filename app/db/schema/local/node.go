package local

// just use as a local copy of cluster version, may not latest, should not depend on this.

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/openPanel/core/app/db/mixin"
)

type Node struct {
	ent.Schema
}

func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("name"),
		field.String("ip"),
		field.Int("port"),
		field.String("comment"),
	}
}

func (Node) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeCreateMixin{},
	}
}
