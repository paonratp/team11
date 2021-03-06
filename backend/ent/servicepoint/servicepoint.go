// Code generated by entc, DO NOT EDIT.

package servicepoint

const (
	// Label holds the string label denoting the servicepoint type in the database.
	Label = "service_point"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCOUNTERNUMBER holds the string denoting the counter_number field in the database.
	FieldCOUNTERNUMBER = "counter_number"

	// EdgeFrom holds the string denoting the from edge name in mutations.
	EdgeFrom = "from"
	// EdgeServicepoint holds the string denoting the servicepoint edge name in mutations.
	EdgeServicepoint = "servicepoint"

	// Table holds the table name of the servicepoint in the database.
	Table = "service_points"
	// FromTable is the table the holds the from relation/edge.
	FromTable = "bookborrows"
	// FromInverseTable is the table name for the Bookborrow entity.
	// It exists in this package in order to avoid circular dependency with the "bookborrow" package.
	FromInverseTable = "bookborrows"
	// FromColumn is the table column denoting the from relation/edge.
	FromColumn = "SERVICEPOINT_ID"
	// ServicepointTable is the table the holds the servicepoint relation/edge.
	ServicepointTable = "bookings"
	// ServicepointInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	ServicepointInverseTable = "bookings"
	// ServicepointColumn is the table column denoting the servicepoint relation/edge.
	ServicepointColumn = "SERVICEPOINT_ID"
)

// Columns holds all SQL columns for servicepoint fields.
var Columns = []string{
	FieldID,
	FieldCOUNTERNUMBER,
}

var (
	// COUNTERNUMBERValidator is a validator for the "COUNTER_NUMBER" field. It is called by the builders before save.
	COUNTERNUMBERValidator func(string) error
)
