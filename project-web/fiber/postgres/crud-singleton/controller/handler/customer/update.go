package hcustomer

import (
	"github.com/gofiber/fiber"

	mErrors "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/models/errors"

	fmts "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/internal/fmts"
	mcustomer "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/models/customer"
	rcustomer "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/repo/customer"
)

func Update(c *fiber.Ctx) {
	var input mcustomer.Customer
	var Errors mErrors.Errors
	if err := c.BodyParser(&input); err != nil {
		Errors.Msg = err.Error()
		c.Status(503).JSON(Errors)
		return
	}
	input.ImpIp = c.IP()
	if err := rcustomer.Update(input); err != nil {
		Errors.Msg = fmts.Concat("Error: ", err.Error())
		//503 Serviço indisponivel
		c.Status(503).JSON(Errors)
		return
	}
	c.Status(200)
	return

}
