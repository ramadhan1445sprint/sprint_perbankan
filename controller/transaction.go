package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramadhan1445sprint/sprint_segokuning/customErr"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
	"github.com/ramadhan1445sprint/sprint_segokuning/svc"
)

type TransactionCtr struct {
	svc svc.TransactionSvc
}

func NewTransactionCtr(svc svc.TransactionSvc) *TransactionCtr {
	return &TransactionCtr{svc: svc}
}

func (c *TransactionCtr) AddTransaction(ctx *fiber.Ctx) error {
	var payload entity.TransactionPayload
	if err := ctx.BodyParser(&payload); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	userId := ctx.Locals("user_id").(string)

	err := c.svc.AddTransaction(userId, payload)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Add Transaction Success!",
	})
}

func (c *TransactionCtr) GetBalance(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)

	listBalance, err := c.svc.GetBalance(userId)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    listBalance,
	})
}
