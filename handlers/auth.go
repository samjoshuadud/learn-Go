package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	
	"time"
	"os"

	"learnGo/db"
	"learnGo/models"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err !=nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)

	result, err := db.DB.Exec("INSERT INTO users (name, age, password) VALUES (?, ?, ?)", user.Name, user.Age, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	id, _ := result.LastInsertId()
	user.ID = int(id)
	user.Password = ""
	json.NewEncoder(w).Encode(user)

}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))


func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	err := db.DB.QueryRow("SELECT id, name, password FROM users WHERE name = ?", req.Name).Scan(&user.ID, &user.Name, &user.Password)

	if err != nil {
		http.Error(w, "User not Found", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		Subject:   fmt.Sprint(user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})

}
