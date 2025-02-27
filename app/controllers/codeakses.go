package controllers

import (
	"NewBEUjian/pkg/tools"

	"github.com/gofiber/fiber/v2"
)

func GetCodeAcsess(c *fiber.Ctx) error {

	idAdmin, ok := c.Locals("id_admin").(int)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "ID Admin tidak valid",
		})
	}
	role, ok := c.Locals("role").(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "Role tidak valid",
		})
	}
	names, ok := c.Locals("name").(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "Name tidak valid",
		})
	}

	// Validasi JWT dan akses
	if err := tools.ValidationJwt(c, role, idAdmin, names); err != nil {
		return c.Status(403).JSON(fiber.Map{
			"Pesan": "Unauthorized",
		})

	}

	return c.JSON(fiber.Map{
		"pesan": "berhasil get data kode akses",
		"data":  "",
	})
}
