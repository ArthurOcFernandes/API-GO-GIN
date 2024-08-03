package main

import (
	"gin-api/models"
	"gin-api/routes"
)

func main() {

	models.Alunos = []models.Aluno{
		{Nome: "Arthur", CPF: "121313131", RG: "124141241241"},
		{Nome: "Gui", CPF: "1313131", RG: "4141241241"},
	}
	routes.HandleRequests()
}
