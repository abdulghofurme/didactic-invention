package helper

import "database/sql"

func SQLTimeNullToString(timeNull *sql.NullTime) string {
	timeString := timeNull.Time.String()
	if !timeNull.Valid {
		return ""
	}
	return timeString
}
