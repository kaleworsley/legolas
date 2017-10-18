package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jamiealquiza/envy"
	"github.com/kaleworsley/legolas"
	"github.com/unrolled/render"
)

var config struct {
	Port      string
	Host      string
	Templates string
}

func main() {
	flag.StringVar(&(config.Port), "port", "8080", "http port")
	flag.StringVar(&(config.Host), "host", "127.0.0.1", "http host")
	flag.StringVar(&(config.Templates), "templates", "./templates", "path to templates directory")
	envy.Parse("LEGOLAS")
	flag.Parse()

	logger := log.New(os.Stderr, "[LEGOLAS] ", log.LstdFlags)

	render := render.New(render.Options{
		Directory:     config.Templates,
		IsDevelopment: true,
		Layout:        "application/application",
	})

	server := &legolas.Server{
		Render: render,
		Logger: logger,
	}
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Get("/", server.IndexRoute)

	svr := http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: router,
	}

	logger.Printf("Starting server on %v:%v...\n", config.Host, config.Port)
	logger.Fatalln(svr.ListenAndServe())
}
