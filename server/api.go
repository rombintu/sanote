package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rombintu/sanote/models"
	"github.com/rombintu/sanote/store"
	"github.com/rombintu/sanote/tools"
)

type API struct {
	Router *gin.Engine
	Store  *store.Store
}

func NewAPI() *API {
	return &API{
		Router: gin.Default(),
		Store:  store.NewStore(),
	}
}

func (api *API) Start() error {
	api.ConfigureRouter()

	return http.ListenAndServe(
		tools.GetEnvOrDefault("PORT", ":8081"),
		api.Router,
	)
}

func (api *API) ConfigureRouter() {
	if tools.GetEnvOrDefault("DEV", "false") == "true" {
		api.Router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "*"
			},
			MaxAge: 12 * time.Hour,
		}),
		)
	}

	api.Router.GET("/", api.Index())

	api.Router.GET("/note/id", api.GetNoteById())   // ?_id=?
	api.Router.GET("/note", api.GetNotesByAuthor()) // ?author=?
	api.Router.POST("/note", api.CreateNote())
	api.Router.PUT("/note", api.UpdateNoteById())
	api.Router.DELETE("/note", api.DeleteNoteById())

	api.Router.GET("/user", api.GetUserByLogin()) // ?login=?
	api.Router.POST("/user", api.CreateUser())
	api.Router.PUT("/user", api.UpdateUserById())
	api.Router.DELETE("/user", api.DeleteUserById())
}

func (api *API) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.REST{
			Message: "Server is starting", Error: 0},
		)
	}
}
