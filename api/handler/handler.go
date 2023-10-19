package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

    r.GET("/transaction", Get_transaction)
    r.POST("/transaction", Create_transaction)

    return r
}
