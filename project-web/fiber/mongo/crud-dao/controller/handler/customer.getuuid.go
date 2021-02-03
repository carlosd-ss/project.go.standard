package handler

import (
	"github.com/gofiber/fiber"
	"github.com/project.go.standard/project-web/fiber/mongo/crud-dao/repo"

	merrors "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/models/errors"
)

func (s *Conn) GetUuid(c *fiber.Ctx) {
	var msgE merrors.Errors
	id := c.Params("id")
	c.Set("Content-Type", "application/json")

	if len(c.Params("id")) <= 0 {
		msgE.Msg = "Id é obrigatorio "
		c.Status(400).JSON(msgE)
		return
	}

	err := repo.Delete(id, s.Db, s.CollectionName)
	if err != nil {
		msgE.Msg = err.Error()
		c.Status(400).JSON(msgE)
		return
	}

	c.Status(200).SendString("")
	return
}
