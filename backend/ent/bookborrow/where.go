// Code generated by entc, DO NOT EDIT.

package bookborrow

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team11/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BORROWDATE applies equality check predicate on the "BORROW_DATE" field. It's identical to BORROWDATEEQ.
func BORROWDATE(v time.Time) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBORROWDATE), v))
	})
}

// BORROWDATEEQ applies the EQ predicate on the "BORROW_DATE" field.
func BORROWDATEEQ(v time.Time) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBORROWDATE), v))
	})
}

// BORROWDATENEQ applies the NEQ predicate on the "BORROW_DATE" field.
func BORROWDATENEQ(v time.Time) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBORROWDATE), v))
	})
}

// BORROWDATEIn applies the In predicate on the "BORROW_DATE" field.
func BORROWDATEIn(vs ...time.Time) predicate.Bookborrow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookborrow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBORROWDATE), v...))
	})
}

// BORROWDATENotIn applies the NotIn predicate on the "BORROW_DATE" field.
func BORROWDATENotIn(vs ...time.Time) predicate.Bookborrow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookborrow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBORROWDATE), v...))
	})
}

// BORROWDATEGT applies the GT predicate on the "BORROW_DATE" field.
func BORROWDATEGT(v time.Time) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBORROWDATE), v))
	})
}

// BORROWDATEGTE applies the GTE predicate on the "BORROW_DATE" field.
func BORROWDATEGTE(v time.Time) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBORROWDATE), v))
	})
}

// BORROWDATELT applies the LT predicate on the "BORROW_DATE" field.
func BORROWDATELT(v time.Time) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBORROWDATE), v))
	})
}

// BORROWDATELTE applies the LTE predicate on the "BORROW_DATE" field.
func BORROWDATELTE(v time.Time) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBORROWDATE), v))
	})
}

// HasUSER applies the HasEdge predicate on the "USER" edge.
func HasUSER() predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(USERTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, USERTable, USERColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUSERWith applies the HasEdge predicate on the "USER" edge with a given conditions (other predicates).
func HasUSERWith(preds ...predicate.User) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(USERInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, USERTable, USERColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBOOK applies the HasEdge predicate on the "BOOK" edge.
func HasBOOK() predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BOOKTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BOOKTable, BOOKColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBOOKWith applies the HasEdge predicate on the "BOOK" edge with a given conditions (other predicates).
func HasBOOKWith(preds ...predicate.Book) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BOOKInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BOOKTable, BOOKColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSERVICEPOINT applies the HasEdge predicate on the "SERVICEPOINT" edge.
func HasSERVICEPOINT() predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SERVICEPOINTTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SERVICEPOINTTable, SERVICEPOINTColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSERVICEPOINTWith applies the HasEdge predicate on the "SERVICEPOINT" edge with a given conditions (other predicates).
func HasSERVICEPOINTWith(preds ...predicate.ServicePoint) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SERVICEPOINTInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SERVICEPOINTTable, SERVICEPOINTColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBorrowed applies the HasEdge predicate on the "borrowed" edge.
func HasBorrowed() predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BorrowedTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BorrowedTable, BorrowedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBorrowedWith applies the HasEdge predicate on the "borrowed" edge with a given conditions (other predicates).
func HasBorrowedWith(preds ...predicate.Bookreturn) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BorrowedInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BorrowedTable, BorrowedColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Bookborrow) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Bookborrow) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
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
func Not(p predicate.Bookborrow) predicate.Bookborrow {
	return predicate.Bookborrow(func(s *sql.Selector) {
		p(s.Not())
	})
}
