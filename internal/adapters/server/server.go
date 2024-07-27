package server

import (
	"fmt"
	"messenger_service/internal/routes"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct{}

type IServer interface {
	Run()
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.Use()
	routes.Routes(router)
	http.Handle("/", router)

	port := os.Getenv("PORT")
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{allowedOrigins}))(router)

	fmt.Println(fmt.Printf("SERVER LISTENNING ON PORT: %s", port))
	panic(http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler))
}
