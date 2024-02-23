package main

import (
	"github.com/benebobaa/go-apigateway/pkg/auth"
	"github.com/benebobaa/go-apigateway/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	_ = *auth.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
