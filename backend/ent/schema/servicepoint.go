package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// ServicePoint holds the schema definition for the ServicePoint entity.
type ServicePoint struct {
	ent.Schema
}

// Fields of the ServicePoint.
func (ServicePoint) Fields() []ent.Field {
	return []ent.Field{
		field.String("COUNTER_NUMBER").
			Unique().
			NotEmpty(),
	}
}

// Edges of the ServicePoint.
func (ServicePoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("from", Bookborrow.Type).
			StorageKey(edge.Column("SERVICEPOINT_ID")),
		edge.To("servicepoint", Booking.Type).
			StorageKey(edge.Column("SERVICEPOINT_ID")),
	}
}
