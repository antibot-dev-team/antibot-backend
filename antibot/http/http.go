package http

import (
	"os"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/antibot-dev-team/antibot-backend/antibot"
	"github.com/antibot-dev-team/antibot-backend/internal/middlewares"
	"github.com/antibot-dev-team/antibot-backend/internal/web"
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

	app.Handle(fasthttp.MethodPost, "/api/v1/analyze", antibotHandlers.Analyze)

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: false,
		AllowMaxAge:      5600,
		Debug:            false,
	})

	return withCors.CorsMiddleware(app.Handler())
}
