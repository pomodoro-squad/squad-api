package models

type User struct {

	FirstName string `json:"first_name,omitempty"`
	LastName string	`json:"last_name,omitempty"`
	TimeOn bool	`json:"timer_on,omitempty"`
}

type UserList struct {
	Users []User
}
