// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"plant-watering/ent/waterlog"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// WaterLog is the model entity for the WaterLog schema.
type WaterLog struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Seconds holds the value of the "seconds" field.
	Seconds int32 `json:"seconds,omitempty"`
	// Channel holds the value of the "channel" field.
	Channel string `json:"channel,omitempty"`
	// Manual holds the value of the "manual" field.
	Manual bool `json:"manual,omitempty"`
	// Time holds the value of the "time" field.
	Time         time.Time `json:"time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WaterLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case waterlog.FieldManual:
			values[i] = new(sql.NullBool)
		case waterlog.FieldSeconds:
			values[i] = new(sql.NullInt64)
		case waterlog.FieldChannel:
			values[i] = new(sql.NullString)
		case waterlog.FieldTime:
			values[i] = new(sql.NullTime)
		case waterlog.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WaterLog fields.
func (wl *WaterLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case waterlog.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				wl.ID = *value
			}
		case waterlog.FieldSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field seconds", values[i])
			} else if value.Valid {
				wl.Seconds = int32(value.Int64)
			}
		case waterlog.FieldChannel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel", values[i])
			} else if value.Valid {
				wl.Channel = value.String
			}
		case waterlog.FieldManual:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field manual", values[i])
			} else if value.Valid {
				wl.Manual = value.Bool
			}
		case waterlog.FieldTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				wl.Time = value.Time
			}
		default:
			wl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WaterLog.
// This includes values selected through modifiers, order, etc.
func (wl *WaterLog) Value(name string) (ent.Value, error) {
	return wl.selectValues.Get(name)
}

// Update returns a builder for updating this WaterLog.
// Note that you need to call WaterLog.Unwrap() before calling this method if this WaterLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (wl *WaterLog) Update() *WaterLogUpdateOne {
	return NewWaterLogClient(wl.config).UpdateOne(wl)
}

// Unwrap unwraps the WaterLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wl *WaterLog) Unwrap() *WaterLog {
	_tx, ok := wl.config.driver.(*txDriver)
	if !ok {
		panic("ent: WaterLog is not a transactional entity")
	}
	wl.config.driver = _tx.drv
	return wl
}

// String implements the fmt.Stringer.
func (wl *WaterLog) String() string {
	var builder strings.Builder
	builder.WriteString("WaterLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wl.ID))
	builder.WriteString("seconds=")
	builder.WriteString(fmt.Sprintf("%v", wl.Seconds))
	builder.WriteString(", ")
	builder.WriteString("channel=")
	builder.WriteString(wl.Channel)
	builder.WriteString(", ")
	builder.WriteString("manual=")
	builder.WriteString(fmt.Sprintf("%v", wl.Manual))
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(wl.Time.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// WaterLogs is a parsable slice of WaterLog.
type WaterLogs []*WaterLog
