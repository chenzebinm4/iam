package main

import (
	"github.com/chenzebinm4/iam/internal/apiserver/v1/user"
	"github.com/gin-gonic/gin"
	"log"

	_ "github.com/chenzebinm4/iam/api/swagger/docs"
)

func main() {
	r := gin.Default()
	r.POST("/users", user.Create)
	r.GET("/users/:name", user.Get)

	log.Fatal(r.Run(":80"))
}
