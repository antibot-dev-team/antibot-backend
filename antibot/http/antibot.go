package http

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/antibot-dev-team/antibot-backend/antibot"
	"github.com/antibot-dev-team/antibot-backend/antibot/analyzer"
	"github.com/antibot-dev-team/antibot-backend/internal/web"
)

const validClient = "VALID_CLIENT"
const invalidClient = "INVALID_CLIENT"

type AntibotHandlers struct {
	Version  string
	Cfg      *antibot.Config
	Logger   *logrus.Logger
	Analyzer *analyzer.Analyzer
}

type analyzeResponse struct {
	// TODO: Think about response format
	Decision string `json:"decision"`
}

type analyzeRequest struct {
	// TODO: Think about request format
	Properties analyzer.ClientProperties `json:"properties"`
}

func (a *AntibotHandlers) Analyze(ctx *fasthttp.RequestCtx) error {
	var requestBody analyzeRequest
	if err := web.DecodeJSON(ctx, &requestBody); err != nil {
		return errors.Wrap(err, "undefined data structure")
	}

	human := a.Analyzer.AnalyzeProperties(requestBody.Properties)

	decision := validClient
	if !human {
		decision = invalidClient
	}

	return web.RespondJSON(ctx, analyzeResponse{
		Decision: decision,
	}, fasthttp.StatusOK)
}
