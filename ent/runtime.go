// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/lbrictson/TinyMonitor/ent/schema"
	"github.com/lbrictson/TinyMonitor/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[3].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[4].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(string)
	// userDescLocked is the schema descriptor for locked field.
	userDescLocked := userFields[5].Descriptor()
	// user.DefaultLocked holds the default value on creation for the locked field.
	user.DefaultLocked = userDescLocked.Default.(bool)
}
