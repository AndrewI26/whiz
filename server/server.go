package server

import (
	"fmt"
	"net/http"

	"github.com/AndrewI26/whiz/routing"
)

type Router interface {
	Handle(method string, path string)
}

type Server struct {
	port    int
	routers []routing.Router
}

func NewServer(port int) *Server {
	return &Server{port: port}
}

func (s *Server) Serve() {
	fmt.Printf("Whiz server is running on http://localhost:%d\n", s.port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method

		for _, router := range s.routers {
			params, handler := router.FindRoute(method, path)

			// If we find the route
			if handler != nil {
				handler(params)
			}
		}
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
	if err != nil {
		panic(err)
	}

}

func (s *Server) AddRouter(router *routing.Router) {
	s.routers = append(s.routers, *router)
}
