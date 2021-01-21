package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	v1 "github.com/lucabecci/docker-golang/internal/server/v1"
)

//Server struct is for my server
type Server struct {
	server *http.Server
}

//New is for create the server
func New(port string) (*Server, error) {
	router := chi.NewRouter()

	router.Mount("/", v1.New())

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

//Close is for close the connection in the server
func (serv *Server) Close() error {
	return nil
}

//Start is for start the server
func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
