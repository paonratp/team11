// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/bookborrow"
	"github.com/team11/app/ent/booking"
	"github.com/team11/app/ent/servicepoint"
)

// ServicePointCreate is the builder for creating a ServicePoint entity.
type ServicePointCreate struct {
	config
	mutation *ServicePointMutation
	hooks    []Hook
}

// SetBUILDINGNAME sets the BUILDING_NAME field.
func (spc *ServicePointCreate) SetBUILDINGNAME(s string) *ServicePointCreate {
	spc.mutation.SetBUILDINGNAME(s)
	return spc
}

// SetCOUNTERNUMBER sets the COUNTER_NUMBER field.
func (spc *ServicePointCreate) SetCOUNTERNUMBER(s string) *ServicePointCreate {
	spc.mutation.SetCOUNTERNUMBER(s)
	return spc
}

// AddFromIDs adds the from edge to Bookborrow by ids.
func (spc *ServicePointCreate) AddFromIDs(ids ...int) *ServicePointCreate {
	spc.mutation.AddFromIDs(ids...)
	return spc
}

// AddFrom adds the from edges to Bookborrow.
func (spc *ServicePointCreate) AddFrom(b ...*Bookborrow) *ServicePointCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spc.AddFromIDs(ids...)
}

// AddServicepointIDs adds the servicepoint edge to Booking by ids.
func (spc *ServicePointCreate) AddServicepointIDs(ids ...int) *ServicePointCreate {
	spc.mutation.AddServicepointIDs(ids...)
	return spc
}

// AddServicepoint adds the servicepoint edges to Booking.
func (spc *ServicePointCreate) AddServicepoint(b ...*Booking) *ServicePointCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spc.AddServicepointIDs(ids...)
}

// Mutation returns the ServicePointMutation object of the builder.
func (spc *ServicePointCreate) Mutation() *ServicePointMutation {
	return spc.mutation
}

// Save creates the ServicePoint in the database.
func (spc *ServicePointCreate) Save(ctx context.Context) (*ServicePoint, error) {
	if _, ok := spc.mutation.BUILDINGNAME(); !ok {
		return nil, &ValidationError{Name: "BUILDING_NAME", err: errors.New("ent: missing required field \"BUILDING_NAME\"")}
	}
	if v, ok := spc.mutation.BUILDINGNAME(); ok {
		if err := servicepoint.BUILDINGNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "BUILDING_NAME", err: fmt.Errorf("ent: validator failed for field \"BUILDING_NAME\": %w", err)}
		}
	}
	if _, ok := spc.mutation.COUNTERNUMBER(); !ok {
		return nil, &ValidationError{Name: "COUNTER_NUMBER", err: errors.New("ent: missing required field \"COUNTER_NUMBER\"")}
	}
	if v, ok := spc.mutation.COUNTERNUMBER(); ok {
		if err := servicepoint.COUNTERNUMBERValidator(v); err != nil {
			return nil, &ValidationError{Name: "COUNTER_NUMBER", err: fmt.Errorf("ent: validator failed for field \"COUNTER_NUMBER\": %w", err)}
		}
	}
	var (
		err  error
		node *ServicePoint
	)
	if len(spc.hooks) == 0 {
		node, err = spc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServicePointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spc.mutation = mutation
			node, err = spc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(spc.hooks) - 1; i >= 0; i-- {
			mut = spc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (spc *ServicePointCreate) SaveX(ctx context.Context) *ServicePoint {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (spc *ServicePointCreate) sqlSave(ctx context.Context) (*ServicePoint, error) {
	sp, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	sp.ID = int(id)
	return sp, nil
}

func (spc *ServicePointCreate) createSpec() (*ServicePoint, *sqlgraph.CreateSpec) {
	var (
		sp    = &ServicePoint{config: spc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: servicepoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servicepoint.FieldID,
			},
		}
	)
	if value, ok := spc.mutation.BUILDINGNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicepoint.FieldBUILDINGNAME,
		})
		sp.BUILDINGNAME = value
	}
	if value, ok := spc.mutation.COUNTERNUMBER(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicepoint.FieldCOUNTERNUMBER,
		})
		sp.COUNTERNUMBER = value
	}
	if nodes := spc.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.FromTable,
			Columns: []string{servicepoint.FromColumn},
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
	if nodes := spc.mutation.ServicepointIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.ServicepointTable,
			Columns: []string{servicepoint.ServicepointColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return sp, _spec
}