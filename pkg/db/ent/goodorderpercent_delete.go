// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/goodorderpercent"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
)

// GoodOrderPercentDelete is the builder for deleting a GoodOrderPercent entity.
type GoodOrderPercentDelete struct {
	config
	hooks    []Hook
	mutation *GoodOrderPercentMutation
}

// Where appends a list predicates to the GoodOrderPercentDelete builder.
func (gopd *GoodOrderPercentDelete) Where(ps ...predicate.GoodOrderPercent) *GoodOrderPercentDelete {
	gopd.mutation.Where(ps...)
	return gopd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gopd *GoodOrderPercentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gopd.hooks) == 0 {
		affected, err = gopd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodOrderPercentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gopd.mutation = mutation
			affected, err = gopd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gopd.hooks) - 1; i >= 0; i-- {
			if gopd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gopd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gopd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gopd *GoodOrderPercentDelete) ExecX(ctx context.Context) int {
	n, err := gopd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gopd *GoodOrderPercentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: goodorderpercent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodorderpercent.FieldID,
			},
		},
	}
	if ps := gopd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gopd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GoodOrderPercentDeleteOne is the builder for deleting a single GoodOrderPercent entity.
type GoodOrderPercentDeleteOne struct {
	gopd *GoodOrderPercentDelete
}

// Exec executes the deletion query.
func (gopdo *GoodOrderPercentDeleteOne) Exec(ctx context.Context) error {
	n, err := gopdo.gopd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodorderpercent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gopdo *GoodOrderPercentDeleteOne) ExecX(ctx context.Context) {
	gopdo.gopd.ExecX(ctx)
}
