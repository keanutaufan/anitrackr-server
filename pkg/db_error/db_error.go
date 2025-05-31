package db_error

import (
	"database/sql"
	"errors"
	"strings"
)

func IsSqlStateError(err error, code string) bool {
	if err != nil {
		return strings.Contains(err.Error(), "SQLSTATE="+code)
	}
	return false
}

func IsNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func IsUniqueViolation(err error) bool {
	return IsSqlStateError(err, "23505")
}
