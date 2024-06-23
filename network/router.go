package network

import (
	"github.com/D0K-ich/KanopyService/handlers/admins"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"runtime/debug"
)

const (
	cookieKeyUser   = "user_id"
	cookieKeyAdmin  = "admin_id"
)

func CreateRouter(files_path string,  admin_token string) (main_router *router.Router) {
	main_router = router.New()

	main_router.PanicHandler = func(ctx *fasthttp.RequestCtx, i interface{}) {
		debug.PrintStack()
		log.Error("Panic detected", zap.Any("err", i))
		ctx.Error("Internal error", fasthttp.StatusInternalServerError)
		return
	}

	main_router.NotFound = func(ctx *fasthttp.RequestCtx) {
		log.Info("(rest) >> NOT FOUND HANDLER", zap.Any("method", string(ctx.Method())), zap.Any("path", string(ctx.Path())))
		if ctx.IsGet() {return}
		writeError(ctx, "handler not found")
		return
	}

	main_router.MethodNotAllowed = func(ctx *fasthttp.RequestCtx) {
		log.Info("(rest) >> NOT ALLOWED METHOD", zap.Any("method", string(ctx.Method())), zap.Any("path", string(ctx.Path())))
		writeError(ctx, "method not allowed")
		return
	}

	main_router.POST("/admin/{subject}/{action?}", func(ctx *fasthttp.RequestCtx) {
		var response
		response = admins.NewHandler()

		ctx.SetStatusCode(response.StatusCode())
		ctx.SetBody(response.Serialize())
		ctx.SetContentType(message.ContentTypeJson)
	})



	return
}

func fsHandler(root string, stripSlashes int) fasthttp.RequestHandler {
	var fs = &fasthttp.FS{
		Root:               root,
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: false,
		AcceptByteRange:    false,
	}
	if stripSlashes > 0 {
		fs.PathRewrite = fasthttp.NewPathSlashesStripper(stripSlashes)
	}
	return fs.NewRequestHandler()
}
