package routers

import "github.com/gofiber/fiber/v2"

func userAccountRouters(router fiber.Router) {
	router.Post("/import", conts.UserAccount.ImportUser)
}
