package models

type User struct {
	FirstName string
	LastName string
	TimeOn bool
}

type UserList struct {
	Users []User
}
