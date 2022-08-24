package controllers

import (
	"github.com/chirag3003/vote-back/models"
	"github.com/chirag3003/vote-back/repository"
	"github.com/gofiber/fiber/v2"
)

type UserAccountsController interface {
	ImportUser(ctx *fiber.Ctx) error
}

type userRoutes struct {
	userAcc repository.UserAccountRepository
}

func userAccountController() UserAccountsController {
	return &userRoutes{
		userAcc: repository.NewUserAccountRepository(conn.DB),
	}
}

func (u *userRoutes) ImportUser(ctx *fiber.Ctx) error {
	body := &models.UserAccountImportInput{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	username, err := u.userAcc.FindByUsername(body.UserName)
	if username != nil {
		err = u.userAcc.UpdateAccount(body.UserName, body)

	} else {
		err = u.userAcc.CreateAccount(body)
	}
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
