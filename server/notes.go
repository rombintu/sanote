package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/sanote/models"
	"github.com/rombintu/sanote/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GET (GET)
func (api *API) GetNotesByAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		author := c.Query("author")
		notes, err := api.Store.GetNotesByAuthor(author)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}
		c.JSON(http.StatusOK, notes)
	}
}

// GET (GET) ID
func (api *API) GetNotesById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Query("_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}
		notes, err := api.Store.GetNoteById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}
		c.JSON(http.StatusOK, notes)
	}
}

// CREATE (POST)
func (api *API) CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var note models.Note
		if err := c.Bind(&note); err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}

		if err := api.Store.CreateNote(note); err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
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
func (api *API) UpdateNoteById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var note models.Note
		if err := c.Bind(&note); err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}

		if err := api.Store.UpdateNoteById(note); err != nil {
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
func (api *API) DeleteNoteById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Query("_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}
		if err := api.Store.DeleteNoteById(id); err != nil {
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
