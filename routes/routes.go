package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletarUmAluno)
	r.PATCH("/alunos/:id", controllers.AlterandoUmAluno)
	r.GET("/alunos/cpf/:cpf", controllers.AlterandoUmAluno)
	r.Run()
}
