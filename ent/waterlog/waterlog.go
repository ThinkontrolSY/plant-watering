// Code generated by ent, DO NOT EDIT.

package waterlog

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the waterlog type in the database.
	Label = "water_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSeconds holds the string denoting the seconds field in the database.
	FieldSeconds = "seconds"
	// FieldChannel holds the string denoting the channel field in the database.
	FieldChannel = "channel"
	// FieldManual holds the string denoting the manual field in the database.
	FieldManual = "manual"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// Table holds the table name of the waterlog in the database.
	Table = "water_logs"
)

// Columns holds all SQL columns for waterlog fields.
var Columns = []string{
	FieldID,
	FieldSeconds,
	FieldChannel,
	FieldManual,
	FieldTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTime holds the default value on creation for the "time" field.
	DefaultTime func() time.Time
)

// OrderOption defines the ordering options for the WaterLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySeconds orders the results by the seconds field.
func BySeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeconds, opts...).ToFunc()
}

// ByChannel orders the results by the channel field.
func ByChannel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannel, opts...).ToFunc()
}

// ByManual orders the results by the manual field.
func ByManual(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManual, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}
