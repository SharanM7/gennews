package utils

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

func summarizeNews(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()

	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		log.Println("Error while unmarshaling data to request")
		SetResponse(ctx, 400, "Incorrect request")
	}

	log.Printf("got req : %+v", req)
	ctx.WriteString(`{"statusCode":200,"message":""}`)
}
