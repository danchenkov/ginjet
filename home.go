package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeIndex(c *gin.Context) {
	c.String(http.StatusOK, "This is my home page")
}
