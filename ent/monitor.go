// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/TinyMonitor/ent/monitor"
)

// Monitor is the model entity for the Monitor schema.
type Monitor struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CurrentDownReason holds the value of the "current_down_reason" field.
	CurrentDownReason string `json:"current_down_reason,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// LastCheckedAt holds the value of the "last_checked_at" field.
	LastCheckedAt *time.Time `json:"last_checked_at,omitempty"`
	// StatusLastChangedAt holds the value of the "status_last_changed_at" field.
	StatusLastChangedAt time.Time `json:"status_last_changed_at,omitempty"`
	// MonitorType holds the value of the "monitor_type" field.
	MonitorType string `json:"monitor_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Config holds the value of the "config" field.
	Config map[string]interface{} `json:"config,omitempty"`
	// IntervalSeconds holds the value of the "interval_seconds" field.
	IntervalSeconds int `json:"interval_seconds,omitempty"`
	// Paused holds the value of the "paused" field.
	Paused bool `json:"paused,omitempty"`
	// FailureCount holds the value of the "failure_count" field.
	FailureCount int `json:"failure_count,omitempty"`
	// SuccessThreshold holds the value of the "success_threshold" field.
	SuccessThreshold int `json:"success_threshold,omitempty"`
	// FailureThreshold holds the value of the "failure_threshold" field.
	FailureThreshold int `json:"failure_threshold,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Monitor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case monitor.FieldConfig:
			values[i] = new([]byte)
		case monitor.FieldPaused:
			values[i] = new(sql.NullBool)
		case monitor.FieldIntervalSeconds, monitor.FieldFailureCount, monitor.FieldSuccessThreshold, monitor.FieldFailureThreshold:
			values[i] = new(sql.NullInt64)
		case monitor.FieldID, monitor.FieldDescription, monitor.FieldCurrentDownReason, monitor.FieldStatus, monitor.FieldMonitorType:
			values[i] = new(sql.NullString)
		case monitor.FieldLastCheckedAt, monitor.FieldStatusLastChangedAt, monitor.FieldCreatedAt, monitor.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Monitor", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Monitor fields.
func (m *Monitor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case monitor.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				m.ID = value.String
			}
		case monitor.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				m.Description = value.String
			}
		case monitor.FieldCurrentDownReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field current_down_reason", values[i])
			} else if value.Valid {
				m.CurrentDownReason = value.String
			}
		case monitor.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = value.String
			}
		case monitor.FieldLastCheckedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_checked_at", values[i])
			} else if value.Valid {
				m.LastCheckedAt = new(time.Time)
				*m.LastCheckedAt = value.Time
			}
		case monitor.FieldStatusLastChangedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field status_last_changed_at", values[i])
			} else if value.Valid {
				m.StatusLastChangedAt = value.Time
			}
		case monitor.FieldMonitorType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field monitor_type", values[i])
			} else if value.Valid {
				m.MonitorType = value.String
			}
		case monitor.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case monitor.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case monitor.FieldConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Config); err != nil {
					return fmt.Errorf("unmarshal field config: %w", err)
				}
			}
		case monitor.FieldIntervalSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field interval_seconds", values[i])
			} else if value.Valid {
				m.IntervalSeconds = int(value.Int64)
			}
		case monitor.FieldPaused:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field paused", values[i])
			} else if value.Valid {
				m.Paused = value.Bool
			}
		case monitor.FieldFailureCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field failure_count", values[i])
			} else if value.Valid {
				m.FailureCount = int(value.Int64)
			}
		case monitor.FieldSuccessThreshold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field success_threshold", values[i])
			} else if value.Valid {
				m.SuccessThreshold = int(value.Int64)
			}
		case monitor.FieldFailureThreshold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field failure_threshold", values[i])
			} else if value.Valid {
				m.FailureThreshold = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Monitor.
// Note that you need to call Monitor.Unwrap() before calling this method if this Monitor
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Monitor) Update() *MonitorUpdateOne {
	return (&MonitorClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Monitor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Monitor) Unwrap() *Monitor {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Monitor is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Monitor) String() string {
	var builder strings.Builder
	builder.WriteString("Monitor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("description=")
	builder.WriteString(m.Description)
	builder.WriteString(", ")
	builder.WriteString("current_down_reason=")
	builder.WriteString(m.CurrentDownReason)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(m.Status)
	builder.WriteString(", ")
	if v := m.LastCheckedAt; v != nil {
		builder.WriteString("last_checked_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status_last_changed_at=")
	builder.WriteString(m.StatusLastChangedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("monitor_type=")
	builder.WriteString(m.MonitorType)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("config=")
	builder.WriteString(fmt.Sprintf("%v", m.Config))
	builder.WriteString(", ")
	builder.WriteString("interval_seconds=")
	builder.WriteString(fmt.Sprintf("%v", m.IntervalSeconds))
	builder.WriteString(", ")
	builder.WriteString("paused=")
	builder.WriteString(fmt.Sprintf("%v", m.Paused))
	builder.WriteString(", ")
	builder.WriteString("failure_count=")
	builder.WriteString(fmt.Sprintf("%v", m.FailureCount))
	builder.WriteString(", ")
	builder.WriteString("success_threshold=")
	builder.WriteString(fmt.Sprintf("%v", m.SuccessThreshold))
	builder.WriteString(", ")
	builder.WriteString("failure_threshold=")
	builder.WriteString(fmt.Sprintf("%v", m.FailureThreshold))
	builder.WriteByte(')')
	return builder.String()
}

// Monitors is a parsable slice of Monitor.
type Monitors []*Monitor

func (m Monitors) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
