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

func DeletaAluno(c *gin.Context) {
	id := c.Params.ByName("id")

	var aluno models.Aluno

	database.DB.First(&aluno, id)

	database.DB.Delete(&aluno)

	c.JSON(http.StatusNoContent, gin.H{
		"Status:": "Deletado com sucesso",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)

	c.JSON(http.StatusCreated, aluno)

}

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno

	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	err := c.ShouldBindJSON(&aluno)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error:": "Aluno não existe no banco de dados",
		})
		return
	}

	database.DB.Model(&aluno).Updates(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscarAlunoPorCPF(c *gin.Context) {

	var aluno models.Aluno
	cpf := c.Param("cpf")

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Error:": "cpf inexistente na base de dados",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func ExibeIndex(c *gin.Context) {

	var alunos []models.Aluno

	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
