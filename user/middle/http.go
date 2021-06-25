package middle

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/liov/tiga/utils/log"
)

func Log(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.RequestURI)
}

func FiberLog(ctx *fiber.Ctx) error {
	log.Debug(ctx.BaseURL())
	return ctx.Next()
}
