package main

import (
	"github.com/gin-gonic/gin"
	srv "my_project/project-common"
)

func main() {
	r := gin.Default()
	srv.Run(r, "project-user", ":8082")
}
