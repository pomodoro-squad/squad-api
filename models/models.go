package models

import (
	"time"
)

type User struct {
	SBasicModel

	ReferralId int64  `json:"referral_id"`
	NickName   string `json:"nick_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	AvatarId   string `json:"avatar_id"`
	Age        int    `json:"age"`

	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PasswordHash  string `json:"-"`
	PasswordSalt  string `json:"-"`
	Password      string `json:"password"`
	ResetPassword string `json:"reset_password"`

	Communities    []string  `json:"communities"`
	UserConnection []int64   `json:"user_connection"`
	Timer          Iteration `json:"timer"`
	Busy           bool      `json:"busy"`
}

type Iteration struct {
	LastTimeOn           time.Time `json:"last_time_on"`
	LastTimeOFf          time.Time `json:"last_time_off"`
	CompleteIterations   int64     `json:"complete_iterations"`
	IncompleteIterations int64     `json:"incomplete_iterations"`
}

type UserList struct {
	Users []User
}
