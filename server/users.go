package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/sanote/models"
	"github.com/rombintu/sanote/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GET (GET)
func (api *API) GetUserByLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.Query("login")
		user, err := api.Store.GetUsersByLogin(login)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// CREATE (POST)
func (api *API) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}

		if err := api.Store.CreateUser(user); err != nil {
			c.JSON(http.StatusBadRequest, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}

		c.JSON(http.StatusCreated, models.REST{
			Message: tools.Created, Error: 0},
		)
	}
}

// UPDATE (PUT)
func (api *API) UpdateUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}

		if err := api.Store.UpdateUserById(user); err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}

		c.JSON(http.StatusCreated, models.REST{
			Message: tools.Updated, Error: 0},
		)
	}
}

// DELETE (DELETE)
func (api *API) DeleteUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Query("_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}
		if err := api.Store.DeleteUserById(id); err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}

		c.JSON(http.StatusAccepted, models.REST{
			Message: tools.Deleted, Error: 0},
		)
	}
}
