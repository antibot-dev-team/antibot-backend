package middlewares

import (
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/antibot-dev-team/antibot-backend/internal/web"
)

var errorStatus = struct {
	Error string `json:"error"`
}{
	"unexpected error",
}

// Errors handles unexpected errors
func Errors(logger *logrus.Logger) web.Middleware {

	m := func(before web.Handler) web.Handler {

		h := func(ctx *fasthttp.RequestCtx) error {
			if err := before(ctx); err != nil {
				logger.Error(err)
				ctx.ResetBody()
				return web.RespondJSON(ctx, errorStatus, fasthttp.StatusInternalServerError)
			}

			return nil
		}

		return h
	}

	return m
}
