package models

import (
	"time"
)

type User struct {
	SBasicModel
	Timer          Timer

	ReferralId     int64    `json:"referral_id"`
	NickName       string   `json:"nick_name"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	AvatarId       string   `json:"avatar_id"`
	Age            int      `json:"age"`
	Email          string   `json:"email"`
	EmailVerified  bool     `json:"email_verified"`
	Groups    []string `json:"communities"`
	Referral       *int64   `json:"referral"`
}

type Iteration struct {
	SBasicModel
	UserID
	Duration
	Completed bool
}


