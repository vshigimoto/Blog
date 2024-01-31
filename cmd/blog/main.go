package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/vshigimoto/LinkedIn-clone/internal/user/config"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)

	l := logger.Sugar()
	l = l.With(zap.String("app", "Blog"))

	var cfg config.Config
	err := cleanenv.ReadConfig("config/user/config.yaml", &cfg)
	if err != nil {
		l.Fatalf("fail read config err: %v", err)
		return
	}
}
