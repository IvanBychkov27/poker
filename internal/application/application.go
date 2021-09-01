package application

import (
	"context"
	"errors"
	"github.com/IvanBychkov27/poker/internal/config"
	"go.uber.org/zap"
	"net"
	"net/http"
	"sync"
)

type Application struct {
	server *http.Server
	logger *zap.Logger
	cfg    *config.Config
}

func New(logger *zap.Logger, cfg *config.Config) *Application {
	app := &Application{
		logger: logger,
		cfg:    cfg,
	}

	router := http.NewServeMux()
	router.HandleFunc("/", app.mainpoker)

	app.server = &http.Server{}
	app.server.Handler = router

	return app
}

func (app *Application) Run(cancel context.CancelFunc, wg *sync.WaitGroup, ln net.Listener) {
	defer wg.Done()
	defer cancel()

	app.logger.Info("start server poker listen", zap.String("address", ln.Addr().String()))

	err := app.server.Serve(ln)
	if !errors.Is(http.ErrServerClosed, err) {
		app.logger.Error("error serve poker", zap.Error(err))
	}
}

func (app *Application) Stop() {
	app.logger.Info("stop server poker...")
	err := app.server.Shutdown(context.Background())
	if err != nil {
		app.logger.Error("error stop server poker", zap.Error(err))
	}
}
