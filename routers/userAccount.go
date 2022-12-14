package routers

import "github.com/gofiber/fiber/v2"

func userAccountRouters(router fiber.Router) {
	router.Post("/import", conts.UserAccount.ImportUser)
	router.Get("/", conts.UserAccount.FindUsers)
	router.Post("/upVote", conts.UserAccount.VoteUserUp)
	router.Post("/downVote", conts.UserAccount.VoteUserDown)
}
