package web

import (
	"bytes"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// DecodeJSON decodes JSON body and stores it in the data structure
func DecodeJSON(ctx *fasthttp.RequestCtx, data interface{}) error {
	requestBody := bytes.NewReader(ctx.PostBody())
	decoder := json.NewDecoder(requestBody)

	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}
