package web

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func DecodeBody(ctx *fasthttp.RequestCtx, data interface{}) error {
	requestBody := bytes.NewReader(ctx.PostBody())
	decoder := json.NewDecoder(requestBody)

	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}
