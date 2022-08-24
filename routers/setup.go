package routers

import (
	"github.com/chirag3003/vote-back/controllers"
	"github.com/gofiber/fiber/v2"
)

var conts *controllers.Controllers

func Setup(controllers *controllers.Controllers, app *fiber.App) {
	conts = controllers
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})
	userAccountRouters(app.Group("/userAccount"))
}
