// Code generated by entc, DO NOT EDIT.

package alert

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/predicate"
	"github.com/m-mizutani/alertchain/types"
)

// ID filters vertices based on their ID field.
func ID(id types.AlertID) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id types.AlertID) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id types.AlertID) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...types.AlertID) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...types.AlertID) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id types.AlertID) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id types.AlertID) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id types.AlertID) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id types.AlertID) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Detector applies equality check predicate on the "detector" field. It's identical to DetectorEQ.
func Detector(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetector), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), vc))
	})
}

// Severity applies equality check predicate on the "severity" field. It's identical to SeverityEQ.
func Severity(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeverity), vc))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// DetectedAt applies equality check predicate on the "detected_at" field. It's identical to DetectedAtEQ.
func DetectedAt(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetectedAt), v))
	})
}

// ClosedAt applies equality check predicate on the "closed_at" field. It's identical to ClosedAtEQ.
func ClosedAt(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClosedAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTitle)))
	})
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTitle)))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// DetectorEQ applies the EQ predicate on the "detector" field.
func DetectorEQ(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetector), v))
	})
}

// DetectorNEQ applies the NEQ predicate on the "detector" field.
func DetectorNEQ(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetector), v))
	})
}

// DetectorIn applies the In predicate on the "detector" field.
func DetectorIn(vs ...string) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDetector), v...))
	})
}

// DetectorNotIn applies the NotIn predicate on the "detector" field.
func DetectorNotIn(vs ...string) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDetector), v...))
	})
}

// DetectorGT applies the GT predicate on the "detector" field.
func DetectorGT(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetector), v))
	})
}

// DetectorGTE applies the GTE predicate on the "detector" field.
func DetectorGTE(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetector), v))
	})
}

// DetectorLT applies the LT predicate on the "detector" field.
func DetectorLT(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetector), v))
	})
}

// DetectorLTE applies the LTE predicate on the "detector" field.
func DetectorLTE(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetector), v))
	})
}

// DetectorContains applies the Contains predicate on the "detector" field.
func DetectorContains(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetector), v))
	})
}

// DetectorHasPrefix applies the HasPrefix predicate on the "detector" field.
func DetectorHasPrefix(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetector), v))
	})
}

// DetectorHasSuffix applies the HasSuffix predicate on the "detector" field.
func DetectorHasSuffix(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetector), v))
	})
}

// DetectorIsNil applies the IsNil predicate on the "detector" field.
func DetectorIsNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDetector)))
	})
}

// DetectorNotNil applies the NotNil predicate on the "detector" field.
func DetectorNotNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDetector)))
	})
}

// DetectorEqualFold applies the EqualFold predicate on the "detector" field.
func DetectorEqualFold(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetector), v))
	})
}

// DetectorContainsFold applies the ContainsFold predicate on the "detector" field.
func DetectorContainsFold(v string) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetector), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), vc))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), vc))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...types.AlertStatus) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...types.AlertStatus) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), vc))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), vc))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), vc))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), vc))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), vc))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), vc))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), vc))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), vc))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v types.AlertStatus) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), vc))
	})
}

// SeverityEQ applies the EQ predicate on the "severity" field.
func SeverityEQ(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeverity), vc))
	})
}

// SeverityNEQ applies the NEQ predicate on the "severity" field.
func SeverityNEQ(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeverity), vc))
	})
}

// SeverityIn applies the In predicate on the "severity" field.
func SeverityIn(vs ...types.Severity) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeverity), v...))
	})
}

// SeverityNotIn applies the NotIn predicate on the "severity" field.
func SeverityNotIn(vs ...types.Severity) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeverity), v...))
	})
}

// SeverityGT applies the GT predicate on the "severity" field.
func SeverityGT(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeverity), vc))
	})
}

// SeverityGTE applies the GTE predicate on the "severity" field.
func SeverityGTE(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeverity), vc))
	})
}

// SeverityLT applies the LT predicate on the "severity" field.
func SeverityLT(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeverity), vc))
	})
}

// SeverityLTE applies the LTE predicate on the "severity" field.
func SeverityLTE(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeverity), vc))
	})
}

// SeverityContains applies the Contains predicate on the "severity" field.
func SeverityContains(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeverity), vc))
	})
}

// SeverityHasPrefix applies the HasPrefix predicate on the "severity" field.
func SeverityHasPrefix(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeverity), vc))
	})
}

// SeverityHasSuffix applies the HasSuffix predicate on the "severity" field.
func SeverityHasSuffix(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeverity), vc))
	})
}

// SeverityIsNil applies the IsNil predicate on the "severity" field.
func SeverityIsNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSeverity)))
	})
}

// SeverityNotNil applies the NotNil predicate on the "severity" field.
func SeverityNotNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSeverity)))
	})
}

// SeverityEqualFold applies the EqualFold predicate on the "severity" field.
func SeverityEqualFold(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeverity), vc))
	})
}

// SeverityContainsFold applies the ContainsFold predicate on the "severity" field.
func SeverityContainsFold(v types.Severity) predicate.Alert {
	vc := string(v)
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeverity), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// DetectedAtEQ applies the EQ predicate on the "detected_at" field.
func DetectedAtEQ(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetectedAt), v))
	})
}

// DetectedAtNEQ applies the NEQ predicate on the "detected_at" field.
func DetectedAtNEQ(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetectedAt), v))
	})
}

// DetectedAtIn applies the In predicate on the "detected_at" field.
func DetectedAtIn(vs ...int64) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDetectedAt), v...))
	})
}

// DetectedAtNotIn applies the NotIn predicate on the "detected_at" field.
func DetectedAtNotIn(vs ...int64) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDetectedAt), v...))
	})
}

// DetectedAtGT applies the GT predicate on the "detected_at" field.
func DetectedAtGT(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetectedAt), v))
	})
}

// DetectedAtGTE applies the GTE predicate on the "detected_at" field.
func DetectedAtGTE(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetectedAt), v))
	})
}

// DetectedAtLT applies the LT predicate on the "detected_at" field.
func DetectedAtLT(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetectedAt), v))
	})
}

// DetectedAtLTE applies the LTE predicate on the "detected_at" field.
func DetectedAtLTE(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetectedAt), v))
	})
}

// DetectedAtIsNil applies the IsNil predicate on the "detected_at" field.
func DetectedAtIsNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDetectedAt)))
	})
}

// DetectedAtNotNil applies the NotNil predicate on the "detected_at" field.
func DetectedAtNotNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDetectedAt)))
	})
}

// ClosedAtEQ applies the EQ predicate on the "closed_at" field.
func ClosedAtEQ(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClosedAt), v))
	})
}

// ClosedAtNEQ applies the NEQ predicate on the "closed_at" field.
func ClosedAtNEQ(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClosedAt), v))
	})
}

// ClosedAtIn applies the In predicate on the "closed_at" field.
func ClosedAtIn(vs ...int64) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldClosedAt), v...))
	})
}

// ClosedAtNotIn applies the NotIn predicate on the "closed_at" field.
func ClosedAtNotIn(vs ...int64) predicate.Alert {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Alert(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldClosedAt), v...))
	})
}

// ClosedAtGT applies the GT predicate on the "closed_at" field.
func ClosedAtGT(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClosedAt), v))
	})
}

// ClosedAtGTE applies the GTE predicate on the "closed_at" field.
func ClosedAtGTE(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClosedAt), v))
	})
}

// ClosedAtLT applies the LT predicate on the "closed_at" field.
func ClosedAtLT(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClosedAt), v))
	})
}

// ClosedAtLTE applies the LTE predicate on the "closed_at" field.
func ClosedAtLTE(v int64) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClosedAt), v))
	})
}

// ClosedAtIsNil applies the IsNil predicate on the "closed_at" field.
func ClosedAtIsNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClosedAt)))
	})
}

// ClosedAtNotNil applies the NotNil predicate on the "closed_at" field.
func ClosedAtNotNil() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClosedAt)))
	})
}

// HasAttributes applies the HasEdge predicate on the "attributes" edge.
func HasAttributes() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttributesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttributesTable, AttributesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttributesWith applies the HasEdge predicate on the "attributes" edge with a given conditions (other predicates).
func HasAttributesWith(preds ...predicate.Attribute) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttributesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttributesTable, AttributesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReferences applies the HasEdge predicate on the "references" edge.
func HasReferences() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReferencesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReferencesTable, ReferencesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReferencesWith applies the HasEdge predicate on the "references" edge with a given conditions (other predicates).
func HasReferencesWith(preds ...predicate.Reference) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReferencesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReferencesTable, ReferencesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTaskLogs applies the HasEdge predicate on the "task_logs" edge.
func HasTaskLogs() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskLogsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TaskLogsTable, TaskLogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskLogsWith applies the HasEdge predicate on the "task_logs" edge with a given conditions (other predicates).
func HasTaskLogsWith(preds ...predicate.TaskLog) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskLogsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TaskLogsTable, TaskLogsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasActionLogs applies the HasEdge predicate on the "action_logs" edge.
func HasActionLogs() predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActionLogsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActionLogsTable, ActionLogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActionLogsWith applies the HasEdge predicate on the "action_logs" edge with a given conditions (other predicates).
func HasActionLogsWith(preds ...predicate.ActionLog) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActionLogsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActionLogsTable, ActionLogsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Alert) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Alert) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
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
func Not(p predicate.Alert) predicate.Alert {
	return predicate.Alert(func(s *sql.Selector) {
		p(s.Not())
	})
}
