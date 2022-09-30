package user

import (
	"context"
	"fmt"
	"finaltask/config"
	"finaltask/models"
	"log"
	"time"
)

const (
	table          = "user"
	layoutDateTime = "2006-01-02 15:04:05"
)


// Insert
func Insert(ctx context.Context, usr models.User) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id, username, email, password, createat, updateat) values(%v,'%v','%v','%v','%v','%v')", table,
		usr.ID,
		usr.Username,
		usr.Email,
		usr.Password,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Take for GetLogin
func Take(ctx context.Context) ([]models.User, error) {

	var users []models.User

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var user models.User
		

		if err = rowQuery.Scan(&user.ID,
			&user.Password); err != nil {
			return nil, err
		}


		users = append(users, user)
	}

	return users, nil
}