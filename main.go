package main

import ( //modulos
	"fmt"

	"github.com/nataliakdiniz/api-rest-go/database"
	"github.com/nataliakdiniz/api-rest-go/models"
	"github.com/nataliakdiniz/api-rest-go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor!")
	routes.HandleResquest()
}
