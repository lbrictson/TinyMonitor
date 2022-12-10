// Code generated by ent, DO NOT EDIT.

package monitor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/TinyMonitor/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// LastCheckedAt applies equality check predicate on the "last_checked_at" field. It's identical to LastCheckedAtEQ.
func LastCheckedAt(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastCheckedAt), v))
	})
}

// StatusLastChangedAt applies equality check predicate on the "status_last_changed_at" field. It's identical to StatusLastChangedAtEQ.
func StatusLastChangedAt(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusLastChangedAt), v))
	})
}

// MonitorType applies equality check predicate on the "monitor_type" field. It's identical to MonitorTypeEQ.
func MonitorType(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMonitorType), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// IntervalSeconds applies equality check predicate on the "interval_seconds" field. It's identical to IntervalSecondsEQ.
func IntervalSeconds(v int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntervalSeconds), v))
	})
}

// Paused applies equality check predicate on the "paused" field. It's identical to PausedEQ.
func Paused(v bool) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaused), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// LastCheckedAtEQ applies the EQ predicate on the "last_checked_at" field.
func LastCheckedAtEQ(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastCheckedAt), v))
	})
}

// LastCheckedAtNEQ applies the NEQ predicate on the "last_checked_at" field.
func LastCheckedAtNEQ(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastCheckedAt), v))
	})
}

// LastCheckedAtIn applies the In predicate on the "last_checked_at" field.
func LastCheckedAtIn(vs ...time.Time) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastCheckedAt), v...))
	})
}

// LastCheckedAtNotIn applies the NotIn predicate on the "last_checked_at" field.
func LastCheckedAtNotIn(vs ...time.Time) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastCheckedAt), v...))
	})
}

// LastCheckedAtGT applies the GT predicate on the "last_checked_at" field.
func LastCheckedAtGT(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastCheckedAt), v))
	})
}

// LastCheckedAtGTE applies the GTE predicate on the "last_checked_at" field.
func LastCheckedAtGTE(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastCheckedAt), v))
	})
}

// LastCheckedAtLT applies the LT predicate on the "last_checked_at" field.
func LastCheckedAtLT(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastCheckedAt), v))
	})
}

// LastCheckedAtLTE applies the LTE predicate on the "last_checked_at" field.
func LastCheckedAtLTE(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastCheckedAt), v))
	})
}

// LastCheckedAtIsNil applies the IsNil predicate on the "last_checked_at" field.
func LastCheckedAtIsNil() predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastCheckedAt)))
	})
}

// LastCheckedAtNotNil applies the NotNil predicate on the "last_checked_at" field.
func LastCheckedAtNotNil() predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastCheckedAt)))
	})
}

// StatusLastChangedAtEQ applies the EQ predicate on the "status_last_changed_at" field.
func StatusLastChangedAtEQ(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusLastChangedAt), v))
	})
}

// StatusLastChangedAtNEQ applies the NEQ predicate on the "status_last_changed_at" field.
func StatusLastChangedAtNEQ(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusLastChangedAt), v))
	})
}

// StatusLastChangedAtIn applies the In predicate on the "status_last_changed_at" field.
func StatusLastChangedAtIn(vs ...time.Time) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatusLastChangedAt), v...))
	})
}

// StatusLastChangedAtNotIn applies the NotIn predicate on the "status_last_changed_at" field.
func StatusLastChangedAtNotIn(vs ...time.Time) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatusLastChangedAt), v...))
	})
}

// StatusLastChangedAtGT applies the GT predicate on the "status_last_changed_at" field.
func StatusLastChangedAtGT(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatusLastChangedAt), v))
	})
}

// StatusLastChangedAtGTE applies the GTE predicate on the "status_last_changed_at" field.
func StatusLastChangedAtGTE(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatusLastChangedAt), v))
	})
}

// StatusLastChangedAtLT applies the LT predicate on the "status_last_changed_at" field.
func StatusLastChangedAtLT(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatusLastChangedAt), v))
	})
}

// StatusLastChangedAtLTE applies the LTE predicate on the "status_last_changed_at" field.
func StatusLastChangedAtLTE(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatusLastChangedAt), v))
	})
}

// MonitorTypeEQ applies the EQ predicate on the "monitor_type" field.
func MonitorTypeEQ(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeNEQ applies the NEQ predicate on the "monitor_type" field.
func MonitorTypeNEQ(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeIn applies the In predicate on the "monitor_type" field.
func MonitorTypeIn(vs ...string) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMonitorType), v...))
	})
}

// MonitorTypeNotIn applies the NotIn predicate on the "monitor_type" field.
func MonitorTypeNotIn(vs ...string) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMonitorType), v...))
	})
}

// MonitorTypeGT applies the GT predicate on the "monitor_type" field.
func MonitorTypeGT(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeGTE applies the GTE predicate on the "monitor_type" field.
func MonitorTypeGTE(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeLT applies the LT predicate on the "monitor_type" field.
func MonitorTypeLT(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeLTE applies the LTE predicate on the "monitor_type" field.
func MonitorTypeLTE(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeContains applies the Contains predicate on the "monitor_type" field.
func MonitorTypeContains(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeHasPrefix applies the HasPrefix predicate on the "monitor_type" field.
func MonitorTypeHasPrefix(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeHasSuffix applies the HasSuffix predicate on the "monitor_type" field.
func MonitorTypeHasSuffix(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeEqualFold applies the EqualFold predicate on the "monitor_type" field.
func MonitorTypeEqualFold(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMonitorType), v))
	})
}

// MonitorTypeContainsFold applies the ContainsFold predicate on the "monitor_type" field.
func MonitorTypeContainsFold(v string) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMonitorType), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// IntervalSecondsEQ applies the EQ predicate on the "interval_seconds" field.
func IntervalSecondsEQ(v int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntervalSeconds), v))
	})
}

// IntervalSecondsNEQ applies the NEQ predicate on the "interval_seconds" field.
func IntervalSecondsNEQ(v int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIntervalSeconds), v))
	})
}

// IntervalSecondsIn applies the In predicate on the "interval_seconds" field.
func IntervalSecondsIn(vs ...int) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIntervalSeconds), v...))
	})
}

// IntervalSecondsNotIn applies the NotIn predicate on the "interval_seconds" field.
func IntervalSecondsNotIn(vs ...int) predicate.Monitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIntervalSeconds), v...))
	})
}

// IntervalSecondsGT applies the GT predicate on the "interval_seconds" field.
func IntervalSecondsGT(v int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIntervalSeconds), v))
	})
}

// IntervalSecondsGTE applies the GTE predicate on the "interval_seconds" field.
func IntervalSecondsGTE(v int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIntervalSeconds), v))
	})
}

// IntervalSecondsLT applies the LT predicate on the "interval_seconds" field.
func IntervalSecondsLT(v int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIntervalSeconds), v))
	})
}

// IntervalSecondsLTE applies the LTE predicate on the "interval_seconds" field.
func IntervalSecondsLTE(v int) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIntervalSeconds), v))
	})
}

// PausedEQ applies the EQ predicate on the "paused" field.
func PausedEQ(v bool) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaused), v))
	})
}

// PausedNEQ applies the NEQ predicate on the "paused" field.
func PausedNEQ(v bool) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaused), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Monitor) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Monitor) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
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
func Not(p predicate.Monitor) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		p(s.Not())
	})
}
