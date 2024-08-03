package routes

import (
	"gin-api/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controller.ExibeTodosAlunos)

	r.Run()
}
