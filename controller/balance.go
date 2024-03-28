package controller

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ramadhan1445sprint/sprint_segokuning/customErr"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
	"github.com/ramadhan1445sprint/sprint_segokuning/svc"
)

type BalanceController struct {
	svc svc.BalanceSvc
	validate *validator.Validate
}

func NewBalanceController(svc svc.BalanceSvc, validate *validator.Validate) *BalanceController {
	return &BalanceController{
		svc: svc,
		validate: validate,
	}
}

func (c *BalanceController) AddBalance(ctx *fiber.Ctx) error {
	var balanceReq entity.AddBalanceRequest
	userId := ctx.Locals("user_id").(string)

	if err := ctx.BodyParser(&balanceReq); err != nil {
		custErr := customErr.NewBadRequestError(err.Error())
		return ctx.Status(custErr.StatusCode).JSON(fiber.Map{"message": custErr.Message})
	}

	if err := c.validate.Struct(balanceReq); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			custErr := customErr.NewBadRequestError(e.Error())
			return ctx.Status(custErr.StatusCode).JSON(fiber.Map{"message": custErr.Message})
		}
	}

	balanceReq.UserID = userId

	if err := c.svc.AddBankAccountBalance(balanceReq); err != nil {
		return ctx.Status(err.StatusCode).JSON(fiber.Map{"message": err.Message})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "success"})
}

func (c *BalanceController) GetTransactionHistory(ctx *fiber.Ctx) error {
	var filterReq entity.BalanceHistoryMeta
	userId := ctx.Locals("user_id").(string)

	limit := ctx.Query("limit")
	offset := ctx.Query("offset")

	if limit == "" && offset == "" {
		fmt.Println(limit)
		fmt.Println(offset)
		filterReq.Limit = 5
		filterReq.Offset = 0
	}else {
		custErr := customErr.NewBadRequestError("invalid query param")
		return ctx.Status(custErr.StatusCode).JSON(fiber.Map{"message": custErr.Message})
	}

	if err := ctx.QueryParser(&filterReq); err != nil {
		custErr := customErr.NewBadRequestError(err.Error())
		return ctx.Status(custErr.StatusCode).JSON(fiber.Map{"message": custErr.Message})
	}

	if err := c.validate.Struct(filterReq); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			custErr := customErr.NewBadRequestError(e.Error())
			return ctx.Status(custErr.StatusCode).JSON(fiber.Map{"message": custErr.Message})
		}
	}

	resp, err := c.svc.GetBalanceHistory(userId, filterReq)

	if err != nil {
		return ctx.Status(err.StatusCode).JSON(fiber.Map{"message": err.Message})
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}
