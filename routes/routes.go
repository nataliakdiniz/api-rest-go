package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nataliakdiniz/api-rest-go/controllers"
	"github.com/nataliakdiniz/api-rest-go/middleware"
)

func HandleResquest() {
	r := mux.NewRouter() //instancia do gorilla-mux
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home) //Lidar com as requisicoes - chega o / quem atende é o controler-home
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")         //criar recurso com postman
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")    //deletar recurso com postman
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")           //editar recurso com postman
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))) //servidor
}
