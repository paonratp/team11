// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/author"
	"github.com/team11/app/ent/book"
	"github.com/team11/app/ent/research"
)

// AuthorCreate is the builder for creating a Author entity.
type AuthorCreate struct {
	config
	mutation *AuthorMutation
	hooks    []Hook
}

// SetName sets the Name field.
func (ac *AuthorCreate) SetName(s string) *AuthorCreate {
	ac.mutation.SetName(s)
	return ac
}

// AddOwnerIDs adds the owner edge to Research by ids.
func (ac *AuthorCreate) AddOwnerIDs(ids ...int) *AuthorCreate {
	ac.mutation.AddOwnerIDs(ids...)
	return ac
}

// AddOwner adds the owner edges to Research.
func (ac *AuthorCreate) AddOwner(r ...*Research) *AuthorCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ac.AddOwnerIDs(ids...)
}

// AddWriterIDs adds the writer edge to Book by ids.
func (ac *AuthorCreate) AddWriterIDs(ids ...int) *AuthorCreate {
	ac.mutation.AddWriterIDs(ids...)
	return ac
}

// AddWriter adds the writer edges to Book.
func (ac *AuthorCreate) AddWriter(b ...*Book) *AuthorCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ac.AddWriterIDs(ids...)
}

// Mutation returns the AuthorMutation object of the builder.
func (ac *AuthorCreate) Mutation() *AuthorMutation {
	return ac.mutation
}

// Save creates the Author in the database.
func (ac *AuthorCreate) Save(ctx context.Context) (*Author, error) {
	if _, ok := ac.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "Name", err: errors.New("ent: missing required field \"Name\"")}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := author.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}
	var (
		err  error
		node *Author
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthorCreate) SaveX(ctx context.Context) *Author {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AuthorCreate) sqlSave(ctx context.Context) (*Author, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *AuthorCreate) createSpec() (*Author, *sqlgraph.CreateSpec) {
	var (
		a     = &Author{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: author.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: author.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: author.FieldName,
		})
		a.Name = value
	}
	if nodes := ac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   author.OwnerTable,
			Columns: []string{author.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: research.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.WriterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   author.WriterTable,
			Columns: []string{author.WriterColumn},
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
	return a, _spec
}
