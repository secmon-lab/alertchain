// Code generated by entc, DO NOT EDIT.

package annotation

const (
	// Label holds the string label denoting the annotation type in the database.
	Label = "annotation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// Table holds the table name of the annotation in the database.
	Table = "annotations"
)

// Columns holds all SQL columns for annotation fields.
var Columns = []string{
	FieldID,
	FieldTimestamp,
	FieldSource,
	FieldName,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "annotations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"attribute_annotations",
	"task_log_annotated",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}