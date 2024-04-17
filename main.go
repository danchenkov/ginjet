package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	httpsRouter := gin.Default()
	httpsRouter.SetTrustedProxies([]string{"127.0.0.1", "ginjet.test"})
	httpsRouter.GET("/", homeIndex)
	httpsRouter.HEAD("/", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	// httpRouter := gin.Default()
	// httpRouter.GET("/*p", func(c *gin.Context) {
	// 	c.Redirect(302, "https://ginjet.test/"+c.Param("p"))
	// })

	// go httpsRouter.RunTLS("ginjet.test:443", "./ssl/ginjet.crt", "./ssl/ginjet.key")
	// httpRouter.Run("ginjet.test:80")

	httpsRouter.RunTLS("ginjet.test:8443", "./ssl/ginjet.crt", "./ssl/ginjet.key")
}
