package hcustomer

import (
	mErrors "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/models/errors"
	fmts "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/pkg/fmts"
	rcustomer "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/repo/customer"
	"github.com/gofiber/fiber"
)

func (s *Server) GetUuid(c *fiber.Ctx) {
	var Errors mErrors.Errors
	uuid := c.Params("uuid")
	if uuid == "" {
		Errors.Msg = "Error: Nenhum uuid informado "
		//400=Bad request
		c.Status(400).JSON(Errors)
		return
	}
	returncustomer, err := rcustomer.GetUuid(s.Db, uuid)
	if err != nil {
		Errors.Msg = fmts.Concat("Error: ", err.Error())
		//404= not found
		c.Status(404).JSON(Errors)
		return
	}
	c.Status(200).Send(returncustomer)

}
