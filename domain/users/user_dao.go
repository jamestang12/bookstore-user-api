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
	indexUniqueEmail = "email_UNIQUE"
	quertGetUser = "SELECT id,first_name, last_name, email, date_created FROM users WHERE id = ?;"
	errorNoRows = "no rows in result set"
)

var(
	usersDB = make(map[int64]*User)
)


// Set it to pointer which when setting the user info, it will 
// pass it into the real user instead killing it after the func is done
func (user *User)Get()  *errors.RestErr{

	stmt, err := users_db.Client.Prepare(quertGetUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	// Close the stmt after the request to the db is done
	defer stmt.Close()

	// No need to close stmt since only getting one row
	result := stmt.QueryRow(user.Id)
	// Insert data into the suer struct
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows){
			return errors.NewInternalServerError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

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
		if strings.Contains(err.Error(), indexUniqueEmail){
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