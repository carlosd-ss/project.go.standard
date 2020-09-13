package hcustomer

import (
	mErrors "github.com/go.standard.project.layout/project.web/fiber/crud.mongoa/models/errors"
	fmts "github.com/go.standard.project.layout/project.web/fiber/crud.mongoa/pkg/fmts"
	rcustomer "github.com/go.standard.project.layout/project.web/fiber/crud.mongoa/repo/customer"
	"github.com/gofiber/fiber"
)

func GetUuid(c *fiber.Ctx) {
	var Errors mErrors.Errors
	uuid := c.Params("uuid")
	if uuid == "" {
		Errors.Msg = "Error: Nenhum uuid informado "
		//400=Bad request
		c.Status(400).JSON(Errors)
		return
	}
	returncustomer, err := rcustomer.GetUuid(uuid)
	if err != nil {
		Errors.Msg = fmts.Concat("Error: ", err.Error())
		//404= not found
		c.Status(404).JSON(Errors)
		return
	}
	c.Status(200).Send(returncustomer)
	return
}
