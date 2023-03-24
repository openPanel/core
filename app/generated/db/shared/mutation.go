// Code generated by ent, DO NOT EDIT.

package shared

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/openPanel/core/app/generated/db/shared/kv"
	"github.com/openPanel/core/app/generated/db/shared/node"
	"github.com/openPanel/core/app/generated/db/shared/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeKV   = "KV"
	TypeNode = "Node"
)

// KVMutation represents an operation that mutates the KV nodes in the graph.
type KVMutation struct {
	config
	op            Op
	typ           string
	id            *int
	created_at    *time.Time
	updated_at    *time.Time
	key           *string
	value         *string
	expires_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*KV, error)
	predicates    []predicate.KV
}

var _ ent.Mutation = (*KVMutation)(nil)

// kvOption allows management of the mutation configuration using functional options.
type kvOption func(*KVMutation)

// newKVMutation creates new mutation for the KV entity.
func newKVMutation(c config, op Op, opts ...kvOption) *KVMutation {
	m := &KVMutation{
		config:        c,
		op:            op,
		typ:           TypeKV,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withKVID sets the ID field of the mutation.
func withKVID(id int) kvOption {
	return func(m *KVMutation) {
		var (
			err   error
			once  sync.Once
			value *KV
		)
		m.oldValue = func(ctx context.Context) (*KV, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().KV.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withKV sets the old KV of the mutation.
func withKV(node *KV) kvOption {
	return func(m *KVMutation) {
		m.oldValue = func(context.Context) (*KV, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m KVMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m KVMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("shared: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *KVMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *KVMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().KV.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *KVMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *KVMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the KV entity.
// If the KV object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KVMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *KVMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *KVMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *KVMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the KV entity.
// If the KV object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KVMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *KVMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetKey sets the "key" field.
func (m *KVMutation) SetKey(s string) {
	m.key = &s
}

// Key returns the value of the "key" field in the mutation.
func (m *KVMutation) Key() (r string, exists bool) {
	v := m.key
	if v == nil {
		return
	}
	return *v, true
}

// OldKey returns the old "key" field's value of the KV entity.
// If the KV object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KVMutation) OldKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKey: %w", err)
	}
	return oldValue.Key, nil
}

// ResetKey resets all changes to the "key" field.
func (m *KVMutation) ResetKey() {
	m.key = nil
}

// SetValue sets the "value" field.
func (m *KVMutation) SetValue(s string) {
	m.value = &s
}

// Value returns the value of the "value" field in the mutation.
func (m *KVMutation) Value() (r string, exists bool) {
	v := m.value
	if v == nil {
		return
	}
	return *v, true
}

// OldValue returns the old "value" field's value of the KV entity.
// If the KV object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KVMutation) OldValue(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldValue is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldValue requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldValue: %w", err)
	}
	return oldValue.Value, nil
}

// ResetValue resets all changes to the "value" field.
func (m *KVMutation) ResetValue() {
	m.value = nil
}

// SetExpiresAt sets the "expires_at" field.
func (m *KVMutation) SetExpiresAt(t time.Time) {
	m.expires_at = &t
}

// ExpiresAt returns the value of the "expires_at" field in the mutation.
func (m *KVMutation) ExpiresAt() (r time.Time, exists bool) {
	v := m.expires_at
	if v == nil {
		return
	}
	return *v, true
}

// OldExpiresAt returns the old "expires_at" field's value of the KV entity.
// If the KV object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KVMutation) OldExpiresAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExpiresAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExpiresAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExpiresAt: %w", err)
	}
	return oldValue.ExpiresAt, nil
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (m *KVMutation) ClearExpiresAt() {
	m.expires_at = nil
	m.clearedFields[kv.FieldExpiresAt] = struct{}{}
}

// ExpiresAtCleared returns if the "expires_at" field was cleared in this mutation.
func (m *KVMutation) ExpiresAtCleared() bool {
	_, ok := m.clearedFields[kv.FieldExpiresAt]
	return ok
}

// ResetExpiresAt resets all changes to the "expires_at" field.
func (m *KVMutation) ResetExpiresAt() {
	m.expires_at = nil
	delete(m.clearedFields, kv.FieldExpiresAt)
}

// Where appends a list predicates to the KVMutation builder.
func (m *KVMutation) Where(ps ...predicate.KV) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the KVMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *KVMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.KV, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *KVMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *KVMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (KV).
func (m *KVMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *KVMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, kv.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, kv.FieldUpdatedAt)
	}
	if m.key != nil {
		fields = append(fields, kv.FieldKey)
	}
	if m.value != nil {
		fields = append(fields, kv.FieldValue)
	}
	if m.expires_at != nil {
		fields = append(fields, kv.FieldExpiresAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *KVMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case kv.FieldCreatedAt:
		return m.CreatedAt()
	case kv.FieldUpdatedAt:
		return m.UpdatedAt()
	case kv.FieldKey:
		return m.Key()
	case kv.FieldValue:
		return m.Value()
	case kv.FieldExpiresAt:
		return m.ExpiresAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *KVMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case kv.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case kv.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case kv.FieldKey:
		return m.OldKey(ctx)
	case kv.FieldValue:
		return m.OldValue(ctx)
	case kv.FieldExpiresAt:
		return m.OldExpiresAt(ctx)
	}
	return nil, fmt.Errorf("unknown KV field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *KVMutation) SetField(name string, value ent.Value) error {
	switch name {
	case kv.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case kv.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case kv.FieldKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKey(v)
		return nil
	case kv.FieldValue:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetValue(v)
		return nil
	case kv.FieldExpiresAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpiresAt(v)
		return nil
	}
	return fmt.Errorf("unknown KV field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *KVMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *KVMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *KVMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown KV numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *KVMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(kv.FieldExpiresAt) {
		fields = append(fields, kv.FieldExpiresAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *KVMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *KVMutation) ClearField(name string) error {
	switch name {
	case kv.FieldExpiresAt:
		m.ClearExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown KV nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *KVMutation) ResetField(name string) error {
	switch name {
	case kv.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case kv.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case kv.FieldKey:
		m.ResetKey()
		return nil
	case kv.FieldValue:
		m.ResetValue()
		return nil
	case kv.FieldExpiresAt:
		m.ResetExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown KV field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *KVMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *KVMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *KVMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *KVMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *KVMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *KVMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *KVMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown KV unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *KVMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown KV edge %s", name)
}

// NodeMutation represents an operation that mutates the Node nodes in the graph.
type NodeMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	name          *string
	ip            *string
	port          *int
	addport       *int
	comment       *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Node, error)
	predicates    []predicate.Node
}

var _ ent.Mutation = (*NodeMutation)(nil)

// nodeOption allows management of the mutation configuration using functional options.
type nodeOption func(*NodeMutation)

// newNodeMutation creates new mutation for the Node entity.
func newNodeMutation(c config, op Op, opts ...nodeOption) *NodeMutation {
	m := &NodeMutation{
		config:        c,
		op:            op,
		typ:           TypeNode,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withNodeID sets the ID field of the mutation.
func withNodeID(id uuid.UUID) nodeOption {
	return func(m *NodeMutation) {
		var (
			err   error
			once  sync.Once
			value *Node
		)
		m.oldValue = func(ctx context.Context) (*Node, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Node.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withNode sets the old Node of the mutation.
func withNode(node *Node) nodeOption {
	return func(m *NodeMutation) {
		m.oldValue = func(context.Context) (*Node, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m NodeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m NodeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("shared: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Node entities.
func (m *NodeMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *NodeMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *NodeMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Node.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *NodeMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *NodeMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Node entity.
// If the Node object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NodeMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *NodeMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *NodeMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *NodeMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Node entity.
// If the Node object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NodeMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *NodeMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *NodeMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *NodeMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Node entity.
// If the Node object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NodeMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *NodeMutation) ResetName() {
	m.name = nil
}

// SetIP sets the "ip" field.
func (m *NodeMutation) SetIP(s string) {
	m.ip = &s
}

// IP returns the value of the "ip" field in the mutation.
func (m *NodeMutation) IP() (r string, exists bool) {
	v := m.ip
	if v == nil {
		return
	}
	return *v, true
}

// OldIP returns the old "ip" field's value of the Node entity.
// If the Node object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NodeMutation) OldIP(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIP is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIP requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIP: %w", err)
	}
	return oldValue.IP, nil
}

// ResetIP resets all changes to the "ip" field.
func (m *NodeMutation) ResetIP() {
	m.ip = nil
}

// SetPort sets the "port" field.
func (m *NodeMutation) SetPort(i int) {
	m.port = &i
	m.addport = nil
}

// Port returns the value of the "port" field in the mutation.
func (m *NodeMutation) Port() (r int, exists bool) {
	v := m.port
	if v == nil {
		return
	}
	return *v, true
}

// OldPort returns the old "port" field's value of the Node entity.
// If the Node object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NodeMutation) OldPort(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPort is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPort requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPort: %w", err)
	}
	return oldValue.Port, nil
}

// AddPort adds i to the "port" field.
func (m *NodeMutation) AddPort(i int) {
	if m.addport != nil {
		*m.addport += i
	} else {
		m.addport = &i
	}
}

// AddedPort returns the value that was added to the "port" field in this mutation.
func (m *NodeMutation) AddedPort() (r int, exists bool) {
	v := m.addport
	if v == nil {
		return
	}
	return *v, true
}

// ResetPort resets all changes to the "port" field.
func (m *NodeMutation) ResetPort() {
	m.port = nil
	m.addport = nil
}

// SetComment sets the "comment" field.
func (m *NodeMutation) SetComment(s string) {
	m.comment = &s
}

// Comment returns the value of the "comment" field in the mutation.
func (m *NodeMutation) Comment() (r string, exists bool) {
	v := m.comment
	if v == nil {
		return
	}
	return *v, true
}

// OldComment returns the old "comment" field's value of the Node entity.
// If the Node object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NodeMutation) OldComment(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldComment is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldComment requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldComment: %w", err)
	}
	return oldValue.Comment, nil
}

// ClearComment clears the value of the "comment" field.
func (m *NodeMutation) ClearComment() {
	m.comment = nil
	m.clearedFields[node.FieldComment] = struct{}{}
}

// CommentCleared returns if the "comment" field was cleared in this mutation.
func (m *NodeMutation) CommentCleared() bool {
	_, ok := m.clearedFields[node.FieldComment]
	return ok
}

// ResetComment resets all changes to the "comment" field.
func (m *NodeMutation) ResetComment() {
	m.comment = nil
	delete(m.clearedFields, node.FieldComment)
}

// Where appends a list predicates to the NodeMutation builder.
func (m *NodeMutation) Where(ps ...predicate.Node) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the NodeMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *NodeMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Node, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *NodeMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *NodeMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Node).
func (m *NodeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *NodeMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.created_at != nil {
		fields = append(fields, node.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, node.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, node.FieldName)
	}
	if m.ip != nil {
		fields = append(fields, node.FieldIP)
	}
	if m.port != nil {
		fields = append(fields, node.FieldPort)
	}
	if m.comment != nil {
		fields = append(fields, node.FieldComment)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *NodeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case node.FieldCreatedAt:
		return m.CreatedAt()
	case node.FieldUpdatedAt:
		return m.UpdatedAt()
	case node.FieldName:
		return m.Name()
	case node.FieldIP:
		return m.IP()
	case node.FieldPort:
		return m.Port()
	case node.FieldComment:
		return m.Comment()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *NodeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case node.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case node.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case node.FieldName:
		return m.OldName(ctx)
	case node.FieldIP:
		return m.OldIP(ctx)
	case node.FieldPort:
		return m.OldPort(ctx)
	case node.FieldComment:
		return m.OldComment(ctx)
	}
	return nil, fmt.Errorf("unknown Node field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NodeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case node.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case node.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case node.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case node.FieldIP:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIP(v)
		return nil
	case node.FieldPort:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPort(v)
		return nil
	case node.FieldComment:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetComment(v)
		return nil
	}
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *NodeMutation) AddedFields() []string {
	var fields []string
	if m.addport != nil {
		fields = append(fields, node.FieldPort)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *NodeMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case node.FieldPort:
		return m.AddedPort()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NodeMutation) AddField(name string, value ent.Value) error {
	switch name {
	case node.FieldPort:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPort(v)
		return nil
	}
	return fmt.Errorf("unknown Node numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *NodeMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(node.FieldComment) {
		fields = append(fields, node.FieldComment)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *NodeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *NodeMutation) ClearField(name string) error {
	switch name {
	case node.FieldComment:
		m.ClearComment()
		return nil
	}
	return fmt.Errorf("unknown Node nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *NodeMutation) ResetField(name string) error {
	switch name {
	case node.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case node.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case node.FieldName:
		m.ResetName()
		return nil
	case node.FieldIP:
		m.ResetIP()
		return nil
	case node.FieldPort:
		m.ResetPort()
		return nil
	case node.FieldComment:
		m.ResetComment()
		return nil
	}
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *NodeMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *NodeMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *NodeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *NodeMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *NodeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *NodeMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *NodeMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Node unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *NodeMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Node edge %s", name)
}
