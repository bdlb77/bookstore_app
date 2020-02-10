package users

import (
	"fmt"

	"github.com/bookstore_app/users_api/datasources/mysql/users_db"
	"github.com/bookstore_app/users_api/utils/date_utils"
	"github.com/bookstore_app/users_api/utils/mysql_utils"
	"github.com/bookstore_app/users_api/utils/rest_errors"
)

const (
	indexUniqEmail  = "email_UNIQUE"
	noRows          = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE  id = ?"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *rest_errors.RestErr {
	stmt, clientErr := users_db.Client.Prepare(queryGetUser)
	if clientErr != nil {
		return rest_errors.HandleInternalServerErr(fmt.Sprintf("Error went trying to read to DB. %s", clientErr.Error()))
	}

	defer stmt.Close()

	queryResult := stmt.QueryRow(user.Id)
	if err := queryResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Destroy() *rest_errors.RestErr {
	result := userDB[user.Id]

	if result == nil {
		return rest_errors.HandleNotFound(fmt.Sprintf("User with id: %d not found.", user.Id))
	}
	delete(userDB, user.Id)
	return nil
}

func (user *User) Save() *rest_errors.RestErr {
	stmt, clientErr := users_db.Client.Prepare(queryInsertUser)
	if clientErr != nil {
		return rest_errors.HandleInternalServerErr(clientErr.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()

	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *rest_errors.RestErr {
	stmt, clientErr := users_db.Client.Prepare(queryUpdateUser)
	if clientErr != nil {
		return rest_errors.HandleInternalServerErr(clientErr.Error())
	}
	defer stmt.Close()

	_, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
