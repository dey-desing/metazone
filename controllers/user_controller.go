package controllers

import (
	"encoding/json"
	"net/http"
	"metazone/services"
	"metazone/models"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/users", UsersHandler).Methods("GET", "POST")
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		users := services.GetUsers()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)
		services.CreateUser(user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}