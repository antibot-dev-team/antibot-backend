package http

import (
	"github.com/antibot-dev-team/antibot-backend/antibot"
	"github.com/antibot-dev-team/antibot-backend/internal/middlewares"
	"github.com/antibot-dev-team/antibot-backend/internal/web"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
)

func HandleAPI(version string,
	logger *logrus.Logger,
	shutdown chan os.Signal,
	cfg *antibot.Config,
) fasthttp.RequestHandler {
	app := web.NewApp(logger, shutdown, middlewares.Errors(logger))

	antibotHandlers := AntibotHandlers{
		Version: version,
		Cfg:     cfg,
		Logger:  logger,
	}

	app.Handle(fasthttp.MethodPost, "/v1/analyze", antibotHandlers.Analyze)

	return app.Handler()
}
