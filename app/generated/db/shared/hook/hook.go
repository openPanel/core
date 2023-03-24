// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/openPanel/core/app/generated/db/shared"
)

// The KVFunc type is an adapter to allow the use of ordinary
// function as KV mutator.
type KVFunc func(context.Context, *shared.KVMutation) (shared.Value, error)

// Mutate calls f(ctx, m).
func (f KVFunc) Mutate(ctx context.Context, m shared.Mutation) (shared.Value, error) {
	if mv, ok := m.(*shared.KVMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *shared.KVMutation", m)
}

// The NodeFunc type is an adapter to allow the use of ordinary
// function as Node mutator.
type NodeFunc func(context.Context, *shared.NodeMutation) (shared.Value, error)

// Mutate calls f(ctx, m).
func (f NodeFunc) Mutate(ctx context.Context, m shared.Mutation) (shared.Value, error) {
	if mv, ok := m.(*shared.NodeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *shared.NodeMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, shared.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m shared.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m shared.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m shared.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op shared.Op) Condition {
	return func(_ context.Context, m shared.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m shared.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m shared.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m shared.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk shared.Hook, cond Condition) shared.Hook {
	return func(next shared.Mutator) shared.Mutator {
		return shared.MutateFunc(func(ctx context.Context, m shared.Mutation) (shared.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, shared.Delete|shared.Create)
func On(hk shared.Hook, op shared.Op) shared.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, shared.Update|shared.UpdateOne)
func Unless(hk shared.Hook, op shared.Op) shared.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) shared.Hook {
	return func(shared.Mutator) shared.Mutator {
		return shared.MutateFunc(func(context.Context, shared.Mutation) (shared.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []shared.Hook {
//		return []shared.Hook{
//			Reject(shared.Delete|shared.Update),
//		}
//	}
func Reject(op shared.Op) shared.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []shared.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...shared.Hook) Chain {
	return Chain{append([]shared.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() shared.Hook {
	return func(mutator shared.Mutator) shared.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...shared.Hook) Chain {
	newHooks := make([]shared.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}