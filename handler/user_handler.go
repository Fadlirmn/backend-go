package handler

import (
	"backend-api-belajar/model"
	"backend-api-belajar/service"
	"encoding/json"
	"net/http"
	
)

type UserHandler struct{
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler  {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetUsers()
	json.NewEncoder(w).Encode(users)
}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	h.service.CreateUser(user)
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
    if id == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("ID is required"))
        return
    }
    
    // 2. Decode Body JSON dari Postman
    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid JSON body"))
        return
    }

    // 3. Panggil Service
    err = h.service.UpdateUser(id, user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Failed to update user in database"))
        return
    }
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request)  {
	id := r.URL.Query().Get("id")
	

	err := h.service.DeleteUser(id)
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User Deleted Successfully"))
}