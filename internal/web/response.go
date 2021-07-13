package web

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func Respond(ctx *fasthttp.RequestCtx, data interface{}, statusCode int) error {
	if statusCode == fasthttp.StatusNoContent {
		ctx.SetStatusCode(statusCode)
		return nil
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	return json.NewEncoder(ctx).Encode(data)
}
