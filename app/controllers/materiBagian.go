package controllers

import (
	"NewBEUjian/app/entity"
	"NewBEUjian/pkg/database"
	"NewBEUjian/pkg/tools"

	"github.com/gofiber/fiber/v2"
)

func GetMateriBagianByIdBagian(c *fiber.Ctx) error {

	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var materiBagian []entity.MateriBagian

	idBagian := c.Query("id_type_ujian")

	database.DB.Where("id_type_ujian = ?", idBagian).Find(&materiBagian)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Mengambil Data",
		"data":  materiBagian,
	})
}



func CreateMateriBagian(c *fiber.Ctx) error {

	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	idBagian := c.Query("id_bagian")

	if idBagian != "" {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "Query Id Bagian Wajib Di Isi",
		})
	}

	var Request entity.MateriBagian


	err := c.BodyParser(&Request)
	if err != nil {
		// Handle the error (for example, return a 400 Bad Request response)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})

	}

	materiBagian := entity.MateriBagian{
		IdTypeUjian: Request.IdTypeUjian,
		NamaMateri: Request.NamaMateri,
		JumlahSoal: Request.JumlahSoal,
		CreateAt: tools.TimeNowJakarta(),
	}

	database.DB.Create(&materiBagian)


	return c.JSON(fiber.Map{
		"Pesan":"Berhasil Membuat Materi Bagian",
	})
}

func UpdateMateriBagian(c *fiber.Ctx) error {
    // Mendapatkan ID admin, peran, dan nama dari konteks
    idAdmin, ok := c.Locals("id_admin").(int)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
    role, ok := c.Locals("role").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
    names, ok := c.Locals("name").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Validasi JWT
    if err := tools.ValidationJwt(c, role, idAdmin, names); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(err)
    }

    // Mendapatkan ID materi bagian dari parameter URL
    id := c.Query("id")
    if id == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak boleh kosong"})
    }

    // Mengonversi ID menjadi integer

    // Mencari materi bagian berdasarkan ID
    var materiBagian entity.MateriBagian
    if err := database.DB.First(&materiBagian, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Materi Bagian tidak ditemukan"})
    }

    // Parsing data dari body permintaan
    var request entity.MateriBagian
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Data tidak valid"})
    }

    // Memperbarui field yang diperlukan
    materiBagian.NamaMateri = request.NamaMateri
    materiBagian.JumlahSoal = request.JumlahSoal

    // Menyimpan perubahan ke database
    if err := database.DB.Save(&materiBagian).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui Materi Bagian"})
    }

    return c.JSON(fiber.Map{
        "Pesan": "Berhasil memperbarui Materi Bagian",
        "data":  materiBagian,
    })
}

func DeleteMateriBagian(c *fiber.Ctx) error {
    // Mendapatkan ID admin, peran, dan nama dari konteks
    idAdmin, ok := c.Locals("id_admin").(int)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
    role, ok := c.Locals("role").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
    names, ok := c.Locals("name").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Validasi JWT
    if err := tools.ValidationJwt(c, role, idAdmin, names); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(err)
    }

    // Mendapatkan ID materi bagian dari parameter URL
    id := c.Query("id")
    if id == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak boleh kosong"})
    }

    // Mengonversi ID menjadi integer
 
    // Mencari materi bagian berdasarkan ID
    var materiBagian entity.MateriBagian
    if err := database.DB.First(&materiBagian, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "MateriBagian tidak ditemukan"})
    }

    // Menghapus materi bagian
    if err := database.DB.Delete(&materiBagian).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus MateriBagian"})
    }

    return c.JSON(fiber.Map{
        "Pesan": "Berhasil menghapus MateriBagian",
    })
}