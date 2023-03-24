// Code generated by ent, DO NOT EDIT.

package shared

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openPanel/core/app/generated/db/shared/kv"
)

// KVCreate is the builder for creating a KV entity.
type KVCreate struct {
	config
	mutation *KVMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (kc *KVCreate) SetCreatedAt(t time.Time) *KVCreate {
	kc.mutation.SetCreatedAt(t)
	return kc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kc *KVCreate) SetNillableCreatedAt(t *time.Time) *KVCreate {
	if t != nil {
		kc.SetCreatedAt(*t)
	}
	return kc
}

// SetUpdatedAt sets the "updated_at" field.
func (kc *KVCreate) SetUpdatedAt(t time.Time) *KVCreate {
	kc.mutation.SetUpdatedAt(t)
	return kc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kc *KVCreate) SetNillableUpdatedAt(t *time.Time) *KVCreate {
	if t != nil {
		kc.SetUpdatedAt(*t)
	}
	return kc
}

// SetKey sets the "key" field.
func (kc *KVCreate) SetKey(s string) *KVCreate {
	kc.mutation.SetKey(s)
	return kc
}

// SetValue sets the "value" field.
func (kc *KVCreate) SetValue(s string) *KVCreate {
	kc.mutation.SetValue(s)
	return kc
}

// SetExpiresAt sets the "expires_at" field.
func (kc *KVCreate) SetExpiresAt(t time.Time) *KVCreate {
	kc.mutation.SetExpiresAt(t)
	return kc
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (kc *KVCreate) SetNillableExpiresAt(t *time.Time) *KVCreate {
	if t != nil {
		kc.SetExpiresAt(*t)
	}
	return kc
}

// Mutation returns the KVMutation object of the builder.
func (kc *KVCreate) Mutation() *KVMutation {
	return kc.mutation
}

// Save creates the KV in the database.
func (kc *KVCreate) Save(ctx context.Context) (*KV, error) {
	kc.defaults()
	return withHooks[*KV, KVMutation](ctx, kc.sqlSave, kc.mutation, kc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KVCreate) SaveX(ctx context.Context) *KV {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KVCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KVCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kc *KVCreate) defaults() {
	if _, ok := kc.mutation.CreatedAt(); !ok {
		v := kv.DefaultCreatedAt()
		kc.mutation.SetCreatedAt(v)
	}
	if _, ok := kc.mutation.UpdatedAt(); !ok {
		v := kv.DefaultUpdatedAt()
		kc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KVCreate) check() error {
	if _, ok := kc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`shared: missing required field "KV.created_at"`)}
	}
	if _, ok := kc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`shared: missing required field "KV.updated_at"`)}
	}
	if _, ok := kc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`shared: missing required field "KV.key"`)}
	}
	if _, ok := kc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`shared: missing required field "KV.value"`)}
	}
	return nil
}

func (kc *KVCreate) sqlSave(ctx context.Context) (*KV, error) {
	if err := kc.check(); err != nil {
		return nil, err
	}
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	kc.mutation.id = &_node.ID
	kc.mutation.done = true
	return _node, nil
}

func (kc *KVCreate) createSpec() (*KV, *sqlgraph.CreateSpec) {
	var (
		_node = &KV{config: kc.config}
		_spec = sqlgraph.NewCreateSpec(kv.Table, sqlgraph.NewFieldSpec(kv.FieldID, field.TypeInt))
	)
	_spec.OnConflict = kc.conflict
	if value, ok := kc.mutation.CreatedAt(); ok {
		_spec.SetField(kv.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := kc.mutation.UpdatedAt(); ok {
		_spec.SetField(kv.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := kc.mutation.Key(); ok {
		_spec.SetField(kv.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := kc.mutation.Value(); ok {
		_spec.SetField(kv.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := kc.mutation.ExpiresAt(); ok {
		_spec.SetField(kv.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.KV.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.KVUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (kc *KVCreate) OnConflict(opts ...sql.ConflictOption) *KVUpsertOne {
	kc.conflict = opts
	return &KVUpsertOne{
		create: kc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.KV.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (kc *KVCreate) OnConflictColumns(columns ...string) *KVUpsertOne {
	kc.conflict = append(kc.conflict, sql.ConflictColumns(columns...))
	return &KVUpsertOne{
		create: kc,
	}
}

type (
	// KVUpsertOne is the builder for "upsert"-ing
	//  one KV node.
	KVUpsertOne struct {
		create *KVCreate
	}

	// KVUpsert is the "OnConflict" setter.
	KVUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *KVUpsert) SetUpdatedAt(v time.Time) *KVUpsert {
	u.Set(kv.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *KVUpsert) UpdateUpdatedAt() *KVUpsert {
	u.SetExcluded(kv.FieldUpdatedAt)
	return u
}

// SetKey sets the "key" field.
func (u *KVUpsert) SetKey(v string) *KVUpsert {
	u.Set(kv.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *KVUpsert) UpdateKey() *KVUpsert {
	u.SetExcluded(kv.FieldKey)
	return u
}

// SetValue sets the "value" field.
func (u *KVUpsert) SetValue(v string) *KVUpsert {
	u.Set(kv.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *KVUpsert) UpdateValue() *KVUpsert {
	u.SetExcluded(kv.FieldValue)
	return u
}

// SetExpiresAt sets the "expires_at" field.
func (u *KVUpsert) SetExpiresAt(v time.Time) *KVUpsert {
	u.Set(kv.FieldExpiresAt, v)
	return u
}

// UpdateExpiresAt sets the "expires_at" field to the value that was provided on create.
func (u *KVUpsert) UpdateExpiresAt() *KVUpsert {
	u.SetExcluded(kv.FieldExpiresAt)
	return u
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (u *KVUpsert) ClearExpiresAt() *KVUpsert {
	u.SetNull(kv.FieldExpiresAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.KV.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *KVUpsertOne) UpdateNewValues() *KVUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(kv.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.KV.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *KVUpsertOne) Ignore() *KVUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *KVUpsertOne) DoNothing() *KVUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the KVCreate.OnConflict
// documentation for more info.
func (u *KVUpsertOne) Update(set func(*KVUpsert)) *KVUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&KVUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *KVUpsertOne) SetUpdatedAt(v time.Time) *KVUpsertOne {
	return u.Update(func(s *KVUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *KVUpsertOne) UpdateUpdatedAt() *KVUpsertOne {
	return u.Update(func(s *KVUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetKey sets the "key" field.
func (u *KVUpsertOne) SetKey(v string) *KVUpsertOne {
	return u.Update(func(s *KVUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *KVUpsertOne) UpdateKey() *KVUpsertOne {
	return u.Update(func(s *KVUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *KVUpsertOne) SetValue(v string) *KVUpsertOne {
	return u.Update(func(s *KVUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *KVUpsertOne) UpdateValue() *KVUpsertOne {
	return u.Update(func(s *KVUpsert) {
		s.UpdateValue()
	})
}

// SetExpiresAt sets the "expires_at" field.
func (u *KVUpsertOne) SetExpiresAt(v time.Time) *KVUpsertOne {
	return u.Update(func(s *KVUpsert) {
		s.SetExpiresAt(v)
	})
}

// UpdateExpiresAt sets the "expires_at" field to the value that was provided on create.
func (u *KVUpsertOne) UpdateExpiresAt() *KVUpsertOne {
	return u.Update(func(s *KVUpsert) {
		s.UpdateExpiresAt()
	})
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (u *KVUpsertOne) ClearExpiresAt() *KVUpsertOne {
	return u.Update(func(s *KVUpsert) {
		s.ClearExpiresAt()
	})
}

// Exec executes the query.
func (u *KVUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("shared: missing options for KVCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *KVUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *KVUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *KVUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// KVCreateBulk is the builder for creating many KV entities in bulk.
type KVCreateBulk struct {
	config
	builders []*KVCreate
	conflict []sql.ConflictOption
}

// Save creates the KV entities in the database.
func (kcb *KVCreateBulk) Save(ctx context.Context) ([]*KV, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*KV, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KVMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = kcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KVCreateBulk) SaveX(ctx context.Context) []*KV {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KVCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KVCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.KV.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.KVUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (kcb *KVCreateBulk) OnConflict(opts ...sql.ConflictOption) *KVUpsertBulk {
	kcb.conflict = opts
	return &KVUpsertBulk{
		create: kcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.KV.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (kcb *KVCreateBulk) OnConflictColumns(columns ...string) *KVUpsertBulk {
	kcb.conflict = append(kcb.conflict, sql.ConflictColumns(columns...))
	return &KVUpsertBulk{
		create: kcb,
	}
}

// KVUpsertBulk is the builder for "upsert"-ing
// a bulk of KV nodes.
type KVUpsertBulk struct {
	create *KVCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.KV.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *KVUpsertBulk) UpdateNewValues() *KVUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(kv.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.KV.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *KVUpsertBulk) Ignore() *KVUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *KVUpsertBulk) DoNothing() *KVUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the KVCreateBulk.OnConflict
// documentation for more info.
func (u *KVUpsertBulk) Update(set func(*KVUpsert)) *KVUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&KVUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *KVUpsertBulk) SetUpdatedAt(v time.Time) *KVUpsertBulk {
	return u.Update(func(s *KVUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *KVUpsertBulk) UpdateUpdatedAt() *KVUpsertBulk {
	return u.Update(func(s *KVUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetKey sets the "key" field.
func (u *KVUpsertBulk) SetKey(v string) *KVUpsertBulk {
	return u.Update(func(s *KVUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *KVUpsertBulk) UpdateKey() *KVUpsertBulk {
	return u.Update(func(s *KVUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *KVUpsertBulk) SetValue(v string) *KVUpsertBulk {
	return u.Update(func(s *KVUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *KVUpsertBulk) UpdateValue() *KVUpsertBulk {
	return u.Update(func(s *KVUpsert) {
		s.UpdateValue()
	})
}

// SetExpiresAt sets the "expires_at" field.
func (u *KVUpsertBulk) SetExpiresAt(v time.Time) *KVUpsertBulk {
	return u.Update(func(s *KVUpsert) {
		s.SetExpiresAt(v)
	})
}

// UpdateExpiresAt sets the "expires_at" field to the value that was provided on create.
func (u *KVUpsertBulk) UpdateExpiresAt() *KVUpsertBulk {
	return u.Update(func(s *KVUpsert) {
		s.UpdateExpiresAt()
	})
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (u *KVUpsertBulk) ClearExpiresAt() *KVUpsertBulk {
	return u.Update(func(s *KVUpsert) {
		s.ClearExpiresAt()
	})
}

// Exec executes the query.
func (u *KVUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("shared: OnConflict was set for builder %d. Set it on the KVCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("shared: missing options for KVCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *KVUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
