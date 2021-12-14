package main

import (
	"fmt"
	"github.com/IvanBychkov27/poker/internal/poker"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"go.uber.org/zap"
	"os"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("error init zap logger, %v \n", err)
		return
	}

	p := poker.NewPoker(logger)

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT must be set")
		return
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.GET("/", p.MainHandler)
	router.Run(":" + port)
}

//=====================================================================
//
//var (
//	version = "undefined"
//)
//
//func main() {
//	cfg := config.New()
//	err := cfg.Load()
//	if err != nil {
//		log.Printf("error load config, %w", err)
//		os.Exit(1)
//	}
//
//	var logger *zap.Logger
//	logger, err = zap.NewDevelopment()
//	if err != nil {
//		log.Printf("error init zap logger, %v", err)
//		os.Exit(1)
//	}
//
//	logger.Info("poker", zap.String("version", version))
//	logger.Info("config loaded", zap.Any("config", cfg))
//
//	err = run(cfg, logger)
//	if err != nil {
//		logger.Error("error run application", zap.Error(err))
//		os.Exit(1)
//	}
//
//	logger.Info("done")
//}
//
//func run(cfg *config.Config, logger *zap.Logger) error {
//	ln, err := net.Listen("tcp", cfg.MainAddress)
//	if err != nil {
//		return fmt.Errorf("error create main listener, %w", err)
//	}
//	defer ln.Close()
//
//	lnControl, err := net.Listen("tcp", cfg.ControlAddress)
//	if err != nil {
//		return fmt.Errorf("error create main listener, %w", err)
//	}
//	defer ln.Close()
//
//	ctx, cancel := context.WithCancel(context.Background())
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//
//	p := poker.NewPoker(logger)
//	app := application.New(logger, cfg, p)
//	go app.Run(cancel, wg, ln, lnControl)
//
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals, syscall.SIGTERM)
//	signal.Notify(signals, syscall.SIGINT)
//
//	select {
//	case <-signals:
//		logger.Info("terminate by signal")
//		cancel()
//	case <-ctx.Done():
//		logger.Info("terminate by context")
//	}
//	app.Stop()
//
//	wg.Wait()
//
//	return nil
//}
