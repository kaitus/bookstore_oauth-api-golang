package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitus/bookstore_oauth-api-golang/src/domain/access_token"
	"github.com/kaitus/bookstore_oauth-api-golang/src/http"
	"github.com/kaitus/bookstore_oauth-api-golang/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")

}
