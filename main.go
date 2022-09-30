package main

import (
	"context"
	"encoding/json"
	"finaltask/user"
	"finaltask/models"
	"finaltask/utils"
	"log"
	"net/http"
	// "finaltask/config"
	"fmt"
)
func main() {
	http.HandleFunc("/user/create", PostResgister)
	http.HandleFunc("/user", GetLogin)
	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}


func PostResgister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mhs models.User

		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := user.Insert(ctx, mhs); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// GetLogin
func GetLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		users, err := user.Take(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, users, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}



//UNTUK TES KONEKSI DATABASE
// func main() {

// 	db, e := config.MySQL()

// 	if e != nil {
// 		log.Fatal(e)
// 	}

// 	eb := db.Ping()
// 	if eb != nil {
// 		panic(eb.Error()) 
// 	}

// 	fmt.Println("Success")

// 	err := http.ListenAndServe(":7000", nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
