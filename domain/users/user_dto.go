package users

import(
	"../../utils/errors"
	"strings"
)

type User struct{
	
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))

	if user.Email == ""{
		return errors.NewBadRequestError("Invaild email address")
	}

	return  nil
}
