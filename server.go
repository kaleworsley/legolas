package legolas

//go:generate yarn run webpack
//go:generate go-bindata -pkg assets -ignore assets/assets\.go -o assets/assets.go public/...
//go:generate go-bindata -pkg templates -ignore templates/templates\.go -o templates/templates.go templates/...

import (
	"io"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

// Server contains all of the http handlers.
type Server struct {
	Render *render.Render
	Logger *log.Logger
}

func (s *Server) renderHTML(w io.Writer, status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) {
	err := s.Render.HTML(w, status, name, binding, htmlOpt...)
	if err != nil {
		s.Logger.Println("HTML Render Error:", err.Error())
	}
}

// IndexRoute is the index handler.
func (s *Server) IndexRoute(w http.ResponseWriter, r *http.Request) {
	s.Render.HTML(w, http.StatusOK, "index", nil)
}
