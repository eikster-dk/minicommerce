package http

import (
	"fmt"
	"net/http"

	"github.com/eikc/minicommerce"
	"github.com/julienschmidt/httprouter"
)

// Server is the http server for serving the minicommerce rest API
type Server struct {
	downloadableRepository minicommerce.DownloadableRepository
	storage                minicommerce.Storage
	router                 *httprouter.Router
}

// NewServer is the constructor for the Http Server
func NewServer(downloadableRepository minicommerce.DownloadableRepository,
	storage minicommerce.Storage) *Server {

	return &Server{
		downloadableRepository: downloadableRepository,
		storage:                storage,
		router:                 httprouter.New(),
	}
}

// Run starts the server with all the given params
func (s *Server) Run(port string) error {
	s.routes()
	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.router)
}