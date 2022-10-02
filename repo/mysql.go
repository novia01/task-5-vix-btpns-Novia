package user

import (
	"context"
	"fmt"
	"finaltask/config"
	"finaltask/models"
	"log"
	"time"
	"database/sql"
	"errors"
)

const (
	table          	= "user"
	table2			= "foto"
	layoutDateTime 	= "2006-01-02 15:04:05"
)


//USER ENDPOINT

// InsertUser
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


//GetUser
func GetAll(ctx context.Context) ([]models.User, error) {

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
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  ubah format dari string ke datetime
		user.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		user.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users, nil
}

// UpdateUser
func Update(ctx context.Context, usr models.User) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set username = '%s', email ='%s', password = '%s', updated_at = '%v' where id = '%d'",
		table,
		usr.Username,
		usr.Email,
		usr.Password,
		time.Now().Format(layoutDateTime),
		usr.ID,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}


// DeleteUSer
func Delete(ctx context.Context, usr models.User) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, usr.ID)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	return nil
}



//FOTO ENDPOINT

// InsertFoto
func InsertFoto(ctx context.Context, ft models.Foto) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id, title, caption, url, userid) values(%v,'%v','%v','%v','%v')", table2,
		ft.ID,
		ft.Title,
		ft.Caption,
		ft.URL,
		ft.User_ID)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}


//GetFoto
func GetFoto(ctx context.Context) ([]models.Foto, error) {

	var fotos []models.Foto

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table2)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var foto models.Foto

		if err = rowQuery.Scan(&foto.ID,
			&foto.ID,
			&foto.Title,
			&foto.Caption,
			&foto.URL,
			&foto.User_ID); err != nil {
			return nil, err
		}
	}

	return fotos, nil
}


// UpdateFoto
func UpdateFoto(ctx context.Context, ft models.Foto) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title = '%s', caption ='%s', url = '%s', userid = '%d' where id = '%d'",
		table2,
		ft.Title,
		ft.Caption,
		ft.URL,
		ft.User_ID,
		ft.ID,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}


// DeleteFoto
func DeleteFoto(ctx context.Context, ft models.Foto) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table2, ft.ID)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	return nil
}