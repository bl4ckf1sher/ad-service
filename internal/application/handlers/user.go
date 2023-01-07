package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bl4ckf1sher/ad-service/internal/domain"
	"github.com/bl4ckf1sher/ad-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type UserHandler struct {
	userService service.User
}

func NewUsersHandler(service service.User) *UserHandler {
	return &UserHandler{userService: service}
}

// Get all users

func (h UserHandler) GetUsers(c *gin.Context) {
	user, err := h.userService.GetUsers(c)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, user)
}

// Get user by id

func (h UserHandler) GetUserById(c *gin.Context) {
	var user *domain.User
	id, _ := uuid.Parse("d7c8c289-9e8a-4e1d-9ef3-8d3b17f6437c")
	user, err := h.userService.GetUserById(c, id)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, user)
}

// Create user

type UserRequest struct {
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func (h UserHandler) CreateUser(c *gin.Context) {
	var err error
	var res []byte

	res, err = io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	var userRequest UserRequest

	err = json.Unmarshal(res, &userRequest)
	if err != nil {
		panic(err)
	}

	var user domain.User

	user.Role = userRequest.Role
	user.Email = userRequest.Email
	user.Password = userRequest.Password
	user.Name = userRequest.Name
	user.Surname = userRequest.Surname

	err = h.userService.CreateUser(c, user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, userRequest)
}
