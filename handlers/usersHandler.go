package handlers

import (
	"encoding/json"
	modelapp "goFinalTask/modelAPP"
	modeldb "goFinalTask/modelDB"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Method Not Found"))
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []modelapp.Users
	// Get users from database
	result := modeldb.DB.Find(&users)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Get Users Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user modelapp.Users
	// Get user from request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Add User Error"))
		return
	}
	// Add user to database
	result := modeldb.DB.Create(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Add User Error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Delete User Error"))
		return
	}
	result := modeldb.DB.Find(&modelapp.Users{}, id)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	} else if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "User with ID " + id + " Not Found"})
		return
	}

	// Delete user from database
	result = modeldb.DB.Delete(&modelapp.Users{}, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// Send response to client

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Delete User with ID " + id + " Successfully"})

}
