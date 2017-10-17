package legolas

import (
	"fmt"
	"net/http"
)

// Server contains all of the http handlers.
type Server struct{}

// IndexRoute is the index handler.
func (s *Server) IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
