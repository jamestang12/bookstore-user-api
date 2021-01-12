package mysql_utils

import (
	"strings"

	"../errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	// Convert err MySQL error
	sqlErr, ok := err.(*mysql.MySQLError)
	// Check if is a vaild MySQL error
	if !ok {
		// Check if it contain the string no rows in result set than return a 404 error
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewBadNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	// Switch case for vaild MySQL error
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
