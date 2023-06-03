package utils

import (
	"log"
	"net/http"
	"net/url"
	"path"
	"sync"

	"github.com/valyala/fasthttp"
)

var server fasthttp.Server

func NewServer(p string) *Server {
	return &Server{
		Port: p,
	}
}

func (s *Server) Start(wg *sync.WaitGroup) {
	server = fasthttp.Server{
		// RemoteAddr: s.Port,
		Handler: HandleRequest,
	}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := server.ListenAndServe(s.Port); err != nil {
			log.Println("http server stopped. Error :", err)
		}

	}(wg)
}

func (s *Server) Stop() {
	server.Shutdown()
}

func HandleRequest(ctx *fasthttp.RequestCtx) {
	uri := ctx.URI().String()
	basePath := path.Base(uri)
	u, _ := url.Parse(basePath)

	log.Printf("reportInitializer - request IP : %v | Endpoint : %v | Data : %v\n", ctx.RemoteAddr(), basePath, string(ctx.PostBody()))

	switch u.Path {
	case "summarize":
		summarizeNews(ctx)
	default:
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.WriteString(`{"statusCode":404,"message":"incorrect url"}`)
	}
}
