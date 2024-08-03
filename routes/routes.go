package routes

import (
	"gin-api/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.POST("/alunos", controller.CriaNovoAluno)
	r.GET("/alunos/:id", controller.ExibeAlunoPorId)
	r.DELETE("alunos/:id", controller.DeletaAluno)
	r.Run()
}
