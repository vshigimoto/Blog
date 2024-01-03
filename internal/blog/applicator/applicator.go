package applicator

import (
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
	mainDb, err := database.New(a.cfg.Database.Main)
	if err != nil {
		a.l.Errorf("fail create mainDb %v", err)
		return
	}
	defer func() {
		if err := mainDb.Close(); err != nil {
			a.l.Errorf("fail close mainDb %v", err)
			return
		}
		a.l.Info("mainDb closed")
	}()
	replicaDb, err := database.New(a.cfg.Database.Replica)
	if err != nil {
		a.l.Errorf("fail create replica Db %v", err)
	}
	defer func() {
		if err := mainDb.Close(); err != nil {
			a.l.Errorf("fail close mainDb %v", err)
			return
		}
		a.l.Info("replicaDb closed")
	}()

	// move to the handlers in server package
	// think about replication
	//server := server.New(gin.Default(), a.cfg, mainDb)
}
