package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sampado/bookstore_oauth-api/src/domain/access_token"
	"github.com/sampado/bookstore_oauth-api/src/http"
	"github.com/sampado/bookstore_oauth-api/src/repository/db"
	"github.com/sampado/bookstore_oauth-api/src/repository/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	restRepo := rest.NewRepository()
	atService := access_token.NewService(dbRepository, restRepo)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run("127.0.0.1:8080")
}
