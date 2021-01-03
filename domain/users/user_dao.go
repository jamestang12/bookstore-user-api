package users

import (
	"../../utils/errors"
	"fmt"
	"../../utils/date_utils"
	"../../datasources/mysql/users_db"
	"strings"
	
)

const(
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"

)

var(
	usersDB = make(map[int64]*User)
)


// Set it to pointer which when setting the user info, it will 
// pass it into the real user instead killing it after the func is done
func (user *User)Get()  *errors.RestErr{

	if err := users_db.Client.Ping(); err != nil{
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil{
		return errors.NewBadNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User)Save() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	// Close the stmt after the request to the db is done
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil{
		if strings.Contains(err.Error(), "email_UNIQUE"){
			return errors.NewInternalServerError(fmt.Sprintf("email: %s alreayd exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))		
	}	

	// result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	// if err != nil {

	// }

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))		
	}
	
	user.Id = userId

	return nil
}