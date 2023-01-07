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

// REQUEST MODELS

type UserByIdRequest struct {
	Id string `json:"id"`
}

type CreateUserRequest struct {
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

// REQUEST MODELS

func NewUsersHandler(service service.User) *UserHandler {
	return &UserHandler{userService: service}
}

// GET ALL USERS

func (h UserHandler) GetUsers(c *gin.Context) {
	user, err := h.userService.GetUsers(c)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, user)
}

// GET ALL USERS

// GET USER BY ID

func (h UserHandler) GetUserById(c *gin.Context) {
	var user *domain.User
	var err error
	var req []byte

	req, err = io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	var requestedId UserByIdRequest

	err = json.Unmarshal(req, &requestedId)
	if err != nil {
		panic(err)
	}

	id, err := uuid.Parse(requestedId.Id)
	if err != nil {
		fmt.Println(err)
	}

	user, err = h.userService.GetUserById(c, id)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, user)
}

// GET USER BY ID

// CREATE USER

func (h UserHandler) CreateUser(c *gin.Context) {
	var err error
	var req []byte

	req, err = io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	var userRequest CreateUserRequest

	err = json.Unmarshal(req, &userRequest)
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

// CREATE USER

// DELETE USER

func (h UserHandler) DeleteUser(c *gin.Context) {
	var err error
	var req []byte

	req, err = io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	var requestedId UserByIdRequest

	err = json.Unmarshal(req, &requestedId)
	if err != nil {
		panic(err)
	}

	id, err := uuid.Parse(requestedId.Id)
	if err != nil {
		fmt.Println(err)
	}

	err = h.userService.DeleteUser(c, id)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": string("User successfully deleted"),
	})
}

// DELETE USER
