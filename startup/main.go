package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"xiaoluokitchen/config"
	"xiaoluokitchen/controller"
)

func main() {
	fmt.Println("---server init start---")

	if err := config.InitComponent(); err != nil {
		log.Fatal("component init fail", err)
		return
	}

	server := gin.Default()
	server = controller.SetRoutes(server)
	if err := server.Run("0.0.0.0:8090"); err != nil {
		log.Fatal("server init fail", err)
		return
	}

	fmt.Println("---server init end---")
}
