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
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCurrentDownReason holds the string denoting the current_down_reason field in the database.
	FieldCurrentDownReason = "current_down_reason"
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
	// FieldFailureCount holds the string denoting the failure_count field in the database.
	FieldFailureCount = "failure_count"
	// FieldSuccessCount holds the string denoting the success_count field in the database.
	FieldSuccessCount = "success_count"
	// FieldSuccessThreshold holds the string denoting the success_threshold field in the database.
	FieldSuccessThreshold = "success_threshold"
	// FieldFailureThreshold holds the string denoting the failure_threshold field in the database.
	FieldFailureThreshold = "failure_threshold"
	// Table holds the table name of the monitor in the database.
	Table = "monitors"
)

// Columns holds all SQL columns for monitor fields.
var Columns = []string{
	FieldID,
	FieldDescription,
	FieldCurrentDownReason,
	FieldStatus,
	FieldLastCheckedAt,
	FieldStatusLastChangedAt,
	FieldMonitorType,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldConfig,
	FieldIntervalSeconds,
	FieldPaused,
	FieldFailureCount,
	FieldSuccessCount,
	FieldSuccessThreshold,
	FieldFailureThreshold,
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
	// DefaultCurrentDownReason holds the default value on creation for the "current_down_reason" field.
	DefaultCurrentDownReason string
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
	// DefaultFailureCount holds the default value on creation for the "failure_count" field.
	DefaultFailureCount int
	// DefaultSuccessCount holds the default value on creation for the "success_count" field.
	DefaultSuccessCount int
	// DefaultSuccessThreshold holds the default value on creation for the "success_threshold" field.
	DefaultSuccessThreshold int
	// DefaultFailureThreshold holds the default value on creation for the "failure_threshold" field.
	DefaultFailureThreshold int
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
