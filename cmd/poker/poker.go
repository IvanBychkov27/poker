package main

import (
	"github.com/IvanBychkov27/internal/config"
	"go.uber.org/zap"
	"log"
	"os"
)

var (
	version = "undefined"
)

func main() {
	cfg := config.New()
	err := cfg.Load()
	if err != nil {
		log.Printf("error load config, %w", err)
		os.Exit(1)
	}

	var logger *zap.Logger
	//if cfg.Debug {
	logger, err = zap.NewDevelopment()
	//} else {
	//	logger, err = zap.NewProduction()
	//}
	if err != nil {
		log.Printf("error init zap logger, %v", err)
		os.Exit(1)
	}

	logger.Info("tsw", zap.String("version", version))
	logger.Info("config loaded", zap.Any("config", cfg))

	//err = run(cfg, logger)
	//if err != nil {
	//	logger.Error("error run application", zap.Error(err))
	//	os.Exit(1)
	//}

	logger.Info("done")
}
