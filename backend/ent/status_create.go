// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/book"
	"github.com/team11/app/ent/bookborrow"
	"github.com/team11/app/ent/cliententity"
	"github.com/team11/app/ent/status"
)

// StatusCreate is the builder for creating a Status entity.
type StatusCreate struct {
	config
	mutation *StatusMutation
	hooks    []Hook
}

// SetSTATUSNAME sets the STATUS_NAME field.
func (sc *StatusCreate) SetSTATUSNAME(s string) *StatusCreate {
	sc.mutation.SetSTATUSNAME(s)
	return sc
}

// AddStatuIDs adds the status edge to ClientEntity by ids.
func (sc *StatusCreate) AddStatuIDs(ids ...int) *StatusCreate {
	sc.mutation.AddStatuIDs(ids...)
	return sc
}

// AddStatus adds the status edges to ClientEntity.
func (sc *StatusCreate) AddStatus(c ...*ClientEntity) *StatusCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddStatuIDs(ids...)
}

// AddStatusofbookIDs adds the statusofbook edge to Book by ids.
func (sc *StatusCreate) AddStatusofbookIDs(ids ...int) *StatusCreate {
	sc.mutation.AddStatusofbookIDs(ids...)
	return sc
}

// AddStatusofbook adds the statusofbook edges to Book.
func (sc *StatusCreate) AddStatusofbook(b ...*Book) *StatusCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return sc.AddStatusofbookIDs(ids...)
}

// AddStatusbookborrowIDs adds the statusbookborrow edge to Bookborrow by ids.
func (sc *StatusCreate) AddStatusbookborrowIDs(ids ...int) *StatusCreate {
	sc.mutation.AddStatusbookborrowIDs(ids...)
	return sc
}

// AddStatusbookborrow adds the statusbookborrow edges to Bookborrow.
func (sc *StatusCreate) AddStatusbookborrow(b ...*Bookborrow) *StatusCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return sc.AddStatusbookborrowIDs(ids...)
}

// Mutation returns the StatusMutation object of the builder.
func (sc *StatusCreate) Mutation() *StatusMutation {
	return sc.mutation
}

// Save creates the Status in the database.
func (sc *StatusCreate) Save(ctx context.Context) (*Status, error) {
	if _, ok := sc.mutation.STATUSNAME(); !ok {
		return nil, &ValidationError{Name: "STATUS_NAME", err: errors.New("ent: missing required field \"STATUS_NAME\"")}
	}
	if v, ok := sc.mutation.STATUSNAME(); ok {
		if err := status.STATUSNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "STATUS_NAME", err: fmt.Errorf("ent: validator failed for field \"STATUS_NAME\": %w", err)}
		}
	}
	var (
		err  error
		node *Status
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StatusCreate) SaveX(ctx context.Context) *Status {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *StatusCreate) sqlSave(ctx context.Context) (*Status, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *StatusCreate) createSpec() (*Status, *sqlgraph.CreateSpec) {
	var (
		s     = &Status{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: status.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: status.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.STATUSNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: status.FieldSTATUSNAME,
		})
		s.STATUSNAME = value
	}
	if nodes := sc.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   status.StatusTable,
			Columns: []string{status.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cliententity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StatusofbookIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   status.StatusofbookTable,
			Columns: []string{status.StatusofbookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StatusbookborrowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   status.StatusbookborrowTable,
			Columns: []string{status.StatusbookborrowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookborrow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
