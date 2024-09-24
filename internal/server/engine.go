package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/voice-line/internal/repositories"
	"github.com/mohidex/voice-line/pkg/auth"
)

type Opt struct {
	Port         string
	Environment  string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Server struct {
	Opt    *Opt
	Engine *gin.Engine
	Routes *Routes
}

func NewServer(opt *Opt, userRepo repositories.UserRepository, auth auth.Authenticator) *Server {

	gin.SetMode(gin.ReleaseMode)
	if opt.Environment == "development" {
		gin.SetMode(gin.DebugMode)
	}

	// Create the Gin engine
	engine := gin.New()

	// Add middleware
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// Initialize the routes
	routes := &Routes{
		UserRepo: userRepo,
		Auth:     auth,
	}
	routes.Setup(engine)

	return &Server{
		Opt:    opt,
		Engine: engine,
		Routes: routes,
	}
}

func (s *Server) Start() error {
	srv := &http.Server{
		Addr:         ":" + s.Opt.Port,
		Handler:      s.Engine,
		ReadTimeout:  s.Opt.ReadTimeout,
		WriteTimeout: s.Opt.WriteTimeout,
	}

	log.Printf("Starting server on port %s in %s mode", s.Opt.Port, s.Opt.Environment)
	return srv.ListenAndServe()
}

func (s *Server) Shutdown() {

	log.Println("Shutting down server...")
}
