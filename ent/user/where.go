// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/TinyMonitor/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// APIKey applies equality check predicate on the "api_key" field. It's identical to APIKeyEQ.
func APIKey(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAPIKey), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRole), v))
	})
}

// Locked applies equality check predicate on the "locked" field. It's identical to LockedEQ.
func Locked(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocked), v))
	})
}

// LockedUntil applies equality check predicate on the "locked_until" field. It's identical to LockedUntilEQ.
func LockedUntil(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLockedUntil), v))
	})
}

// APIKeyEQ applies the EQ predicate on the "api_key" field.
func APIKeyEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAPIKey), v))
	})
}

// APIKeyNEQ applies the NEQ predicate on the "api_key" field.
func APIKeyNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAPIKey), v))
	})
}

// APIKeyIn applies the In predicate on the "api_key" field.
func APIKeyIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAPIKey), v...))
	})
}

// APIKeyNotIn applies the NotIn predicate on the "api_key" field.
func APIKeyNotIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAPIKey), v...))
	})
}

// APIKeyGT applies the GT predicate on the "api_key" field.
func APIKeyGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAPIKey), v))
	})
}

// APIKeyGTE applies the GTE predicate on the "api_key" field.
func APIKeyGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAPIKey), v))
	})
}

// APIKeyLT applies the LT predicate on the "api_key" field.
func APIKeyLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAPIKey), v))
	})
}

// APIKeyLTE applies the LTE predicate on the "api_key" field.
func APIKeyLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAPIKey), v))
	})
}

// APIKeyContains applies the Contains predicate on the "api_key" field.
func APIKeyContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAPIKey), v))
	})
}

// APIKeyHasPrefix applies the HasPrefix predicate on the "api_key" field.
func APIKeyHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAPIKey), v))
	})
}

// APIKeyHasSuffix applies the HasSuffix predicate on the "api_key" field.
func APIKeyHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAPIKey), v))
	})
}

// APIKeyEqualFold applies the EqualFold predicate on the "api_key" field.
func APIKeyEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAPIKey), v))
	})
}

// APIKeyContainsFold applies the ContainsFold predicate on the "api_key" field.
func APIKeyContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAPIKey), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRole), v))
	})
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRole), v))
	})
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRole), v...))
	})
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRole), v...))
	})
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRole), v))
	})
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRole), v))
	})
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRole), v))
	})
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRole), v))
	})
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRole), v))
	})
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRole), v))
	})
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRole), v))
	})
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRole), v))
	})
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRole), v))
	})
}

// LockedEQ applies the EQ predicate on the "locked" field.
func LockedEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocked), v))
	})
}

// LockedNEQ applies the NEQ predicate on the "locked" field.
func LockedNEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocked), v))
	})
}

// LockedUntilEQ applies the EQ predicate on the "locked_until" field.
func LockedUntilEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLockedUntil), v))
	})
}

// LockedUntilNEQ applies the NEQ predicate on the "locked_until" field.
func LockedUntilNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLockedUntil), v))
	})
}

// LockedUntilIn applies the In predicate on the "locked_until" field.
func LockedUntilIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLockedUntil), v...))
	})
}

// LockedUntilNotIn applies the NotIn predicate on the "locked_until" field.
func LockedUntilNotIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLockedUntil), v...))
	})
}

// LockedUntilGT applies the GT predicate on the "locked_until" field.
func LockedUntilGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLockedUntil), v))
	})
}

// LockedUntilGTE applies the GTE predicate on the "locked_until" field.
func LockedUntilGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLockedUntil), v))
	})
}

// LockedUntilLT applies the LT predicate on the "locked_until" field.
func LockedUntilLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLockedUntil), v))
	})
}

// LockedUntilLTE applies the LTE predicate on the "locked_until" field.
func LockedUntilLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLockedUntil), v))
	})
}

// LockedUntilIsNil applies the IsNil predicate on the "locked_until" field.
func LockedUntilIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLockedUntil)))
	})
}

// LockedUntilNotNil applies the NotNil predicate on the "locked_until" field.
func LockedUntilNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLockedUntil)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
