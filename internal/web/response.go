package web

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// RespondJSON responds with JSON body
func RespondJSON(ctx *fasthttp.RequestCtx, data interface{}, statusCode int) error {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	return json.NewEncoder(ctx).Encode(data)
}
