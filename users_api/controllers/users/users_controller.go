package users

import (
	"net/http"
	"strconv"

	"github.com/bookstore_app/users_api/domain/users"
	"github.com/bookstore_app/users_api/services"
	"github.com/bookstore_app/users_api/utils/errors"

	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		// param not valid
		err := errors.HandleBadRequestErr("invalid User id in URL.")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
	// c.String(http.StatusNotImplemented, "TODO MEE!")
}
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.HandleBadRequestErr("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// sends user to save in DB
	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		// Handle Error -> controller just returns this error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func DeleteUser(c *gin.Context) {
	// find user id
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	// if param is bad.. handle
	if userErr != nil {
		err := errors.HandleBadRequestErr("invalid User id in URL.")
		c.JSON(err.Status, err)
		return
	}
	// if id doesn't exist.. respond
	if deleteErr := services.DestroyUser(userId); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
