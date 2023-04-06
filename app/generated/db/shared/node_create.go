// Code generated by ent, DO NOT EDIT.

package shared

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openPanel/core/app/generated/db/shared/node"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	mutation *NodeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (nc *NodeCreate) SetCreatedAt(t time.Time) *NodeCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NodeCreate) SetNillableCreatedAt(t *time.Time) *NodeCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NodeCreate) SetUpdatedAt(t time.Time) *NodeCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NodeCreate) SetNillableUpdatedAt(t *time.Time) *NodeCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetName sets the "name" field.
func (nc *NodeCreate) SetName(s string) *NodeCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nc *NodeCreate) SetNillableName(s *string) *NodeCreate {
	if s != nil {
		nc.SetName(*s)
	}
	return nc
}

// SetIP sets the "ip" field.
func (nc *NodeCreate) SetIP(s string) *NodeCreate {
	nc.mutation.SetIP(s)
	return nc
}

// SetPort sets the "port" field.
func (nc *NodeCreate) SetPort(i int) *NodeCreate {
	nc.mutation.SetPort(i)
	return nc
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (nc *NodeCreate) SetNillablePort(i *int) *NodeCreate {
	if i != nil {
		nc.SetPort(*i)
	}
	return nc
}

// SetComment sets the "comment" field.
func (nc *NodeCreate) SetComment(s string) *NodeCreate {
	nc.mutation.SetComment(s)
	return nc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (nc *NodeCreate) SetNillableComment(s *string) *NodeCreate {
	if s != nil {
		nc.SetComment(*s)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NodeCreate) SetID(s string) *NodeCreate {
	nc.mutation.SetID(s)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NodeCreate) SetNillableID(s *string) *NodeCreate {
	if s != nil {
		nc.SetID(*s)
	}
	return nc
}

// Mutation returns the NodeMutation object of the builder.
func (nc *NodeCreate) Mutation() *NodeMutation {
	return nc.mutation
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	nc.defaults()
	return withHooks[*Node, NodeMutation](ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NodeCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NodeCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NodeCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := node.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := node.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.Port(); !ok {
		v := node.DefaultPort
		nc.mutation.SetPort(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := node.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NodeCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`shared: missing required field "Node.created_at"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`shared: missing required field "Node.updated_at"`)}
	}
	if _, ok := nc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`shared: missing required field "Node.ip"`)}
	}
	if _, ok := nc.mutation.Port(); !ok {
		return &ValidationError{Name: "port", err: errors.New(`shared: missing required field "Node.port"`)}
	}
	return nil
}

func (nc *NodeCreate) sqlSave(ctx context.Context) (*Node, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Node.ID type: %T", _spec.ID.Value)
		}
	}
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NodeCreate) createSpec() (*Node, *sqlgraph.CreateSpec) {
	var (
		_node = &Node{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(node.Table, sqlgraph.NewFieldSpec(node.FieldID, field.TypeString))
	)
	_spec.OnConflict = nc.conflict
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(node.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(node.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.Name(); ok {
		_spec.SetField(node.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := nc.mutation.IP(); ok {
		_spec.SetField(node.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := nc.mutation.Port(); ok {
		_spec.SetField(node.FieldPort, field.TypeInt, value)
		_node.Port = value
	}
	if value, ok := nc.mutation.Comment(); ok {
		_spec.SetField(node.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Node.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NodeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (nc *NodeCreate) OnConflict(opts ...sql.ConflictOption) *NodeUpsertOne {
	nc.conflict = opts
	return &NodeUpsertOne{
		create: nc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nc *NodeCreate) OnConflictColumns(columns ...string) *NodeUpsertOne {
	nc.conflict = append(nc.conflict, sql.ConflictColumns(columns...))
	return &NodeUpsertOne{
		create: nc,
	}
}

type (
	// NodeUpsertOne is the builder for "upsert"-ing
	//  one Node node.
	NodeUpsertOne struct {
		create *NodeCreate
	}

	// NodeUpsert is the "OnConflict" setter.
	NodeUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *NodeUpsert) SetUpdatedAt(v time.Time) *NodeUpsert {
	u.Set(node.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NodeUpsert) UpdateUpdatedAt() *NodeUpsert {
	u.SetExcluded(node.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *NodeUpsert) SetName(v string) *NodeUpsert {
	u.Set(node.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NodeUpsert) UpdateName() *NodeUpsert {
	u.SetExcluded(node.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *NodeUpsert) ClearName() *NodeUpsert {
	u.SetNull(node.FieldName)
	return u
}

// SetIP sets the "ip" field.
func (u *NodeUpsert) SetIP(v string) *NodeUpsert {
	u.Set(node.FieldIP, v)
	return u
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *NodeUpsert) UpdateIP() *NodeUpsert {
	u.SetExcluded(node.FieldIP)
	return u
}

// SetPort sets the "port" field.
func (u *NodeUpsert) SetPort(v int) *NodeUpsert {
	u.Set(node.FieldPort, v)
	return u
}

// UpdatePort sets the "port" field to the value that was provided on create.
func (u *NodeUpsert) UpdatePort() *NodeUpsert {
	u.SetExcluded(node.FieldPort)
	return u
}

// AddPort adds v to the "port" field.
func (u *NodeUpsert) AddPort(v int) *NodeUpsert {
	u.Add(node.FieldPort, v)
	return u
}

// SetComment sets the "comment" field.
func (u *NodeUpsert) SetComment(v string) *NodeUpsert {
	u.Set(node.FieldComment, v)
	return u
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *NodeUpsert) UpdateComment() *NodeUpsert {
	u.SetExcluded(node.FieldComment)
	return u
}

// ClearComment clears the value of the "comment" field.
func (u *NodeUpsert) ClearComment() *NodeUpsert {
	u.SetNull(node.FieldComment)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(node.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NodeUpsertOne) UpdateNewValues() *NodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(node.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(node.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Node.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NodeUpsertOne) Ignore() *NodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NodeUpsertOne) DoNothing() *NodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NodeCreate.OnConflict
// documentation for more info.
func (u *NodeUpsertOne) Update(set func(*NodeUpsert)) *NodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NodeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NodeUpsertOne) SetUpdatedAt(v time.Time) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateUpdatedAt() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *NodeUpsertOne) SetName(v string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateName() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *NodeUpsertOne) ClearName() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.ClearName()
	})
}

// SetIP sets the "ip" field.
func (u *NodeUpsertOne) SetIP(v string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateIP() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateIP()
	})
}

// SetPort sets the "port" field.
func (u *NodeUpsertOne) SetPort(v int) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetPort(v)
	})
}

// AddPort adds v to the "port" field.
func (u *NodeUpsertOne) AddPort(v int) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.AddPort(v)
	})
}

// UpdatePort sets the "port" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdatePort() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdatePort()
	})
}

// SetComment sets the "comment" field.
func (u *NodeUpsertOne) SetComment(v string) *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *NodeUpsertOne) UpdateComment() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateComment()
	})
}

// ClearComment clears the value of the "comment" field.
func (u *NodeUpsertOne) ClearComment() *NodeUpsertOne {
	return u.Update(func(s *NodeUpsert) {
		s.ClearComment()
	})
}

// Exec executes the query.
func (u *NodeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("shared: missing options for NodeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NodeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NodeUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("shared: NodeUpsertOne.ID is not supported by MySQL driver. Use NodeUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NodeUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NodeCreateBulk is the builder for creating many Node entities in bulk.
type NodeCreateBulk struct {
	config
	builders []*NodeCreate
	conflict []sql.ConflictOption
}

// Save creates the Node entities in the database.
func (ncb *NodeCreateBulk) Save(ctx context.Context) ([]*Node, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Node, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NodeCreateBulk) SaveX(ctx context.Context) []*Node {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NodeCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Node.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NodeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ncb *NodeCreateBulk) OnConflict(opts ...sql.ConflictOption) *NodeUpsertBulk {
	ncb.conflict = opts
	return &NodeUpsertBulk{
		create: ncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ncb *NodeCreateBulk) OnConflictColumns(columns ...string) *NodeUpsertBulk {
	ncb.conflict = append(ncb.conflict, sql.ConflictColumns(columns...))
	return &NodeUpsertBulk{
		create: ncb,
	}
}

// NodeUpsertBulk is the builder for "upsert"-ing
// a bulk of Node nodes.
type NodeUpsertBulk struct {
	create *NodeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(node.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NodeUpsertBulk) UpdateNewValues() *NodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(node.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(node.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Node.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NodeUpsertBulk) Ignore() *NodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NodeUpsertBulk) DoNothing() *NodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NodeCreateBulk.OnConflict
// documentation for more info.
func (u *NodeUpsertBulk) Update(set func(*NodeUpsert)) *NodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NodeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NodeUpsertBulk) SetUpdatedAt(v time.Time) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateUpdatedAt() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *NodeUpsertBulk) SetName(v string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateName() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *NodeUpsertBulk) ClearName() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.ClearName()
	})
}

// SetIP sets the "ip" field.
func (u *NodeUpsertBulk) SetIP(v string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateIP() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateIP()
	})
}

// SetPort sets the "port" field.
func (u *NodeUpsertBulk) SetPort(v int) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetPort(v)
	})
}

// AddPort adds v to the "port" field.
func (u *NodeUpsertBulk) AddPort(v int) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.AddPort(v)
	})
}

// UpdatePort sets the "port" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdatePort() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdatePort()
	})
}

// SetComment sets the "comment" field.
func (u *NodeUpsertBulk) SetComment(v string) *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *NodeUpsertBulk) UpdateComment() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.UpdateComment()
	})
}

// ClearComment clears the value of the "comment" field.
func (u *NodeUpsertBulk) ClearComment() *NodeUpsertBulk {
	return u.Update(func(s *NodeUpsert) {
		s.ClearComment()
	})
}

// Exec executes the query.
func (u *NodeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("shared: OnConflict was set for builder %d. Set it on the NodeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("shared: missing options for NodeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NodeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
