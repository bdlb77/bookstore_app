package mysql_utils

import (
	"strings"

	"github.com/bookstore_app/users_api/utils/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const ErrorNoRows = "no rows in result set"

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.HandleNotFound("No record matching given Id.")
		}
		return rest_errors.HandleInternalServerErr("Error parsing Database response.")
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.HandleBadRequestErr("invalid data.")
	}
	return rest_errors.HandleInternalServerErr("Error processing request.")
}
