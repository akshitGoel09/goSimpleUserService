package model

import (
	"errors"
	"fmt"
)

// User represent a user strung containing userId, firstName and lastName
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers functions returns the slice of pointers of user objects
func GetUsers() []*User {
	return users
}

//AddUser function add the user to our users slice
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include an id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//GetUserByID returns the user corrosponding to the passed ID
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

//UpdateUser updates the user
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.id)
}

//RemoveUserByID removes user of given id
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	eturn User{}, fmt.Errorf("User with ID '%v' not found", id)
}
