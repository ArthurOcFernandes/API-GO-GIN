package controller

import (
	"gin-api/database"
	"gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {

	var alunos []models.Aluno

	database.DB.Find(&alunos)

	c.JSON(http.StatusOK, alunos)
}

func ExibeAlunoPorId(c *gin.Context) {

	id := c.Params.ByName("id")

	var aluno models.Aluno

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error:": "Aluno not found",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)

}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
	}

	database.DB.Create(&aluno)

	c.JSON(http.StatusCreated, aluno)

}
