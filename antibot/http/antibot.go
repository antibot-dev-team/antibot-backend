package http

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/antibot-dev-team/antibot-backend/antibot"
	"github.com/antibot-dev-team/antibot-backend/internal/web"
)

type AntibotHandlers struct {
	Version string
	Cfg     *antibot.Config
	Logger  *logrus.Logger
}

type analyzeResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type analyzeRequest struct {
	Data string `json:"data"`
}

func (a *AntibotHandlers) Analyze(ctx *fasthttp.RequestCtx) error {
	var requestBody analyzeRequest
	if err := web.DecodeJSON(ctx, &requestBody); err != nil {
		return errors.Wrap(err, "undefined data structure")
	}
	return web.RespondJSON(ctx, analyzeResponse{
		Status: "success",
		Data:   requestBody.Data,
	}, fasthttp.StatusOK)
}
