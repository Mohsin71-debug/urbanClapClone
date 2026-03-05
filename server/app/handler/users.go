package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/rsj-rishabh/urbanClapClone/server/app/model"
)

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("CreateUser API called")

	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// check if username already exists
	existingUser := model.User{}
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		respondError(w, http.StatusConflict, "User already exists")
		return
	}

	// create new user
	if err := db.Create(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	userFromReqBody := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userFromReqBody); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	username := userFromReqBody.Username
	password := userFromReqBody.Password

	fmt.Printf("username : %s\n", username)
	fmt.Printf("password : %s\n", password)

	user := getUserOr404(db, username, password, w, r)
	if user == nil {
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func getUserOr404(db *gorm.DB, username string, password string, w http.ResponseWriter, r *http.Request) *model.User {

	user := model.User{}

	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		respondError(w, http.StatusNotFound, "Invalid username or password")
		return nil
	}

	return &user
}

func GetUserDetails(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	user := model.User{}

	serviceId := r.URL.Query()["userId"]

	i, err := strconv.Atoi(serviceId[0])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid userId")
		return
	}

	if err := db.Where("id = ?", i).First(&user).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, user)
}
