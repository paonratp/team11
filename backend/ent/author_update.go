// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/author"
	"github.com/team11/app/ent/book"
	"github.com/team11/app/ent/predicate"
	"github.com/team11/app/ent/research"
)

// AuthorUpdate is the builder for updating Author entities.
type AuthorUpdate struct {
	config
	hooks      []Hook
	mutation   *AuthorMutation
	predicates []predicate.Author
}

// Where adds a new predicate for the builder.
func (au *AuthorUpdate) Where(ps ...predicate.Author) *AuthorUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetName sets the Name field.
func (au *AuthorUpdate) SetName(s string) *AuthorUpdate {
	au.mutation.SetName(s)
	return au
}

// AddOwnerIDs adds the owner edge to Research by ids.
func (au *AuthorUpdate) AddOwnerIDs(ids ...int) *AuthorUpdate {
	au.mutation.AddOwnerIDs(ids...)
	return au
}

// AddOwner adds the owner edges to Research.
func (au *AuthorUpdate) AddOwner(r ...*Research) *AuthorUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.AddOwnerIDs(ids...)
}

// AddWriterIDs adds the writer edge to Book by ids.
func (au *AuthorUpdate) AddWriterIDs(ids ...int) *AuthorUpdate {
	au.mutation.AddWriterIDs(ids...)
	return au
}

// AddWriter adds the writer edges to Book.
func (au *AuthorUpdate) AddWriter(b ...*Book) *AuthorUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return au.AddWriterIDs(ids...)
}

// Mutation returns the AuthorMutation object of the builder.
func (au *AuthorUpdate) Mutation() *AuthorMutation {
	return au.mutation
}

// RemoveOwnerIDs removes the owner edge to Research by ids.
func (au *AuthorUpdate) RemoveOwnerIDs(ids ...int) *AuthorUpdate {
	au.mutation.RemoveOwnerIDs(ids...)
	return au
}

// RemoveOwner removes owner edges to Research.
func (au *AuthorUpdate) RemoveOwner(r ...*Research) *AuthorUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.RemoveOwnerIDs(ids...)
}

// RemoveWriterIDs removes the writer edge to Book by ids.
func (au *AuthorUpdate) RemoveWriterIDs(ids ...int) *AuthorUpdate {
	au.mutation.RemoveWriterIDs(ids...)
	return au
}

// RemoveWriter removes writer edges to Book.
func (au *AuthorUpdate) RemoveWriter(b ...*Book) *AuthorUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return au.RemoveWriterIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AuthorUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := au.mutation.Name(); ok {
		if err := author.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuthorUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuthorUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuthorUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AuthorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   author.Table,
			Columns: author.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: author.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: author.FieldName,
		})
	}
	if nodes := au.mutation.RemovedOwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := au.mutation.RemovedWriterIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.WriterIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{author.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AuthorUpdateOne is the builder for updating a single Author entity.
type AuthorUpdateOne struct {
	config
	hooks    []Hook
	mutation *AuthorMutation
}

// SetName sets the Name field.
func (auo *AuthorUpdateOne) SetName(s string) *AuthorUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// AddOwnerIDs adds the owner edge to Research by ids.
func (auo *AuthorUpdateOne) AddOwnerIDs(ids ...int) *AuthorUpdateOne {
	auo.mutation.AddOwnerIDs(ids...)
	return auo
}

// AddOwner adds the owner edges to Research.
func (auo *AuthorUpdateOne) AddOwner(r ...*Research) *AuthorUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.AddOwnerIDs(ids...)
}

// AddWriterIDs adds the writer edge to Book by ids.
func (auo *AuthorUpdateOne) AddWriterIDs(ids ...int) *AuthorUpdateOne {
	auo.mutation.AddWriterIDs(ids...)
	return auo
}

// AddWriter adds the writer edges to Book.
func (auo *AuthorUpdateOne) AddWriter(b ...*Book) *AuthorUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return auo.AddWriterIDs(ids...)
}

// Mutation returns the AuthorMutation object of the builder.
func (auo *AuthorUpdateOne) Mutation() *AuthorMutation {
	return auo.mutation
}

// RemoveOwnerIDs removes the owner edge to Research by ids.
func (auo *AuthorUpdateOne) RemoveOwnerIDs(ids ...int) *AuthorUpdateOne {
	auo.mutation.RemoveOwnerIDs(ids...)
	return auo
}

// RemoveOwner removes owner edges to Research.
func (auo *AuthorUpdateOne) RemoveOwner(r ...*Research) *AuthorUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.RemoveOwnerIDs(ids...)
}

// RemoveWriterIDs removes the writer edge to Book by ids.
func (auo *AuthorUpdateOne) RemoveWriterIDs(ids ...int) *AuthorUpdateOne {
	auo.mutation.RemoveWriterIDs(ids...)
	return auo
}

// RemoveWriter removes writer edges to Book.
func (auo *AuthorUpdateOne) RemoveWriter(b ...*Book) *AuthorUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return auo.RemoveWriterIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (auo *AuthorUpdateOne) Save(ctx context.Context) (*Author, error) {
	if v, ok := auo.mutation.Name(); ok {
		if err := author.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}

	var (
		err  error
		node *Author
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuthorUpdateOne) SaveX(ctx context.Context) *Author {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AuthorUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuthorUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AuthorUpdateOne) sqlSave(ctx context.Context) (a *Author, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   author.Table,
			Columns: author.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: author.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Author.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: author.FieldName,
		})
	}
	if nodes := auo.mutation.RemovedOwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := auo.mutation.RemovedWriterIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.WriterIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	a = &Author{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{author.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
