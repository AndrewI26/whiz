package server

import (
	"fmt"
	"net/http"

	logging "github.com/AndrewI26/whiz/logger"
	"github.com/AndrewI26/whiz/routing"
)

type Router interface {
	Handle(method string, path string)
}

type Server struct {
	port    int
	routers []routing.Router
	logger  *logging.Logger
}

func NewServer(logger *logging.Logger, port int) *Server {
	return &Server{logger: logger, port: port}
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
				res := handler(params)
				for key, val := range res.Headers {
					w.Header().Add(key, val)
				}
				w.WriteHeader(res.Status)
				_, err := w.Write([]byte(res.Data))
				if err != nil {
					s.logger.Error(fmt.Sprintf("writing response: %s", err))
					// Log error
				}
			} else {
				w.WriteHeader(404)
				msg := "404 Page Not Found"
				w.Write([]byte(msg))
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
