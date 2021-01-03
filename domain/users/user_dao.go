package users

import (
	"../../utils/errors"
	"fmt"
	"../../utils/date_utils"
	"../../datasources/mysql/users_db"
	"../../utils/mysql_utils"
)

const(
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
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
	// Insert data into the uer struct
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		// Calling the error handle func in mysql_utils 
		return mysql_utils.ParseError(getErr)
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil{


		// Calling the error handle func in mysql_utils 
		return mysql_utils.ParseError(saveErr)


		// sqlErr, ok := saveErr.(*mysql.MySQLError)
		// if  !ok{
		// 	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
		// }

		// fmt.Println(sqlErr.Number)
		// fmt.Println(sqlErr.Message)
		// // Switch case depend by the sql number
		// switch sqlErr.Number {
		// case 1062:
		// 	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", user.Email))
		// }
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))


		
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