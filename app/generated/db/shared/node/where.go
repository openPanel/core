// Code generated by ent, DO NOT EDIT.

package node

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/openPanel/core/app/generated/db/shared/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldName, v))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldIP, v))
}

// Port applies equality check predicate on the "port" field. It's identical to PortEQ.
func Port(v int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldPort, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldComment, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Node {
	return predicate.Node(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Node {
	return predicate.Node(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldName, v))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldIP, v))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldIP, v))
}

// PortEQ applies the EQ predicate on the "port" field.
func PortEQ(v int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldPort, v))
}

// PortNEQ applies the NEQ predicate on the "port" field.
func PortNEQ(v int) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldPort, v))
}

// PortIn applies the In predicate on the "port" field.
func PortIn(vs ...int) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldPort, vs...))
}

// PortNotIn applies the NotIn predicate on the "port" field.
func PortNotIn(vs ...int) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldPort, vs...))
}

// PortGT applies the GT predicate on the "port" field.
func PortGT(v int) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldPort, v))
}

// PortGTE applies the GTE predicate on the "port" field.
func PortGTE(v int) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldPort, v))
}

// PortLT applies the LT predicate on the "port" field.
func PortLT(v int) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldPort, v))
}

// PortLTE applies the LTE predicate on the "port" field.
func PortLTE(v int) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldPort, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldComment, v))
}

// CommentIsNil applies the IsNil predicate on the "comment" field.
func CommentIsNil() predicate.Node {
	return predicate.Node(sql.FieldIsNull(FieldComment))
}

// CommentNotNil applies the NotNil predicate on the "comment" field.
func CommentNotNil() predicate.Node {
	return predicate.Node(sql.FieldNotNull(FieldComment))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldComment, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		p(s.Not())
	})
}
