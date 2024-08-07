package domain

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserID represents the user unique identifier
type UserID string

func (u UserID) Validate() (*UserID, error) {
	_, err := uuid.Parse(string(u))
	if err != nil {
		return nil, fmt.Errorf("%w = %s", InvalidUserID, err.Error())
	}

	return &u, nil
}

func (u UserID) ID() string {
	return string(u)
}

// UserName represents the user name
type UserName string

func (a UserName) Validate() (*UserName, error) {
	if strings.TrimSpace(string(a)) == "" {
		return nil, fmt.Errorf("%w: %s", InvalidUserName, "the name is invalid")
	}

	return &a, nil
}

func (a UserName) Name() string {
	return string(a)
}

// UserPassword represents the user password
type UserPassword string

func (a UserPassword) Validate() (*UserPassword, error) {
	// TODO: validate length and characters
	if strings.TrimSpace(string(a)) == "" {
		return nil, fmt.Errorf("%w: %s", InvalidUserPassword, "the last password is invalid")
	}

	return &a, nil
}

func (a UserPassword) Encrypt() *UserPassword {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(string(a)), bcrypt.DefaultCost)
	hashPassword := UserPassword(string(bytes))

	return &hashPassword
}

func (a UserPassword) Password() string {
	return string(a)
}

func (a UserPassword) Equals(accountPassword *UserPassword) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(string(a)), []byte(accountPassword.Password()))
	if err != nil {
		return fmt.Errorf("%w: password does not match the account password", PasswordDoesNotMatch)
	}

	return err
}

// User represents the Object Domain of our business
type User struct {
	userID       UserID
	userName     UserName
	userPassword UserPassword
	role         *Role
}

// NewUser creates a new User instance with the provided parameters
func NewUser(
	userID,
	userName,
	userPassword string,
	role *Role,
) (*User, error) {
	userIDVO, err := UserID(userID).Validate()
	if err != nil {
		return nil, err
	}

	userNameVO, err := UserName(userName).Validate()
	if err != nil {
		return nil, err
	}

	userPasswordVO, err := UserPassword(userPassword).Validate()
	if err != nil {
		return nil, err
	}

	User := &User{
		userID:       *userIDVO,
		userName:     *userNameVO,
		userPassword: *userPasswordVO,
		role:         role,
	}

	return User, nil
}

// ID returns the ID of the User
func (a *User) ID() string {
	return a.userID.ID()
}

// Name returns the Name of the User
func (a *User) Name() string {
	return a.userName.Name()
}

// Password returns the Password of the User
func (a *User) Password() string {
	return a.userPassword.Encrypt().Password()
}

// Roles returns the Role of the User
func (a *User) Role() Role {
	return *a.role
}

// PasswordMatches check the match of passwords
func (a *User) PasswordMatches(password *UserPassword) error {
	return a.userPassword.Equals(password)
}
