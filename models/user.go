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
		CreatedAt 	time.Time 	`json:"createdat"`
		UpdatedAt 	time.Time 	`json:"updatedat"`
	}
)