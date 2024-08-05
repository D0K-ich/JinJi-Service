package network

import (
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

var DefaultServer *fasthttp.Server

func cors(h fasthttp.RequestHandler, access_token string) fasthttp.RequestHandler {
	log.Info().Msg("(MW) >> Init CORS middleware...")

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
		log.Debug().Msgf("CORS EXECUTED", "origin", origin, "method", string(ctx.Method()), "url", string(ctx.Request.RequestURI()))
		ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", allowed_headers)
		ctx.Response.Header.Set("Access-Control-Expose-Headers", allowed_headers)
		if ctx.IsOptions() {return} // Log.Debug("Return from options")

		if string(ctx.Request.Header.Peek("x-access-token")) != access_token {
			ctx.Response.SetStatusCode(401)
			ctx.Response.SetBody([]byte("Unauthorized"))
			return
		}

		h(ctx)
	}
}

func writeError(ctx *fasthttp.RequestCtx, err_string string) {
	log.Error().Msgf("http preprocess error", "err_string", err_string)
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetBodyString(`{"type" : "error", "message" : "` + err_string + `"}`)
}
