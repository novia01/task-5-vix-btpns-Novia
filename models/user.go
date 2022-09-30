package models

import (
	"time"
)

type (
	// User
	User struct {
		ID        	int       	`json:"id"`
		Username	string    	`json:"username"`
		Email  		string    	`json:"email"`
		Password 	string		`json:"password"`
		CreateAt 	time.Time 	`json:"createat"`
		UpdateAt 	time.Time 	`json:"updateat"`
	}
)