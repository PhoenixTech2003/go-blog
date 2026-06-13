package api

import (
	"database/sql"
	"time"
)

func nullTimeValue(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}
