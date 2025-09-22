package model

import (
	"mocau-backend/common"
	"database/sql/driver"
	"errors"
	"fmt"
)

const EntityName = "User"

type UserRole int

const (
	RoleUser UserRole = 1 << iota
	RoleAdmin
)

func (role UserRole) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	default:
		return "user"
	}
}

func (role *UserRole) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var r UserRole

	roleValue := string(bytes)

	if roleValue == "user" {
		r = RoleUser
	} else if roleValue == "admin" {
		r = RoleAdmin
	}

	*role = r

	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}
	return role.String(), nil
}

func (role *UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}

type User struct {
	common.SQLModel
	Username  string   `json:"username" gorm:"column:username;"`
	Password  string   `json:"-" gorm:"column:password;"`
	Email     string   `json:"email" gorm:"column:email;"`
	FullName  string   `json:"full_name" gorm:"column:full_name;"`
	Phone     string   `json:"phone" gorm:"column:phone;"`
	Role      UserRole `json:"role" gorm:"column:role;"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role.String()
}

func (User) TableName() string {
	return "Users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Username        string `json:"username" gorm:"column:username;"`
	Password        string `json:"password" gorm:"column:password;"`
	Email           string `json:"email" gorm:"column:email;"`
	FullName        string `json:"full_name" gorm:"column:full_name;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
	Role            string `json:"-" gorm:"column:role;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Username string `json:"username" form:"username" gorm:"column:username;"`
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

//type Account struct {
//	AccessToken  *tokenprovider.Token `json:"access_token"`
//	RefreshToken *tokenprovider.Token `json:"refresh_token"`
//}
//
//func NewAccount(at, rt *tokenprovider.Token) *Account {
//	return &Account{
//		AccessToken:  at,
//		RefreshToken: rt,
//	}
//}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("username/email or password invalid"),
		"username/email or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrUsernameExisted = common.NewCustomError(
		errors.New("username has already existed"),
		"username has already existed",
		"ErrUsernameExisted",
	)
)
