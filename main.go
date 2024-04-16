package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	httpsRouter := gin.Default()
	httpsRouter.SetTrustedProxies([]string{"127.0.0.1", "ginjet.test"})
	httpsRouter.GET("/", homeIndex)

	httpRouter := gin.Default()
	httpRouter.GET("/*p", func(c *gin.Context) {
		c.Redirect(302, "https://ginjet.test/"+c.Param("p"))
	})

	go httpsRouter.RunTLS("ginjet.test:443", "./ssl/ginjet.crt", "./ssl/ginjet.key")
	httpRouter.Run("ginjet.test:80")
}
