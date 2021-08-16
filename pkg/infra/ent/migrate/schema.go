// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AlertsColumns holds the columns for the "alerts" table.
	AlertsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "detector", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeString, Default: "new"},
		{Name: "severity", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "detected_at", Type: field.TypeTime, Nullable: true},
		{Name: "closed_at", Type: field.TypeTime, Nullable: true},
	}
	// AlertsTable holds the schema information for the "alerts" table.
	AlertsTable = &schema.Table{
		Name:       "alerts",
		Columns:    AlertsColumns,
		PrimaryKey: []*schema.Column{AlertsColumns[0]},
	}
	// AttributesColumns holds the columns for the "attributes" table.
	AttributesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "context", Type: field.TypeJSON},
		{Name: "alert_attributes", Type: field.TypeString, Nullable: true},
	}
	// AttributesTable holds the schema information for the "attributes" table.
	AttributesTable = &schema.Table{
		Name:       "attributes",
		Columns:    AttributesColumns,
		PrimaryKey: []*schema.Column{AttributesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attributes_alerts_attributes",
				Columns:    []*schema.Column{AttributesColumns[5]},
				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FindingsColumns holds the columns for the "findings" table.
	FindingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "time", Type: field.TypeTime},
		{Name: "source", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "attribute_findings", Type: field.TypeInt, Nullable: true},
	}
	// FindingsTable holds the schema information for the "findings" table.
	FindingsTable = &schema.Table{
		Name:       "findings",
		Columns:    FindingsColumns,
		PrimaryKey: []*schema.Column{FindingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "findings_attributes_findings",
				Columns:    []*schema.Column{FindingsColumns[5]},
				RefColumns: []*schema.Column{AttributesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AlertsTable,
		AttributesTable,
		FindingsTable,
	}
)

func init() {
	AttributesTable.ForeignKeys[0].RefTable = AlertsTable
	FindingsTable.ForeignKeys[0].RefTable = AttributesTable
}
