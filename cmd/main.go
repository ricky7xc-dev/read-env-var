package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ricky7xc-dev/read-env-var/pkg/books"
	"github.com/ricky7xc-dev/read-env-var/pkg/common/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.Run(port)
}
