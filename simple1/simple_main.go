/**
 * Auth :   liubo
 * Date :   2018/10/22 14:25
 * Comment:
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/fvbock/endless"

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	endless.ListenAndServe(":4242", router)
}
func hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "simple world")
}

