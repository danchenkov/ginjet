package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1", "ginjet.test"})

	router.GET("/", homeIndex)

	// router.Run("ginjet.test:8003")
	router.RunTLS("ginjet.test:443", "./ssl/ginjet.crt", "./ssl/ginjet.key")
}
