package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oriastanjung/01-simple-restAPI/api/users/router"
	"github.com/oriastanjung/01-simple-restAPI/config"
	"github.com/oriastanjung/01-simple-restAPI/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	app := gin.Default()
	db.ConnectDB()
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})

	router.UsersGroup(app)

	fmt.Println("PORT >> ", cfg.PORT)
	app.Run(":" + cfg.PORT)
}
