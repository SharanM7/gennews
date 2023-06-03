package utils

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

func SetResponse(ctx *fasthttp.RequestCtx, code int, s string) {
	res := Response{
		Message: s,
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Error while marshalling response", err)
		data = []byte(`{message:"internal error"}`)
		code = 503
	}

	ctx.SetStatusCode(code)
	ctx.Write(data)
}
