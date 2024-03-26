package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramadhan1445sprint/sprint_segokuning/customErr"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
	"github.com/ramadhan1445sprint/sprint_segokuning/svc"
)

type UserController struct {
	svc svc.UserSvc
}

func NewUserController(svc svc.UserSvc) *UserController {
	return &UserController{svc: svc}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var newUser entity.RegistrationPayload
	if err := ctx.BodyParser(&newUser); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	accessToken, err := c.svc.Register(newUser)
	if err != nil {
		return err
	}

	respData := fiber.Map{
		"email":       newUser.Email,
		"name":        newUser.Name,
		"accessToken": accessToken,
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    respData,
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var user entity.Credential
	if err := ctx.BodyParser(&user); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	name, accessToken, err := c.svc.Login(user)
	if err != nil {
		return err
	}

	respData := fiber.Map{
		"email":       user.Email,
		"name":        name,
		"accessToken": accessToken,
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    respData,
	})
}
