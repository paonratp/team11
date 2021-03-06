// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/booking"
	"github.com/team11/app/ent/cliententity"
	"github.com/team11/app/ent/predicate"
	"github.com/team11/app/ent/servicepoint"
	"github.com/team11/app/ent/user"
)

// BookingUpdate is the builder for updating Booking entities.
type BookingUpdate struct {
	config
	hooks      []Hook
	mutation   *BookingMutation
	predicates []predicate.Booking
}

// Where adds a new predicate for the builder.
func (bu *BookingUpdate) Where(ps ...predicate.Booking) *BookingUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetBOOKINGDATE sets the BOOKING_DATE field.
func (bu *BookingUpdate) SetBOOKINGDATE(t time.Time) *BookingUpdate {
	bu.mutation.SetBOOKINGDATE(t)
	return bu
}

// SetNillableBOOKINGDATE sets the BOOKING_DATE field if the given value is not nil.
func (bu *BookingUpdate) SetNillableBOOKINGDATE(t *time.Time) *BookingUpdate {
	if t != nil {
		bu.SetBOOKINGDATE(*t)
	}
	return bu
}

// SetTIMELEFT sets the TIME_LEFT field.
func (bu *BookingUpdate) SetTIMELEFT(t time.Time) *BookingUpdate {
	bu.mutation.SetTIMELEFT(t)
	return bu
}

// SetUSERNUMBER sets the USER_NUMBER field.
func (bu *BookingUpdate) SetUSERNUMBER(i int) *BookingUpdate {
	bu.mutation.ResetUSERNUMBER()
	bu.mutation.SetUSERNUMBER(i)
	return bu
}

// AddUSERNUMBER adds i to USER_NUMBER.
func (bu *BookingUpdate) AddUSERNUMBER(i int) *BookingUpdate {
	bu.mutation.AddUSERNUMBER(i)
	return bu
}

// SetBORROWITEM sets the BORROW_ITEM field.
func (bu *BookingUpdate) SetBORROWITEM(i int) *BookingUpdate {
	bu.mutation.ResetBORROWITEM()
	bu.mutation.SetBORROWITEM(i)
	return bu
}

// AddBORROWITEM adds i to BORROW_ITEM.
func (bu *BookingUpdate) AddBORROWITEM(i int) *BookingUpdate {
	bu.mutation.AddBORROWITEM(i)
	return bu
}

// SetPHONENUMBER sets the PHONE_NUMBER field.
func (bu *BookingUpdate) SetPHONENUMBER(s string) *BookingUpdate {
	bu.mutation.SetPHONENUMBER(s)
	return bu
}

// SetUsedbyID sets the usedby edge to User by id.
func (bu *BookingUpdate) SetUsedbyID(id int) *BookingUpdate {
	bu.mutation.SetUsedbyID(id)
	return bu
}

// SetNillableUsedbyID sets the usedby edge to User by id if the given value is not nil.
func (bu *BookingUpdate) SetNillableUsedbyID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetUsedbyID(*id)
	}
	return bu
}

// SetUsedby sets the usedby edge to User.
func (bu *BookingUpdate) SetUsedby(u *User) *BookingUpdate {
	return bu.SetUsedbyID(u.ID)
}

// SetGetserviceID sets the getservice edge to ServicePoint by id.
func (bu *BookingUpdate) SetGetserviceID(id int) *BookingUpdate {
	bu.mutation.SetGetserviceID(id)
	return bu
}

// SetNillableGetserviceID sets the getservice edge to ServicePoint by id if the given value is not nil.
func (bu *BookingUpdate) SetNillableGetserviceID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetGetserviceID(*id)
	}
	return bu
}

// SetGetservice sets the getservice edge to ServicePoint.
func (bu *BookingUpdate) SetGetservice(s *ServicePoint) *BookingUpdate {
	return bu.SetGetserviceID(s.ID)
}

// SetUsingID sets the using edge to ClientEntity by id.
func (bu *BookingUpdate) SetUsingID(id int) *BookingUpdate {
	bu.mutation.SetUsingID(id)
	return bu
}

// SetNillableUsingID sets the using edge to ClientEntity by id if the given value is not nil.
func (bu *BookingUpdate) SetNillableUsingID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetUsingID(*id)
	}
	return bu
}

// SetUsing sets the using edge to ClientEntity.
func (bu *BookingUpdate) SetUsing(c *ClientEntity) *BookingUpdate {
	return bu.SetUsingID(c.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bu *BookingUpdate) Mutation() *BookingMutation {
	return bu.mutation
}

// ClearUsedby clears the usedby edge to User.
func (bu *BookingUpdate) ClearUsedby() *BookingUpdate {
	bu.mutation.ClearUsedby()
	return bu
}

// ClearGetservice clears the getservice edge to ServicePoint.
func (bu *BookingUpdate) ClearGetservice() *BookingUpdate {
	bu.mutation.ClearGetservice()
	return bu
}

// ClearUsing clears the using edge to ClientEntity.
func (bu *BookingUpdate) ClearUsing() *BookingUpdate {
	bu.mutation.ClearUsing()
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BookingUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := bu.mutation.USERNUMBER(); ok {
		if err := booking.USERNUMBERValidator(v); err != nil {
			return 0, &ValidationError{Name: "USER_NUMBER", err: fmt.Errorf("ent: validator failed for field \"USER_NUMBER\": %w", err)}
		}
	}
	if v, ok := bu.mutation.BORROWITEM(); ok {
		if err := booking.BORROWITEMValidator(v); err != nil {
			return 0, &ValidationError{Name: "BORROW_ITEM", err: fmt.Errorf("ent: validator failed for field \"BORROW_ITEM\": %w", err)}
		}
	}
	if v, ok := bu.mutation.PHONENUMBER(); ok {
		if err := booking.PHONENUMBERValidator(v); err != nil {
			return 0, &ValidationError{Name: "PHONE_NUMBER", err: fmt.Errorf("ent: validator failed for field \"PHONE_NUMBER\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BookingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   booking.Table,
			Columns: booking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: booking.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.BOOKINGDATE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldBOOKINGDATE,
		})
	}
	if value, ok := bu.mutation.TIMELEFT(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldTIMELEFT,
		})
	}
	if value, ok := bu.mutation.USERNUMBER(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: booking.FieldUSERNUMBER,
		})
	}
	if value, ok := bu.mutation.AddedUSERNUMBER(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: booking.FieldUSERNUMBER,
		})
	}
	if value, ok := bu.mutation.BORROWITEM(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: booking.FieldBORROWITEM,
		})
	}
	if value, ok := bu.mutation.AddedBORROWITEM(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: booking.FieldBORROWITEM,
		})
	}
	if value, ok := bu.mutation.PHONENUMBER(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: booking.FieldPHONENUMBER,
		})
	}
	if bu.mutation.UsedbyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsedbyTable,
			Columns: []string{booking.UsedbyColumn},
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
	if nodes := bu.mutation.UsedbyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsedbyTable,
			Columns: []string{booking.UsedbyColumn},
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
	if bu.mutation.GetserviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.GetserviceTable,
			Columns: []string{booking.GetserviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicepoint.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.GetserviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.GetserviceTable,
			Columns: []string{booking.GetserviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicepoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.UsingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsingTable,
			Columns: []string{booking.UsingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cliententity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.UsingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsingTable,
			Columns: []string{booking.UsingColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BookingUpdateOne is the builder for updating a single Booking entity.
type BookingUpdateOne struct {
	config
	hooks    []Hook
	mutation *BookingMutation
}

// SetBOOKINGDATE sets the BOOKING_DATE field.
func (buo *BookingUpdateOne) SetBOOKINGDATE(t time.Time) *BookingUpdateOne {
	buo.mutation.SetBOOKINGDATE(t)
	return buo
}

// SetNillableBOOKINGDATE sets the BOOKING_DATE field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableBOOKINGDATE(t *time.Time) *BookingUpdateOne {
	if t != nil {
		buo.SetBOOKINGDATE(*t)
	}
	return buo
}

// SetTIMELEFT sets the TIME_LEFT field.
func (buo *BookingUpdateOne) SetTIMELEFT(t time.Time) *BookingUpdateOne {
	buo.mutation.SetTIMELEFT(t)
	return buo
}

// SetUSERNUMBER sets the USER_NUMBER field.
func (buo *BookingUpdateOne) SetUSERNUMBER(i int) *BookingUpdateOne {
	buo.mutation.ResetUSERNUMBER()
	buo.mutation.SetUSERNUMBER(i)
	return buo
}

// AddUSERNUMBER adds i to USER_NUMBER.
func (buo *BookingUpdateOne) AddUSERNUMBER(i int) *BookingUpdateOne {
	buo.mutation.AddUSERNUMBER(i)
	return buo
}

// SetBORROWITEM sets the BORROW_ITEM field.
func (buo *BookingUpdateOne) SetBORROWITEM(i int) *BookingUpdateOne {
	buo.mutation.ResetBORROWITEM()
	buo.mutation.SetBORROWITEM(i)
	return buo
}

// AddBORROWITEM adds i to BORROW_ITEM.
func (buo *BookingUpdateOne) AddBORROWITEM(i int) *BookingUpdateOne {
	buo.mutation.AddBORROWITEM(i)
	return buo
}

// SetPHONENUMBER sets the PHONE_NUMBER field.
func (buo *BookingUpdateOne) SetPHONENUMBER(s string) *BookingUpdateOne {
	buo.mutation.SetPHONENUMBER(s)
	return buo
}

// SetUsedbyID sets the usedby edge to User by id.
func (buo *BookingUpdateOne) SetUsedbyID(id int) *BookingUpdateOne {
	buo.mutation.SetUsedbyID(id)
	return buo
}

// SetNillableUsedbyID sets the usedby edge to User by id if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableUsedbyID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetUsedbyID(*id)
	}
	return buo
}

// SetUsedby sets the usedby edge to User.
func (buo *BookingUpdateOne) SetUsedby(u *User) *BookingUpdateOne {
	return buo.SetUsedbyID(u.ID)
}

// SetGetserviceID sets the getservice edge to ServicePoint by id.
func (buo *BookingUpdateOne) SetGetserviceID(id int) *BookingUpdateOne {
	buo.mutation.SetGetserviceID(id)
	return buo
}

// SetNillableGetserviceID sets the getservice edge to ServicePoint by id if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableGetserviceID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetGetserviceID(*id)
	}
	return buo
}

// SetGetservice sets the getservice edge to ServicePoint.
func (buo *BookingUpdateOne) SetGetservice(s *ServicePoint) *BookingUpdateOne {
	return buo.SetGetserviceID(s.ID)
}

// SetUsingID sets the using edge to ClientEntity by id.
func (buo *BookingUpdateOne) SetUsingID(id int) *BookingUpdateOne {
	buo.mutation.SetUsingID(id)
	return buo
}

// SetNillableUsingID sets the using edge to ClientEntity by id if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableUsingID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetUsingID(*id)
	}
	return buo
}

// SetUsing sets the using edge to ClientEntity.
func (buo *BookingUpdateOne) SetUsing(c *ClientEntity) *BookingUpdateOne {
	return buo.SetUsingID(c.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (buo *BookingUpdateOne) Mutation() *BookingMutation {
	return buo.mutation
}

// ClearUsedby clears the usedby edge to User.
func (buo *BookingUpdateOne) ClearUsedby() *BookingUpdateOne {
	buo.mutation.ClearUsedby()
	return buo
}

// ClearGetservice clears the getservice edge to ServicePoint.
func (buo *BookingUpdateOne) ClearGetservice() *BookingUpdateOne {
	buo.mutation.ClearGetservice()
	return buo
}

// ClearUsing clears the using edge to ClientEntity.
func (buo *BookingUpdateOne) ClearUsing() *BookingUpdateOne {
	buo.mutation.ClearUsing()
	return buo
}

// Save executes the query and returns the updated entity.
func (buo *BookingUpdateOne) Save(ctx context.Context) (*Booking, error) {
	if v, ok := buo.mutation.USERNUMBER(); ok {
		if err := booking.USERNUMBERValidator(v); err != nil {
			return nil, &ValidationError{Name: "USER_NUMBER", err: fmt.Errorf("ent: validator failed for field \"USER_NUMBER\": %w", err)}
		}
	}
	if v, ok := buo.mutation.BORROWITEM(); ok {
		if err := booking.BORROWITEMValidator(v); err != nil {
			return nil, &ValidationError{Name: "BORROW_ITEM", err: fmt.Errorf("ent: validator failed for field \"BORROW_ITEM\": %w", err)}
		}
	}
	if v, ok := buo.mutation.PHONENUMBER(); ok {
		if err := booking.PHONENUMBERValidator(v); err != nil {
			return nil, &ValidationError{Name: "PHONE_NUMBER", err: fmt.Errorf("ent: validator failed for field \"PHONE_NUMBER\": %w", err)}
		}
	}

	var (
		err  error
		node *Booking
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookingUpdateOne) SaveX(ctx context.Context) *Booking {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BookingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BookingUpdateOne) sqlSave(ctx context.Context) (b *Booking, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   booking.Table,
			Columns: booking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: booking.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Booking.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.BOOKINGDATE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldBOOKINGDATE,
		})
	}
	if value, ok := buo.mutation.TIMELEFT(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldTIMELEFT,
		})
	}
	if value, ok := buo.mutation.USERNUMBER(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: booking.FieldUSERNUMBER,
		})
	}
	if value, ok := buo.mutation.AddedUSERNUMBER(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: booking.FieldUSERNUMBER,
		})
	}
	if value, ok := buo.mutation.BORROWITEM(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: booking.FieldBORROWITEM,
		})
	}
	if value, ok := buo.mutation.AddedBORROWITEM(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: booking.FieldBORROWITEM,
		})
	}
	if value, ok := buo.mutation.PHONENUMBER(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: booking.FieldPHONENUMBER,
		})
	}
	if buo.mutation.UsedbyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsedbyTable,
			Columns: []string{booking.UsedbyColumn},
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
	if nodes := buo.mutation.UsedbyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsedbyTable,
			Columns: []string{booking.UsedbyColumn},
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
	if buo.mutation.GetserviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.GetserviceTable,
			Columns: []string{booking.GetserviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicepoint.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.GetserviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.GetserviceTable,
			Columns: []string{booking.GetserviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicepoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.UsingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsingTable,
			Columns: []string{booking.UsingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cliententity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.UsingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsingTable,
			Columns: []string{booking.UsingColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	b = &Booking{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}
