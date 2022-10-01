package main

import (
	"context"
	"encoding/json"
	"finaltask/user"
	"finaltask/models"
	"finaltask/utils"
	"log"
	"fmt"
	"net/http"
	"strconv"
)


func main() {
	http.HandleFunc("/user/create", PostResgister)
	http.HandleFunc("/user", GetUser)
	http.HandleFunc("/user/update", UpdateUser)
	http.HandleFunc("/user/delete", DeleteUser)
	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}


//PostUser
func PostResgister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var usr models.User

		if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := user.Insert(ctx, usr); err != nil {
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

// GetUser
func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		users, err := user.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, users, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

// UpdateUser
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var usr models.User

		if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		fmt.Println(usr)

		if err := user.Update(ctx, usr); err != nil {
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

// DeleteUser
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var usr models.User

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		usr.ID, _ = strconv.Atoi(id)

		if err := user.Delete(ctx, usr); err != nil {

			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
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
