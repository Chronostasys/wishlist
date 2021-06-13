package main

import (
	"log"

	"github.com/Chronostasys/wishlist/controllers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	conn, err := config.String("sqlconn")
	if err != nil {
		log.Panic(err)
	}
	orm.RegisterDataBase("default", "mysql", conn)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/adduser", controllers.NewUserCtr().Register)
	orm.RunSyncdb("default", false, true)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
