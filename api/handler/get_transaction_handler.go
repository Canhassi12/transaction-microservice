package handler

import "github.com/gin-gonic/gin"

func Get_transaction(c *gin.Context) {
	c.String(200, "Hello, World!")
}
