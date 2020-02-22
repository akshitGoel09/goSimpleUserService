package models

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

	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
