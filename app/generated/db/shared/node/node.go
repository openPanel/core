// Code generated by ent, DO NOT EDIT.

package node

import (
	"time"
)

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// Table holds the table name of the node in the database.
	Table = "nodes"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldIP,
	FieldPort,
	FieldComment,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultPort holds the default value on creation for the "port" field.
	DefaultPort int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
