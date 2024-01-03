package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis"
	"github.com/vshigimoto/Blog/internal/blog/config"
	"go.uber.org/zap"
)

type Server struct {
	eng *gin.Engine
	cfg *config.Config
	db  *sqlx.DB
	rdc *redis.Client
	log *zap.SugaredLogger
}

func New(eng *gin.Engine, cfg *config.Config, db *sqlx.DB, rdc *redis.Client, log *zap.SugaredLogger) *Server {
	return &Server{
		eng: eng,
		cfg: cfg,
		db:  db,
		rdc: rdc,
		log: log,
	}
}

func (s *Server) Run() {
	// add context with cancel for gShutdown
	go func() {
		s.log.Infof("Server on port %d start", s.cfg.HttpServer.Port)
		addr := fmt.Sprintf(":%d", s.cfg.HttpServer.Port)
		err := s.eng.Run(addr)
		if err != nil {
			s.log.Infof("Fail run user server %v", err)
			return
		}
	}()
	go func() {
		s.log.Infof("Server on port %d start", s.cfg.HttpServer.AdminPort)
		addr := fmt.Sprintf(":%d", s.cfg.HttpServer.AdminPort)
		err := s.eng.Run(addr)
		if err != nil {
			s.log.Infof("Fail run admin server %v", err)
			return
		}
	}()
	// add debug server with http.ListenAndServe and pprof. Add pprof in config
	// add graceful shutdown
	// init handlers
}
