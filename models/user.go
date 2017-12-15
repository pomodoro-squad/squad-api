package models

import (
	"time"
)

type User struct {
	SBasicModel
	Timer          Timer

	ReferralId
	NickName  
	FirstName 
	LastName  
	AvatarId  
	Age       
	Email     
	EmailVerified 
	Groups    
	Referral  
}

type Iteration struct {
	SBasicModel
	UserID
	Duration
	Completed bool
}

type Group struct {
	SBasicModel
	OwnerID
	Title
	Description
}
