package main

import (
	"encoding/json"
	"fmt"
	"gin-api/controller"
	"gin-api/database"
	"gin-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeGetAllAlunos(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupRotasDeTeste()

	r.GET("/alunos", controller.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	defer DeletaAlunoMock()

}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()

	r := SetupRotasDeTeste()

	r.GET("/alunos/cpf/:cpf", controller.BuscarAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	defer DeletaAlunoMock()
}

func TestListaAlunoPorId(t *testing.T) {

	database.ConectaComBancoDeDados()

	CriaAlunoMock()

	r := SetupRotasDeTeste()

	r.GET("/alunos/:id", controller.ExibeAlunoPorId)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/alunos/%d", ID), nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, "12345678901", alunoMock.CPF)

	defer DeletaAlunoMock()

}

func TestDelecaoAluno(t *testing.T) {

	database.ConectaComBancoDeDados()

	CriaAlunoMock()

	r := SetupRotasDeTeste()

	r.DELETE("/alunos/:id", controller.DeletaAluno)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/alunos/%d", ID), nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusNoContent, resposta.Code)

}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}

	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	database.DB.Delete(&models.Aluno{}, ID)
}
