package controllers

import (
	"NewBEUjian/app/entity"
	"NewBEUjian/pkg/config"
	"NewBEUjian/pkg/database"
	"NewBEUjian/pkg/tools"

	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

//Menampilkan CDN gambar soal

func UploadGambarSoal(c *fiber.Ctx) error {
	// Inisialisasi Viper untuk mengambil konfigurasi
	viper := config.NewViper()
	baseUrl := viper.GetString("web.baseUrl")

	// Mengambil file gambar dari form
	foto, err := c.FormFile("foto")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File tidak ditemukan"})
	}

	// Memeriksa ekstensi file
	ext := filepath.Ext(foto.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File harus berupa gambar (JPG, JPEG, PNG)"})
	}

	if foto != nil {
		if err := c.SaveFile(foto, "public/soal-pelatihan/"+tools.RemoverSpaci(foto.Filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": "Failed to save foto", "Error": err.Error()})
		}
	}
	// Simpan informasi gambar ke database
	newGambar := entity.GambarSoal{
		CodeUnik: "code UniK", // Menyimpan nama file
		Gambar:   baseUrl + "/public/soal-pelatihan/" + tools.RemoverSpaci(foto.Filename),
	}

	if err := database.DB.Create(&newGambar).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan informasi gambar ke database"})
	}

	// Mengembalikan URL gambar yang diupload
	return c.JSON(fiber.Map{
		"Pesan": "Sukses Upload Gambar",
		"Url":   newGambar.Gambar,
	})
}

func GetGambar(c *fiber.Ctx) error {

	var gambar []entity.GambarSoal

	database.DB.Find(&gambar)
	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Ambil Data",
		"data":  gambar,
	})
}

func UpdateSoal(c *fiber.Ctx) error {

	// Validasi role Super Admin atau Admin Pusat
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwtTenagaAhli(c, role, id_admin, names); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi gagal",
		})
	}

	// Ambil ID soal dari query parameter
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID soal harus disertakan",
		})
	}

	// Ambil data soal ujian berdasarkan ID
	var soal entity.SoalUjianTypeUjian
	if err := database.DB.Where("id_soal_ujian_type_ujian  = ?", id).Preload("Jawaban").First(&soal).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Soal tidak ditemukan",
			"error":   err.Error(),
		})
	}

	// Parsing body untuk mendapatkan data baru
	var request ResponDatas

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal memproses input",
			"error":   err.Error(),
		})
	}

	// Perbarui data soal
	updates := entity.SoalUjianTypeUjian{
		Soal:         request.NamaSoal,
		GambarSoal:   tools.RemoverSpaci(request.GambarSoal),
		JawabanBenar: request.JawabanBenar,
		UpdateAt:     tools.TimeNowJakarta(),
	}

	if err := database.DB.Model(&soal).Updates(&updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memperbarui soal",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data soal berhasil diperbarui",
		"data":    soal,
	})
}

func DeleteSoal(c *fiber.Ctx) error {

	// Validasi role Super Admin atau Admin Pusat
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwtTenagaAhli(c, role, id_admin, names); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi gagal",
		})
	}

	// Ambil ID soal dari query parameter
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID soal harus disertakan",
		})
	}

	// Ambil data soal ujian berdasarkan ID
	var soal entity.SoalUjianTypeUjian
	if err := database.DB.Where("id_soal_ujian_type_ujian = ?", id).Preload("Jawaban").First(&soal).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Soal tidak ditemukan",
			"error":   err.Error(),
		})
	}

	database.DB.Delete(&soal)

	return c.JSON(fiber.Map{
		"Pesan": "Sukses Menghapus Pesan",
	})
}

func CreateSoal(c *fiber.Ctx) error {

	// Validasi role Super Admin atau Admin Pusat
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwtTenagaAhli(c, role, id_admin, names); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi gagal",
		})
	}

	return nil
}
