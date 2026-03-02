package handler


import(
	"encoding/json"
	"net/http"
	"backend-api-belajar/model"
	"backend-api-belajar/service"
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