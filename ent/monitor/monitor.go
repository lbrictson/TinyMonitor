// Code generated by ent, DO NOT EDIT.

package monitor

import (
	"time"
)

const (
	// Label holds the string label denoting the monitor type in the database.
	Label = "monitor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLastCheckedAt holds the string denoting the last_checked_at field in the database.
	FieldLastCheckedAt = "last_checked_at"
	// FieldStatusLastChangedAt holds the string denoting the status_last_changed_at field in the database.
	FieldStatusLastChangedAt = "status_last_changed_at"
	// FieldMonitorType holds the string denoting the monitor_type field in the database.
	FieldMonitorType = "monitor_type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldIntervalSeconds holds the string denoting the interval_seconds field in the database.
	FieldIntervalSeconds = "interval_seconds"
	// FieldPaused holds the string denoting the paused field in the database.
	FieldPaused = "paused"
	// Table holds the table name of the monitor in the database.
	Table = "monitors"
)

// Columns holds all SQL columns for monitor fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldStatus,
	FieldLastCheckedAt,
	FieldStatusLastChangedAt,
	FieldMonitorType,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldConfig,
	FieldIntervalSeconds,
	FieldPaused,
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
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultStatusLastChangedAt holds the default value on creation for the "status_last_changed_at" field.
	DefaultStatusLastChangedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIntervalSeconds holds the default value on creation for the "interval_seconds" field.
	DefaultIntervalSeconds int
	// DefaultPaused holds the default value on creation for the "paused" field.
	DefaultPaused bool
)
