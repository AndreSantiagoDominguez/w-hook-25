package main

import (
	infrastructure "web-hook/pull_reques_webhook/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	infrastructure.RegisterRouter(r)

	r.Run()
}