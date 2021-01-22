package hcustomer

import (
	mcustomer "github.com/go.standard.project.layout/project.web/fiber/crud.postgres.dao/models/customer"
	mErrors "github.com/go.standard.project.layout/project.web/fiber/crud.postgres.dao/models/errors"
	fmts "github.com/go.standard.project.layout/project.web/fiber/crud.postgres.dao/pkg/fmts"
	rcustomer "github.com/go.standard.project.layout/project.web/fiber/crud.postgres.dao/repo/customer"
	"github.com/gofiber/fiber"
)

func (s *Server) Post(c *fiber.Ctx) {
	var input mcustomer.CustomerPost
	var Errors mErrors.Errors
	if err := c.BodyParser(&input); err != nil {
		Errors.Msg = err.Error()
		c.Status(503).JSON(Errors)
		return
	}
	input.ImpIp = c.IP()
	if err := rcustomer.Post(s.Db, input); err != nil {
		Errors.Msg = fmts.Concat("Error: ", err.Error())
		//503 Serviço indisponivel
		c.Status(503).JSON(Errors)
		return
	}
	c.Status(200)

}
