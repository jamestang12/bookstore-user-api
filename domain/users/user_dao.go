package users

import (
	"../../utils/errors"
	"fmt"
	// "../../utils/date_utils"
	"../../datasources/mysql/users_db"
	"../../utils/mysql_utils"
)

const(   			   
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?,?,?,?,?,?);"
	quertGetUser = "SELECT id,first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	errorNoRows = "no rows in result set"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindByStatus  = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"

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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
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

	//user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
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

func (user *User)Update() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err1 != nil{
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User)Delete() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err1 := stmt.Exec(user.Id)
	if err1 != nil{
		return mysql_utils.ParseError(err1)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr){
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		return  nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err1 := stmt.Query(status)
	if err1 != nil{
		return nil, errors.NewInternalServerError(err1.Error())
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next(){
		var user User
		// Using & while doing scan so that it take the actual value instead of the copy of that value
		if err2 := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err2 != nil{
			return nil, mysql_utils.ParseError(err2)
		}
		result = append(result, user)
	}
	if len(result) == 0{
		return nil, errors.NewBadNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return result, nil
}