package routes

import (
	"gin-api/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.POST("/alunos", controller.CriaNovoAluno)
	r.GET("/alunos/:id", controller.ExibeAlunoPorId)
	r.DELETE("alunos/:id", controller.DeletaAluno)
	r.PUT("/alunos/:id", controller.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controller.BuscarAlunoPorCPF)
	r.GET("/", controller.ExibeIndex)
	r.NoRoute(controller.RotaNaoEncontrada)
	r.Run()
}
