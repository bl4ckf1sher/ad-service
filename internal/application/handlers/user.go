package handlers

import (
	"encoding/json"
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

// REQUEST MODELS

type UserRequest struct {
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

// REQUEST MODELS

// GET ALL USERS

func (h UserHandler) GetUsers(c *gin.Context) {
	user, err := h.userService.GetUsers(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
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
		if json.Valid(req) {
			c.AbortWithError(http.StatusUnprocessableEntity, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		panic(err)
	}

	user, err = h.userService.GetUserById(c, id)
	if err != nil {
		//TODO:
		// If id is valid, to check if user exists and throw 404 if so,
		// instead of just throwing 500
		c.AbortWithError(http.StatusInternalServerError, err)
		return
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
		//TODO:
		//Throw 500 if something else
		if json.Valid(req) {
			c.AbortWithError(http.StatusUnprocessableEntity, err)
			return
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	var userRequest UserRequest

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
		if json.Valid(req) {
			c.AbortWithError(http.StatusUnprocessableEntity, err)
			return
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = h.userService.DeleteUser(c, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User successfully deleted",
	})
}

// DELETE USER

// UPDATE USER

func (h UserHandler) UpdateUser(c *gin.Context) {
	var err error
	var req []byte

	req, err = io.ReadAll(c.Request.Body)
	if err != nil {
		if json.Valid(req) {
			c.AbortWithError(http.StatusUnprocessableEntity, err)
			return
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	var userRequest UserRequest

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(req, &userRequest)
	if err != nil {
		panic(err)
	}

	var user domain.User

	user.ID = id
	user.Role = userRequest.Role
	user.Email = userRequest.Email
	user.Password = userRequest.Password
	user.Name = userRequest.Name
	user.Surname = userRequest.Surname

	err = h.userService.UpdateUser(c, user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User successfully updated",
	})
}

// UPDATE USER
