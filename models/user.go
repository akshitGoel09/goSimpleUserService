package models

// User represent a user strung containing userId, firstName and lastName
type User struct{
	ID int
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)