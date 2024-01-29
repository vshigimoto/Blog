package applicator

import (
	"github.com/redis/go-redis"
	"github.com/vshigimoto/Blog/internal/blog/config"
	"github.com/vshigimoto/Blog/internal/blog/database"
	"go.uber.org/zap"
)

type Applicator struct {
	l   *zap.SugaredLogger
	cfg *config.Config
}

func New(l *zap.SugaredLogger, cfg *config.Config) *Applicator {
	return &Applicator{
		l:   l,
		cfg: cfg,
	}
}

func (a *Applicator) Run() {
	mainDB, err := database.New(a.cfg.Database.Main)
	if err != nil {
		a.l.Errorf("fail create mainDB %v", err)
		return
	}

	defer func() {
		if err := mainDB.Close(); err != nil {
			a.l.Errorf("fail close mainDb %v", err)
			return
		}
		a.l.Info("mainDb closed")
	}()

	replicaDB, err := database.New(a.cfg.Database.Replica)
	if err != nil {
		a.l.Errorf("fail create replica Db %v", err)
	}

	defer func() {
		if err := replicaDB.Close(); err != nil {
			a.l.Errorf("fail close replicaDB %v", err)
			return
		}
		a.l.Info("replicaDb closed")
	}()

	rdb := redis.NewClient(&redis.Options{
		Addr:     a.cfg.Database.Redis.Addr,
		Password: a.cfg.Database.Redis.Password, // no password set
		DB:       a.cfg.Database.Redis.DB,       // use default DB
	})
}
