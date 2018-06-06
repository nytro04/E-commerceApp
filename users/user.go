package users

import (
	"golang.org/x/crypto/bcrypt"
)

// user properties
type User struct {
	ID      int64
	Name    string
	Hash    []byte
	Address string
}

// user db interface
type UserDB interface {
	CreateUser(*User) (int64, error)
	GetUserByID(id int64) (*User, error)
	GetUserByName(name string) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(user *User) error
	RemoveUser(id int64) error
}

//create user func
func CreateUser(db UserDB, name, password, address string) (*User, error) {
	//Generate the password from the hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	up := User{
		Name:    name,
		Hash:    hash,
		Address: address,
	}

	up.ID, err = db.CreateUser(&up)
	if err != nil {
		return nil, err
	}

	return &up, nil
}

func Login(db UserDB, user, password string) (*User, error) {

	up, err := db.GetUserByName(user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(up.Hash, []byte(password))
	if err != nil {
		return nil, err
	}

	return up, err
}
