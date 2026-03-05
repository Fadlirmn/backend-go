package handler

import (
	"backend-api-belajar/model"
	"backend-api-belajar/service"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

type UserHandler struct{
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler  {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.service.GetUsers()
	c.JSON(http.StatusOK, users)
}
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid json: "+err.Error()})
		return 
	}

	h.service.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{"success":"User has been created"})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Query("id")
    if id == "" {
         c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid ID"})
        return
    }
    
    // 2. Decode Body JSON dari Postman
    var user model.User
    err := c.ShouldBindJSON(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error":"Invalid Body JSON "})
        return
    }

    // 3. Panggil Service
    err = h.service.UpdateUser(id, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
}

func (h *UserHandler) DeleteUser(c *gin.Context)  {
	id := c.Query("id")
	

	err := h.service.DeleteUser(id)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Id Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success":"users han been deleted"})
}