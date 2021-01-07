// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/bookreturn"
)

// BookreturnCreate is the builder for creating a Bookreturn entity.
type BookreturnCreate struct {
	config
	mutation *BookreturnMutation
	hooks    []Hook
}

// SetBookName sets the book_name field.
func (bc *BookreturnCreate) SetBookName(s string) *BookreturnCreate {
	bc.mutation.SetBookName(s)
	return bc
}

// Mutation returns the BookreturnMutation object of the builder.
func (bc *BookreturnCreate) Mutation() *BookreturnMutation {
	return bc.mutation
}

// Save creates the Bookreturn in the database.
func (bc *BookreturnCreate) Save(ctx context.Context) (*Bookreturn, error) {
	if _, ok := bc.mutation.BookName(); !ok {
		return nil, &ValidationError{Name: "book_name", err: errors.New("ent: missing required field \"book_name\"")}
	}
	if v, ok := bc.mutation.BookName(); ok {
		if err := bookreturn.BookNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "book_name", err: fmt.Errorf("ent: validator failed for field \"book_name\": %w", err)}
		}
	}
	var (
		err  error
		node *Bookreturn
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookreturnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookreturnCreate) SaveX(ctx context.Context) *Bookreturn {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BookreturnCreate) sqlSave(ctx context.Context) (*Bookreturn, error) {
	b, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	b.ID = int(id)
	return b, nil
}

func (bc *BookreturnCreate) createSpec() (*Bookreturn, *sqlgraph.CreateSpec) {
	var (
		b     = &Bookreturn{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bookreturn.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookreturn.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.BookName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bookreturn.FieldBookName,
		})
		b.BookName = value
	}
	return b, _spec
}