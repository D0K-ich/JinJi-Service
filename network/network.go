package network

import (
	"go.uber.org/zap"
	"strings"

	"github.com/D0K-ich/KanopyService/logs"
	"github.com/valyala/fasthttp"
)

var log = logs.NewLog()

var DefaultServer *fasthttp.Server

func cors(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	log.Info("(MW) >> Init CORS middleware...")

	var allowed_headers = strings.Join([]string{
		"origin",
		"content-type",
		"x-uuid",
		"x-filename",
		"x-link-uuid",
		"x-client-uuid",
		"x-access-token",
	}, ", ")

	return func(ctx *fasthttp.RequestCtx) {
		var origin string
		if origin = string(ctx.Request.Header.Peek("Origin")); origin == "" {origin = "*"}
		//Log.Debug("CORS EXECUTED", "origin", origin, "method", string(ctx.Method()), "url", string(ctx.Request.RequestURI()))
		ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", allowed_headers)
		ctx.Response.Header.Set("Access-Control-Expose-Headers", allowed_headers)
		if ctx.IsOptions() {return}	// Log.Debug("Return from options")
		h(ctx)
	}
}

func writeError(ctx *fasthttp.RequestCtx, err_string string) {
	log.Error("http preprocess error", zap.Any("err_string", err_string))
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetBodyString(`{"type" : "error", "message" : "` + err_string +`"}`)
}