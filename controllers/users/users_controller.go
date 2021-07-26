package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mahdipardat/bookstore_user-api/domain/users"
	"github.com/mahdipardat/bookstore_user-api/services"
	"github.com/mahdipardat/bookstore_user-api/utils/errors"
)


func CreateUser(c *gin.Context) {
	var user users.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		resErr := errors.BadRequestError("invalid json body")
		c.JSON(resErr.Status, resErr)
		return
	}

	result, err := services.CreateUser(user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		rErr := errors.BadRequestError("invalid user_id")
		c.JSON(rErr.Status, rErr)
		return
	}

	result, userErr := services.GetUser(userId)

	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	c.JSON(http.StatusOK, result)
}