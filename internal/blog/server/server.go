package server

import (
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

func (s *Server) Run() error {
	
}
