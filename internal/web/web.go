package web

import (
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
	"syscall"
)

type Handler func(ctx *fasthttp.RequestCtx) error

type App struct {
	logger   *logrus.Logger
	mux      *router.Router
	shutdown chan os.Signal
	mw       []Middleware
}

func NewApp(logger *logrus.Logger, shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{
		logger:   logger,
		shutdown: shutdown,
		mux:      router.New(),
		mw:       mw,
	}
}

func (a *App) Handle(verb, path string, handler Handler, mw ...Middleware) {
	handler = wrapMiddleware(mw, handler)

	handler = wrapMiddleware(a.mw, handler)

	fn := func(ctx *fasthttp.RequestCtx) {
		if err := handler(ctx); err != nil {
			a.logger.Errorf("Unhandled error: %+v", err)
		}
	}

	a.mux.Handle(verb, path, fn)
}

func (a *App) Handler() fasthttp.RequestHandler {
	return a.mux.Handler
}

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
