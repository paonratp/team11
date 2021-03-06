// Code generated by entc, DO NOT EDIT.

package preemption

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team11/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PreemptTime applies equality check predicate on the "PreemptTime" field. It's identical to PreemptTimeEQ.
func PreemptTime(v time.Time) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPreemptTime), v))
	})
}

// Phonenumber applies equality check predicate on the "Phonenumber" field. It's identical to PhonenumberEQ.
func Phonenumber(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhonenumber), v))
	})
}

// Surrogateid applies equality check predicate on the "Surrogateid" field. It's identical to SurrogateidEQ.
func Surrogateid(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSurrogateid), v))
	})
}

// Surrogatephone applies equality check predicate on the "Surrogatephone" field. It's identical to SurrogatephoneEQ.
func Surrogatephone(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSurrogatephone), v))
	})
}

// PreemptTimeEQ applies the EQ predicate on the "PreemptTime" field.
func PreemptTimeEQ(v time.Time) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPreemptTime), v))
	})
}

// PreemptTimeNEQ applies the NEQ predicate on the "PreemptTime" field.
func PreemptTimeNEQ(v time.Time) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPreemptTime), v))
	})
}

// PreemptTimeIn applies the In predicate on the "PreemptTime" field.
func PreemptTimeIn(vs ...time.Time) predicate.Preemption {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Preemption(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPreemptTime), v...))
	})
}

// PreemptTimeNotIn applies the NotIn predicate on the "PreemptTime" field.
func PreemptTimeNotIn(vs ...time.Time) predicate.Preemption {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Preemption(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPreemptTime), v...))
	})
}

// PreemptTimeGT applies the GT predicate on the "PreemptTime" field.
func PreemptTimeGT(v time.Time) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPreemptTime), v))
	})
}

// PreemptTimeGTE applies the GTE predicate on the "PreemptTime" field.
func PreemptTimeGTE(v time.Time) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPreemptTime), v))
	})
}

// PreemptTimeLT applies the LT predicate on the "PreemptTime" field.
func PreemptTimeLT(v time.Time) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPreemptTime), v))
	})
}

// PreemptTimeLTE applies the LTE predicate on the "PreemptTime" field.
func PreemptTimeLTE(v time.Time) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPreemptTime), v))
	})
}

// PhonenumberEQ applies the EQ predicate on the "Phonenumber" field.
func PhonenumberEQ(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberNEQ applies the NEQ predicate on the "Phonenumber" field.
func PhonenumberNEQ(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberIn applies the In predicate on the "Phonenumber" field.
func PhonenumberIn(vs ...string) predicate.Preemption {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Preemption(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhonenumber), v...))
	})
}

// PhonenumberNotIn applies the NotIn predicate on the "Phonenumber" field.
func PhonenumberNotIn(vs ...string) predicate.Preemption {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Preemption(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhonenumber), v...))
	})
}

// PhonenumberGT applies the GT predicate on the "Phonenumber" field.
func PhonenumberGT(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberGTE applies the GTE predicate on the "Phonenumber" field.
func PhonenumberGTE(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberLT applies the LT predicate on the "Phonenumber" field.
func PhonenumberLT(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberLTE applies the LTE predicate on the "Phonenumber" field.
func PhonenumberLTE(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberContains applies the Contains predicate on the "Phonenumber" field.
func PhonenumberContains(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberHasPrefix applies the HasPrefix predicate on the "Phonenumber" field.
func PhonenumberHasPrefix(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberHasSuffix applies the HasSuffix predicate on the "Phonenumber" field.
func PhonenumberHasSuffix(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberEqualFold applies the EqualFold predicate on the "Phonenumber" field.
func PhonenumberEqualFold(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberContainsFold applies the ContainsFold predicate on the "Phonenumber" field.
func PhonenumberContainsFold(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhonenumber), v))
	})
}

// SurrogateidEQ applies the EQ predicate on the "Surrogateid" field.
func SurrogateidEQ(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidNEQ applies the NEQ predicate on the "Surrogateid" field.
func SurrogateidNEQ(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidIn applies the In predicate on the "Surrogateid" field.
func SurrogateidIn(vs ...string) predicate.Preemption {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Preemption(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSurrogateid), v...))
	})
}

// SurrogateidNotIn applies the NotIn predicate on the "Surrogateid" field.
func SurrogateidNotIn(vs ...string) predicate.Preemption {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Preemption(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSurrogateid), v...))
	})
}

// SurrogateidGT applies the GT predicate on the "Surrogateid" field.
func SurrogateidGT(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidGTE applies the GTE predicate on the "Surrogateid" field.
func SurrogateidGTE(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidLT applies the LT predicate on the "Surrogateid" field.
func SurrogateidLT(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidLTE applies the LTE predicate on the "Surrogateid" field.
func SurrogateidLTE(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidContains applies the Contains predicate on the "Surrogateid" field.
func SurrogateidContains(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidHasPrefix applies the HasPrefix predicate on the "Surrogateid" field.
func SurrogateidHasPrefix(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidHasSuffix applies the HasSuffix predicate on the "Surrogateid" field.
func SurrogateidHasSuffix(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidEqualFold applies the EqualFold predicate on the "Surrogateid" field.
func SurrogateidEqualFold(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSurrogateid), v))
	})
}

// SurrogateidContainsFold applies the ContainsFold predicate on the "Surrogateid" field.
func SurrogateidContainsFold(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSurrogateid), v))
	})
}

// SurrogatephoneEQ applies the EQ predicate on the "Surrogatephone" field.
func SurrogatephoneEQ(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneNEQ applies the NEQ predicate on the "Surrogatephone" field.
func SurrogatephoneNEQ(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneIn applies the In predicate on the "Surrogatephone" field.
func SurrogatephoneIn(vs ...string) predicate.Preemption {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Preemption(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSurrogatephone), v...))
	})
}

// SurrogatephoneNotIn applies the NotIn predicate on the "Surrogatephone" field.
func SurrogatephoneNotIn(vs ...string) predicate.Preemption {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Preemption(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSurrogatephone), v...))
	})
}

// SurrogatephoneGT applies the GT predicate on the "Surrogatephone" field.
func SurrogatephoneGT(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneGTE applies the GTE predicate on the "Surrogatephone" field.
func SurrogatephoneGTE(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneLT applies the LT predicate on the "Surrogatephone" field.
func SurrogatephoneLT(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneLTE applies the LTE predicate on the "Surrogatephone" field.
func SurrogatephoneLTE(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneContains applies the Contains predicate on the "Surrogatephone" field.
func SurrogatephoneContains(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneHasPrefix applies the HasPrefix predicate on the "Surrogatephone" field.
func SurrogatephoneHasPrefix(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneHasSuffix applies the HasSuffix predicate on the "Surrogatephone" field.
func SurrogatephoneHasSuffix(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneEqualFold applies the EqualFold predicate on the "Surrogatephone" field.
func SurrogatephoneEqualFold(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSurrogatephone), v))
	})
}

// SurrogatephoneContainsFold applies the ContainsFold predicate on the "Surrogatephone" field.
func SurrogatephoneContainsFold(v string) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSurrogatephone), v))
	})
}

// HasUserID applies the HasEdge predicate on the "User_ID" edge.
func HasUserID() predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserIDWith applies the HasEdge predicate on the "User_ID" edge with a given conditions (other predicates).
func HasUserIDWith(preds ...predicate.User) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPurposeID applies the HasEdge predicate on the "PurposeID" edge.
func HasPurposeID() predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PurposeIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PurposeIDTable, PurposeIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPurposeIDWith applies the HasEdge predicate on the "PurposeID" edge with a given conditions (other predicates).
func HasPurposeIDWith(preds ...predicate.Purpose) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PurposeIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PurposeIDTable, PurposeIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoomID applies the HasEdge predicate on the "RoomID" edge.
func HasRoomID() predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomIDTable, RoomIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomIDWith applies the HasEdge predicate on the "RoomID" edge with a given conditions (other predicates).
func HasRoomIDWith(preds ...predicate.Roominfo) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomIDTable, RoomIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Preemption) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Preemption) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
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
func Not(p predicate.Preemption) predicate.Preemption {
	return predicate.Preemption(func(s *sql.Selector) {
		p(s.Not())
	})
}
