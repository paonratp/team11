// Code generated by entc, DO NOT EDIT.

package booking

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team11/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BOOKINGDATE applies equality check predicate on the "BOOKING_DATE" field. It's identical to BOOKINGDATEEQ.
func BOOKINGDATE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBOOKINGDATE), v))
	})
}

// TIMELEFT applies equality check predicate on the "TIME_LEFT" field. It's identical to TIMELEFTEQ.
func TIMELEFT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTIMELEFT), v))
	})
}

// USERNUMBER applies equality check predicate on the "USER_NUMBER" field. It's identical to USERNUMBEREQ.
func USERNUMBER(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUSERNUMBER), v))
	})
}

// BORROWITEM applies equality check predicate on the "BORROW_ITEM" field. It's identical to BORROWITEMEQ.
func BORROWITEM(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBORROWITEM), v))
	})
}

// PHONENUMBER applies equality check predicate on the "PHONE_NUMBER" field. It's identical to PHONENUMBEREQ.
func PHONENUMBER(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHONENUMBER), v))
	})
}

// BOOKINGDATEEQ applies the EQ predicate on the "BOOKING_DATE" field.
func BOOKINGDATEEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBOOKINGDATE), v))
	})
}

// BOOKINGDATENEQ applies the NEQ predicate on the "BOOKING_DATE" field.
func BOOKINGDATENEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBOOKINGDATE), v))
	})
}

// BOOKINGDATEIn applies the In predicate on the "BOOKING_DATE" field.
func BOOKINGDATEIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBOOKINGDATE), v...))
	})
}

// BOOKINGDATENotIn applies the NotIn predicate on the "BOOKING_DATE" field.
func BOOKINGDATENotIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBOOKINGDATE), v...))
	})
}

// BOOKINGDATEGT applies the GT predicate on the "BOOKING_DATE" field.
func BOOKINGDATEGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBOOKINGDATE), v))
	})
}

// BOOKINGDATEGTE applies the GTE predicate on the "BOOKING_DATE" field.
func BOOKINGDATEGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBOOKINGDATE), v))
	})
}

// BOOKINGDATELT applies the LT predicate on the "BOOKING_DATE" field.
func BOOKINGDATELT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBOOKINGDATE), v))
	})
}

// BOOKINGDATELTE applies the LTE predicate on the "BOOKING_DATE" field.
func BOOKINGDATELTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBOOKINGDATE), v))
	})
}

// TIMELEFTEQ applies the EQ predicate on the "TIME_LEFT" field.
func TIMELEFTEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTIMELEFT), v))
	})
}

// TIMELEFTNEQ applies the NEQ predicate on the "TIME_LEFT" field.
func TIMELEFTNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTIMELEFT), v))
	})
}

// TIMELEFTIn applies the In predicate on the "TIME_LEFT" field.
func TIMELEFTIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTIMELEFT), v...))
	})
}

// TIMELEFTNotIn applies the NotIn predicate on the "TIME_LEFT" field.
func TIMELEFTNotIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTIMELEFT), v...))
	})
}

// TIMELEFTGT applies the GT predicate on the "TIME_LEFT" field.
func TIMELEFTGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTIMELEFT), v))
	})
}

// TIMELEFTGTE applies the GTE predicate on the "TIME_LEFT" field.
func TIMELEFTGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTIMELEFT), v))
	})
}

// TIMELEFTLT applies the LT predicate on the "TIME_LEFT" field.
func TIMELEFTLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTIMELEFT), v))
	})
}

// TIMELEFTLTE applies the LTE predicate on the "TIME_LEFT" field.
func TIMELEFTLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTIMELEFT), v))
	})
}

// USERNUMBEREQ applies the EQ predicate on the "USER_NUMBER" field.
func USERNUMBEREQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUSERNUMBER), v))
	})
}

// USERNUMBERNEQ applies the NEQ predicate on the "USER_NUMBER" field.
func USERNUMBERNEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUSERNUMBER), v))
	})
}

// USERNUMBERIn applies the In predicate on the "USER_NUMBER" field.
func USERNUMBERIn(vs ...int) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUSERNUMBER), v...))
	})
}

// USERNUMBERNotIn applies the NotIn predicate on the "USER_NUMBER" field.
func USERNUMBERNotIn(vs ...int) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUSERNUMBER), v...))
	})
}

// USERNUMBERGT applies the GT predicate on the "USER_NUMBER" field.
func USERNUMBERGT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUSERNUMBER), v))
	})
}

// USERNUMBERGTE applies the GTE predicate on the "USER_NUMBER" field.
func USERNUMBERGTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUSERNUMBER), v))
	})
}

// USERNUMBERLT applies the LT predicate on the "USER_NUMBER" field.
func USERNUMBERLT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUSERNUMBER), v))
	})
}

// USERNUMBERLTE applies the LTE predicate on the "USER_NUMBER" field.
func USERNUMBERLTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUSERNUMBER), v))
	})
}

// BORROWITEMEQ applies the EQ predicate on the "BORROW_ITEM" field.
func BORROWITEMEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBORROWITEM), v))
	})
}

// BORROWITEMNEQ applies the NEQ predicate on the "BORROW_ITEM" field.
func BORROWITEMNEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBORROWITEM), v))
	})
}

// BORROWITEMIn applies the In predicate on the "BORROW_ITEM" field.
func BORROWITEMIn(vs ...int) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBORROWITEM), v...))
	})
}

// BORROWITEMNotIn applies the NotIn predicate on the "BORROW_ITEM" field.
func BORROWITEMNotIn(vs ...int) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBORROWITEM), v...))
	})
}

// BORROWITEMGT applies the GT predicate on the "BORROW_ITEM" field.
func BORROWITEMGT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBORROWITEM), v))
	})
}

// BORROWITEMGTE applies the GTE predicate on the "BORROW_ITEM" field.
func BORROWITEMGTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBORROWITEM), v))
	})
}

// BORROWITEMLT applies the LT predicate on the "BORROW_ITEM" field.
func BORROWITEMLT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBORROWITEM), v))
	})
}

// BORROWITEMLTE applies the LTE predicate on the "BORROW_ITEM" field.
func BORROWITEMLTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBORROWITEM), v))
	})
}

// PHONENUMBEREQ applies the EQ predicate on the "PHONE_NUMBER" field.
func PHONENUMBEREQ(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERNEQ applies the NEQ predicate on the "PHONE_NUMBER" field.
func PHONENUMBERNEQ(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERIn applies the In predicate on the "PHONE_NUMBER" field.
func PHONENUMBERIn(vs ...string) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPHONENUMBER), v...))
	})
}

// PHONENUMBERNotIn applies the NotIn predicate on the "PHONE_NUMBER" field.
func PHONENUMBERNotIn(vs ...string) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPHONENUMBER), v...))
	})
}

// PHONENUMBERGT applies the GT predicate on the "PHONE_NUMBER" field.
func PHONENUMBERGT(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERGTE applies the GTE predicate on the "PHONE_NUMBER" field.
func PHONENUMBERGTE(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERLT applies the LT predicate on the "PHONE_NUMBER" field.
func PHONENUMBERLT(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERLTE applies the LTE predicate on the "PHONE_NUMBER" field.
func PHONENUMBERLTE(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERContains applies the Contains predicate on the "PHONE_NUMBER" field.
func PHONENUMBERContains(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERHasPrefix applies the HasPrefix predicate on the "PHONE_NUMBER" field.
func PHONENUMBERHasPrefix(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERHasSuffix applies the HasSuffix predicate on the "PHONE_NUMBER" field.
func PHONENUMBERHasSuffix(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBEREqualFold applies the EqualFold predicate on the "PHONE_NUMBER" field.
func PHONENUMBEREqualFold(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERContainsFold applies the ContainsFold predicate on the "PHONE_NUMBER" field.
func PHONENUMBERContainsFold(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPHONENUMBER), v))
	})
}

// HasUsedby applies the HasEdge predicate on the "usedby" edge.
func HasUsedby() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsedbyTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsedbyTable, UsedbyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsedbyWith applies the HasEdge predicate on the "usedby" edge with a given conditions (other predicates).
func HasUsedbyWith(preds ...predicate.User) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsedbyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsedbyTable, UsedbyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGetservice applies the HasEdge predicate on the "getservice" edge.
func HasGetservice() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GetserviceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GetserviceTable, GetserviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGetserviceWith applies the HasEdge predicate on the "getservice" edge with a given conditions (other predicates).
func HasGetserviceWith(preds ...predicate.ServicePoint) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GetserviceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GetserviceTable, GetserviceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsing applies the HasEdge predicate on the "using" edge.
func HasUsing() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsingTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsingTable, UsingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsingWith applies the HasEdge predicate on the "using" edge with a given conditions (other predicates).
func HasUsingWith(preds ...predicate.ClientEntity) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsingInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsingTable, UsingColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		p(s.Not())
	})
}
