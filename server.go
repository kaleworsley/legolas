package legolas

import (
	"net/http"

	"github.com/unrolled/render"
)

// Server contains all of the http handlers.
type Server struct {
	Render *render.Render
}

// IndexRoute is the index handler.
func (s *Server) IndexRoute(w http.ResponseWriter, r *http.Request) {
	s.Render.HTML(w, http.StatusOK, "index", nil)
}
