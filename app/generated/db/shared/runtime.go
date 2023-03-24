// Code generated by ent, DO NOT EDIT.

package shared

import (
	"time"

	"github.com/openPanel/core/app/db/schema/shared"
	"github.com/openPanel/core/app/generated/db/shared/kv"
	"github.com/openPanel/core/app/generated/db/shared/node"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	kvMixin := shared.KV{}.Mixin()
	kvMixinFields0 := kvMixin[0].Fields()
	_ = kvMixinFields0
	kvMixinFields1 := kvMixin[1].Fields()
	_ = kvMixinFields1
	kvFields := shared.KV{}.Fields()
	_ = kvFields
	// kvDescCreatedAt is the schema descriptor for created_at field.
	kvDescCreatedAt := kvMixinFields0[0].Descriptor()
	// kv.DefaultCreatedAt holds the default value on creation for the created_at field.
	kv.DefaultCreatedAt = kvDescCreatedAt.Default.(func() time.Time)
	// kvDescUpdatedAt is the schema descriptor for updated_at field.
	kvDescUpdatedAt := kvMixinFields1[0].Descriptor()
	// kv.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	kv.DefaultUpdatedAt = kvDescUpdatedAt.Default.(func() time.Time)
	// kv.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	kv.UpdateDefaultUpdatedAt = kvDescUpdatedAt.UpdateDefault.(func() time.Time)
	nodeMixin := shared.Node{}.Mixin()
	nodeMixinFields0 := nodeMixin[0].Fields()
	_ = nodeMixinFields0
	nodeMixinFields1 := nodeMixin[1].Fields()
	_ = nodeMixinFields1
	nodeFields := shared.Node{}.Fields()
	_ = nodeFields
	// nodeDescCreatedAt is the schema descriptor for created_at field.
	nodeDescCreatedAt := nodeMixinFields0[0].Descriptor()
	// node.DefaultCreatedAt holds the default value on creation for the created_at field.
	node.DefaultCreatedAt = nodeDescCreatedAt.Default.(func() time.Time)
	// nodeDescUpdatedAt is the schema descriptor for updated_at field.
	nodeDescUpdatedAt := nodeMixinFields1[0].Descriptor()
	// node.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	node.DefaultUpdatedAt = nodeDescUpdatedAt.Default.(func() time.Time)
	// node.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	node.UpdateDefaultUpdatedAt = nodeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// nodeDescPort is the schema descriptor for port field.
	nodeDescPort := nodeFields[3].Descriptor()
	// node.DefaultPort holds the default value on creation for the port field.
	node.DefaultPort = nodeDescPort.Default.(int)
}
