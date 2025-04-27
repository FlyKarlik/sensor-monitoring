package dao

import (
	"database/sql"
	"time"
)

func toNullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: *value, Valid: true}
}

func toNullBool(value *bool) sql.NullBool {
	if value == nil {
		return sql.NullBool{Bool: false, Valid: false}
	}
	return sql.NullBool{Bool: *value, Valid: true}
}

func toNullInt64(value *int) sql.NullInt64 {
	if value == nil {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: int64(*value), Valid: true}
}

func fromNullInt64(value sql.NullInt64) *int {
	if !value.Valid {
		return nil
	}
	valueInt := int(value.Int64)
	return &valueInt
}

func fromNullBool(value sql.NullBool) *bool {
	if !value.Valid {
		return nil
	}
	return &value.Bool
}

func fromNullString(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}

func fromNullTime(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}
	return &value.Time
}
