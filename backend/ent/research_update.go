// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/author"
	"github.com/team11/app/ent/predicate"
	"github.com/team11/app/ent/research"
	"github.com/team11/app/ent/researchtype"
	"github.com/team11/app/ent/user"
)

// ResearchUpdate is the builder for updating Research entities.
type ResearchUpdate struct {
	config
	hooks      []Hook
	mutation   *ResearchMutation
	predicates []predicate.Research
}

// Where adds a new predicate for the builder.
func (ru *ResearchUpdate) Where(ps ...predicate.Research) *ResearchUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetDOCNAME sets the DOC_NAME field.
func (ru *ResearchUpdate) SetDOCNAME(s string) *ResearchUpdate {
	ru.mutation.SetDOCNAME(s)
	return ru
}

// SetDATE sets the DATE field.
func (ru *ResearchUpdate) SetDATE(t time.Time) *ResearchUpdate {
	ru.mutation.SetDATE(t)
	return ru
}

// SetNillableDATE sets the DATE field if the given value is not nil.
func (ru *ResearchUpdate) SetNillableDATE(t *time.Time) *ResearchUpdate {
	if t != nil {
		ru.SetDATE(*t)
	}
	return ru
}

// SetRegisterID sets the register edge to User by id.
func (ru *ResearchUpdate) SetRegisterID(id int) *ResearchUpdate {
	ru.mutation.SetRegisterID(id)
	return ru
}

// SetNillableRegisterID sets the register edge to User by id if the given value is not nil.
func (ru *ResearchUpdate) SetNillableRegisterID(id *int) *ResearchUpdate {
	if id != nil {
		ru = ru.SetRegisterID(*id)
	}
	return ru
}

// SetRegister sets the register edge to User.
func (ru *ResearchUpdate) SetRegister(u *User) *ResearchUpdate {
	return ru.SetRegisterID(u.ID)
}

// SetMyDocID sets the myDoc edge to Author by id.
func (ru *ResearchUpdate) SetMyDocID(id int) *ResearchUpdate {
	ru.mutation.SetMyDocID(id)
	return ru
}

// SetNillableMyDocID sets the myDoc edge to Author by id if the given value is not nil.
func (ru *ResearchUpdate) SetNillableMyDocID(id *int) *ResearchUpdate {
	if id != nil {
		ru = ru.SetMyDocID(*id)
	}
	return ru
}

// SetMyDoc sets the myDoc edge to Author.
func (ru *ResearchUpdate) SetMyDoc(a *Author) *ResearchUpdate {
	return ru.SetMyDocID(a.ID)
}

// SetDocTypeID sets the docType edge to Researchtype by id.
func (ru *ResearchUpdate) SetDocTypeID(id int) *ResearchUpdate {
	ru.mutation.SetDocTypeID(id)
	return ru
}

// SetNillableDocTypeID sets the docType edge to Researchtype by id if the given value is not nil.
func (ru *ResearchUpdate) SetNillableDocTypeID(id *int) *ResearchUpdate {
	if id != nil {
		ru = ru.SetDocTypeID(*id)
	}
	return ru
}

// SetDocType sets the docType edge to Researchtype.
func (ru *ResearchUpdate) SetDocType(r *Researchtype) *ResearchUpdate {
	return ru.SetDocTypeID(r.ID)
}

// Mutation returns the ResearchMutation object of the builder.
func (ru *ResearchUpdate) Mutation() *ResearchMutation {
	return ru.mutation
}

// ClearRegister clears the register edge to User.
func (ru *ResearchUpdate) ClearRegister() *ResearchUpdate {
	ru.mutation.ClearRegister()
	return ru
}

// ClearMyDoc clears the myDoc edge to Author.
func (ru *ResearchUpdate) ClearMyDoc() *ResearchUpdate {
	ru.mutation.ClearMyDoc()
	return ru
}

// ClearDocType clears the docType edge to Researchtype.
func (ru *ResearchUpdate) ClearDocType() *ResearchUpdate {
	ru.mutation.ClearDocType()
	return ru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *ResearchUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ru.mutation.DOCNAME(); ok {
		if err := research.DOCNAMEValidator(v); err != nil {
			return 0, &ValidationError{Name: "DOC_NAME", err: fmt.Errorf("ent: validator failed for field \"DOC_NAME\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResearchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResearchUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResearchUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResearchUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ResearchUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   research.Table,
			Columns: research.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: research.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.DOCNAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: research.FieldDOCNAME,
		})
	}
	if value, ok := ru.mutation.DATE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: research.FieldDATE,
		})
	}
	if ru.mutation.RegisterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.RegisterTable,
			Columns: []string{research.RegisterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RegisterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.RegisterTable,
			Columns: []string{research.RegisterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.MyDocCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.MyDocTable,
			Columns: []string{research.MyDocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: author.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.MyDocIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.MyDocTable,
			Columns: []string{research.MyDocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: author.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.DocTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.DocTypeTable,
			Columns: []string{research.DocTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: researchtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.DocTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.DocTypeTable,
			Columns: []string{research.DocTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: researchtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{research.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ResearchUpdateOne is the builder for updating a single Research entity.
type ResearchUpdateOne struct {
	config
	hooks    []Hook
	mutation *ResearchMutation
}

// SetDOCNAME sets the DOC_NAME field.
func (ruo *ResearchUpdateOne) SetDOCNAME(s string) *ResearchUpdateOne {
	ruo.mutation.SetDOCNAME(s)
	return ruo
}

// SetDATE sets the DATE field.
func (ruo *ResearchUpdateOne) SetDATE(t time.Time) *ResearchUpdateOne {
	ruo.mutation.SetDATE(t)
	return ruo
}

// SetNillableDATE sets the DATE field if the given value is not nil.
func (ruo *ResearchUpdateOne) SetNillableDATE(t *time.Time) *ResearchUpdateOne {
	if t != nil {
		ruo.SetDATE(*t)
	}
	return ruo
}

// SetRegisterID sets the register edge to User by id.
func (ruo *ResearchUpdateOne) SetRegisterID(id int) *ResearchUpdateOne {
	ruo.mutation.SetRegisterID(id)
	return ruo
}

// SetNillableRegisterID sets the register edge to User by id if the given value is not nil.
func (ruo *ResearchUpdateOne) SetNillableRegisterID(id *int) *ResearchUpdateOne {
	if id != nil {
		ruo = ruo.SetRegisterID(*id)
	}
	return ruo
}

// SetRegister sets the register edge to User.
func (ruo *ResearchUpdateOne) SetRegister(u *User) *ResearchUpdateOne {
	return ruo.SetRegisterID(u.ID)
}

// SetMyDocID sets the myDoc edge to Author by id.
func (ruo *ResearchUpdateOne) SetMyDocID(id int) *ResearchUpdateOne {
	ruo.mutation.SetMyDocID(id)
	return ruo
}

// SetNillableMyDocID sets the myDoc edge to Author by id if the given value is not nil.
func (ruo *ResearchUpdateOne) SetNillableMyDocID(id *int) *ResearchUpdateOne {
	if id != nil {
		ruo = ruo.SetMyDocID(*id)
	}
	return ruo
}

// SetMyDoc sets the myDoc edge to Author.
func (ruo *ResearchUpdateOne) SetMyDoc(a *Author) *ResearchUpdateOne {
	return ruo.SetMyDocID(a.ID)
}

// SetDocTypeID sets the docType edge to Researchtype by id.
func (ruo *ResearchUpdateOne) SetDocTypeID(id int) *ResearchUpdateOne {
	ruo.mutation.SetDocTypeID(id)
	return ruo
}

// SetNillableDocTypeID sets the docType edge to Researchtype by id if the given value is not nil.
func (ruo *ResearchUpdateOne) SetNillableDocTypeID(id *int) *ResearchUpdateOne {
	if id != nil {
		ruo = ruo.SetDocTypeID(*id)
	}
	return ruo
}

// SetDocType sets the docType edge to Researchtype.
func (ruo *ResearchUpdateOne) SetDocType(r *Researchtype) *ResearchUpdateOne {
	return ruo.SetDocTypeID(r.ID)
}

// Mutation returns the ResearchMutation object of the builder.
func (ruo *ResearchUpdateOne) Mutation() *ResearchMutation {
	return ruo.mutation
}

// ClearRegister clears the register edge to User.
func (ruo *ResearchUpdateOne) ClearRegister() *ResearchUpdateOne {
	ruo.mutation.ClearRegister()
	return ruo
}

// ClearMyDoc clears the myDoc edge to Author.
func (ruo *ResearchUpdateOne) ClearMyDoc() *ResearchUpdateOne {
	ruo.mutation.ClearMyDoc()
	return ruo
}

// ClearDocType clears the docType edge to Researchtype.
func (ruo *ResearchUpdateOne) ClearDocType() *ResearchUpdateOne {
	ruo.mutation.ClearDocType()
	return ruo
}

// Save executes the query and returns the updated entity.
func (ruo *ResearchUpdateOne) Save(ctx context.Context) (*Research, error) {
	if v, ok := ruo.mutation.DOCNAME(); ok {
		if err := research.DOCNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "DOC_NAME", err: fmt.Errorf("ent: validator failed for field \"DOC_NAME\": %w", err)}
		}
	}

	var (
		err  error
		node *Research
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResearchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResearchUpdateOne) SaveX(ctx context.Context) *Research {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *ResearchUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResearchUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ResearchUpdateOne) sqlSave(ctx context.Context) (r *Research, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   research.Table,
			Columns: research.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: research.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Research.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.DOCNAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: research.FieldDOCNAME,
		})
	}
	if value, ok := ruo.mutation.DATE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: research.FieldDATE,
		})
	}
	if ruo.mutation.RegisterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.RegisterTable,
			Columns: []string{research.RegisterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RegisterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.RegisterTable,
			Columns: []string{research.RegisterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.MyDocCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.MyDocTable,
			Columns: []string{research.MyDocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: author.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.MyDocIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.MyDocTable,
			Columns: []string{research.MyDocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: author.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.DocTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.DocTypeTable,
			Columns: []string{research.DocTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: researchtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.DocTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.DocTypeTable,
			Columns: []string{research.DocTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: researchtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Research{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{research.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}