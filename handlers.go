package main


import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
	"encoding/json"
)
func findUserIndexById(id int) int {
	for i, user := range users {
		if user.Id == id {
			return i
		}
	}
	return -1
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	idx := findUserIndexById(id)
	if idx == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}


	json.NewEncoder(w).Encode(users[idx])
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	idx := findUserIndexById(id)
	if idx == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}


	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser.Id = id
	users[idx] = updatedUser

	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	
	idx := findUserIndexById(id)
	if findUserIndexById(id) == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	userName := users[idx].Name

	users = append(users[:idx], users[idx+1:]...)
	json.NewEncoder(w).Encode(map[string]string{
					"message": fmt.Sprintf("User '%s' deleted", userName),
			})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(users) == 0 {
		newUser.Id = 1

	} else {
		newUser.Id = users[len(users)-1].Id + 1

	}
	users = append(users, newUser)

	json.NewEncoder(w).Encode(newUser)
}
