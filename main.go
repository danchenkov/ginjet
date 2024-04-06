package main

import (
	// jet "github.com/CloudyKit/jet/v6"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	// engine := jet.NewSet(
	// 	jet.NewOSFileSystemLoader("./views"),
	// )

	// r.HTMLRender = engine

	router.GET("/", homeIndex)

	// router.RunTLS(":8443", "./ginjet.crt", "./ginjet.key")
	router.Run(":8082")
}
