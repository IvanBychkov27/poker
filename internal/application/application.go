package application

import (
	"context"
	"errors"
	"github.com/IvanBychkov27/poker/internal/config"
	"github.com/IvanBychkov27/poker/internal/poker"
	"go.uber.org/zap"
	"net"
	"net/http"
	"sync"
)

type Application struct {
	server        *http.Server
	serverControl *http.Server
	logger        *zap.Logger
	cfg           *config.Config
	p             *poker.Poker
}

func New(logger *zap.Logger, cfg *config.Config, p *poker.Poker) *Application {
	app := &Application{
		logger: logger,
		cfg:    cfg,
		p:      p,
	}
	app.p.PageTop = pageTop
	app.p.Form = form

	router := http.NewServeMux()
	router.HandleFunc("/", app.poker)
	router.HandleFunc("/pic/", app.handler)

	app.server = &http.Server{}
	app.server.Handler = router

	controlRouter := http.NewServeMux()
	controlRouter.HandleFunc("/liveness", app.liveness)

	app.serverControl = &http.Server{}
	app.serverControl.Handler = controlRouter

	return app
}

func (app *Application) Run(cancel context.CancelFunc, wg *sync.WaitGroup, ln, lnControl net.Listener) {
	defer wg.Done()

	wg.Add(2)
	go app.run(ln, cancel, wg)
	go app.runControl(lnControl, cancel, wg)
}

func (app *Application) run(ln net.Listener, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	defer cancel()

	app.logger.Info("start main server poker listen", zap.String("address", ln.Addr().String()))

	err := app.server.Serve(ln)
	if !errors.Is(http.ErrServerClosed, err) {
		app.logger.Error("error main serve poker", zap.Error(err))
	}
}

func (app *Application) runControl(ln net.Listener, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	defer cancel()

	app.logger.Info("start control server poker listen", zap.String("address", ln.Addr().String()))

	err := app.serverControl.Serve(ln)
	if !errors.Is(http.ErrServerClosed, err) {
		app.logger.Error("error control serve poker", zap.Error(err))
	}
}

func (app *Application) Stop() {
	app.logger.Info("stop main server poker...")
	err := app.server.Shutdown(context.Background())
	if err != nil {
		app.logger.Error("error stop main server poker", zap.Error(err))
	}

	app.logger.Info("stop control server poker...")
	err = app.serverControl.Shutdown(context.Background())
	if err != nil {
		app.logger.Error("error stop control server poker", zap.Error(err))
	}
}
