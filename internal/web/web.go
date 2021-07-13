package web

import (
	"os"
	"syscall"

	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Handler func(ctx *fasthttp.RequestCtx) error

type App struct {
	logger      *logrus.Logger
	router      *router.Router
	shutdown    chan os.Signal
	middlewares []Middleware
}

// NewApp creates new application
func NewApp(logger *logrus.Logger, shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{
		logger:      logger,
		shutdown:    shutdown,
		router:      router.New(),
		middlewares: mw,
	}
}

// Handle operates with handlers and middlewares around them
func (a *App) Handle(verb, path string, handler Handler, mw ...Middleware) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.middlewares, handler)
	wrapFunction := func(ctx *fasthttp.RequestCtx) {
		if err := handler(ctx); err != nil {
			a.logger.Error(err)
		}
	}

	a.router.Handle(verb, path, wrapFunction)
}

// Handler wraps fasthttp router handler
func (a *App) Handler() fasthttp.RequestHandler {
	return a.router.Handler
}

// SignalShutdown sends SIGTERM to the shutdown channel
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
