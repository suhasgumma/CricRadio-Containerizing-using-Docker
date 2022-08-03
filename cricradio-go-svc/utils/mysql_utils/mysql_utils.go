package mysql_utils

import (
	"bookstoreusers-api/utils/errors"
	_ "fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	ErrorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?;"
)

func ParseError(err error) *errors.RestErr {
	sqlerr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")

	}

	switch sqlerr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
