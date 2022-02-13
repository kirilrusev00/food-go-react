/*
	Package server contains functions for starting and running the server.
*/
package server

import (
	"net/http"

	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/kirilrusev00/food-go-react/pkg/database"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*
	Server contains configurations for the server.
*/
type Server struct {
	config config.Config
	dbConn *database.DbConn
	router *mux.Router
}

/*
	NewServer creates a new server.
*/
func NewServer(config config.Config, dbConn *database.DbConn) (*Server, error) {
	server := &Server{
		config: config,
		dbConn: dbConn,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/image", server.createImage).Methods("POST")

	server.router = router
}

/*
	Start starts the server.
*/
func (server *Server) Start() {
	handler := server.corsHandler()
	http.ListenAndServe(server.config.Server.Address, handler)
}

func (server *Server) corsHandler() (handler http.Handler) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{server.config.Server.ClientAddress},
		AllowCredentials: true,
	})

	handler = c.Handler(server.router)
	return
}
