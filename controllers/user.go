package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shanmuganthan/go-lang-mongo/models"
)

func CreateAdminUser(w http.ResponseWriter, r *http.Request) {

	body := r.Body
	var userModel models.UserModel

	json.NewDecoder(body).Decode(&userModel)

	fmt.Printf("BODY %+v", userModel)

	json.NewEncoder(w).Encode(map[string]string{"message": "user has been create successfully."})
}

func UpdateAdminUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "user has been create successfully."})
}

func DeleteAdminUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "user has been create successfully."})
}

func GetByIdAdminUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "user has been create successfully."})
}

func GetAllAdminUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "user has been create successfully."})
}
