package mixin

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/openPanel/core/app/constant"
)

type TimeCreateMixin struct {
	mixin.Schema
}

func (TimeCreateMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(entproto.Field(constant.ProtoMixinIdStart)),
	}
}

type TimeUpdateMixin struct {
	mixin.Schema
}

func (TimeUpdateMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entproto.Field(constant.ProtoMixinIdStart + 1)),
	}
}
