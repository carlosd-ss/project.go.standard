package hcustomer

import (
	fmts "github.com/go.standard.project.layout/project.web/fiber/crud.postgresa/internal/fmts"
	mErrors "github.com/go.standard.project.layout/project.web/fiber/crud.postgresa/models/errors"
	rcustomer "github.com/go.standard.project.layout/project.web/fiber/crud.postgresa/repo/customer"
	"github.com/gofiber/fiber"
)

func GetAll(c *fiber.Ctx) {
	var Errors mErrors.Errors
	if c.Params("offset") == "" {
		Errors.Msg = "Error: offset nao informado"
		//bad request
		c.Status(400).JSON(Errors)
		return
	}
	if c.Params("limit") == "" {
		Errors.Msg = "Error: limit nao informado"
		//bad request
		c.Status(400).JSON(Errors)
		return
	}
	returncustomer, err := rcustomer.GetAll(c.Params("offset"), c.Params("limit"))
	if err != nil {
		Errors.Msg = fmts.Concat("Error: ", err.Error())
		//404= not found
		c.Status(404).JSON(Errors)
		return
	}
	c.Status(200).Send(returncustomer)
	return
}
