package controllers

import (
	"NewBEUjian/app/entity"
	"NewBEUjian/pkg/database"
	"NewBEUjian/pkg/tools"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// hanya admin Pusat/DJKAP yang dapat mengenerate ujian  ini
func CreateTypeUjian(c *fiber.Ctx) error {

	//Pake Role Super admin/ admin pusat
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	newTypes := entity.TypeUjian{
		NamaTypeUjian: data["typeUjian"],
		CreateAt: tools.TimeNowJakarta(),
	}

	err := database.DB.Create(&newTypes)
	fmt.Println(err)
	return c.JSON(fiber.Map{
		"Pesan": "Terlah Berhasil Membuat Type Ujian",
		"data":  newTypes,
	})
}

func UpdateTypeUjian(c *fiber.Ctx) error {

	//Pake Role Super admin/ admin pusat
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var data entity.TypeUjian

	id := c.Query("id")

	var datas map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return c.JSON(err)
	}

	database.DB.Where("id_type_ujian = ? ", id).Find(&data)

	newTypes := entity.TypeUjian{
		NamaTypeUjian: datas["typeUjian"],
	}

	database.DB.Model(&data).Updates(&newTypes)

	return c.JSON(fiber.Map{
		"Pesan": "Telah Berhasil Update Ujian",
		"data":  data,
	})
}

func GetTypeUjian(c *fiber.Ctx) error {
	//Pake Role Super admin/ admin pusat
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	id := c.Query("id")

	var data []entity.TypeUjian

	baseQuery := database.DB

	if id != "" {
		baseQuery = baseQuery.Where("id_type_ujian = ? ", id)
	}

	err := baseQuery.Preload("Ujian.UsersUjian.CodeAksesUsersBagian").Preload("Soal.Jawaban").Preload("MateriBagian").Find(&data)
	fmt.Println(err)
	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Mendapatkan data",
		"data":  data,
	})

}

func DeleteTypeUjian(c *fiber.Ctx) error {

	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Pake Id aja Untuk menghapusnya
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	var typeUjian entity.TypeUjian

	if err := database.DB.Where("id_type_ujian = ?", id).Find(&typeUjian).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "TypeUjian not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	if err := database.DB.Delete(&typeUjian).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}
