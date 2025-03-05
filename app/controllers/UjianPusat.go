package controllers

import (
	"NewBEUjian/app/entity"
	"NewBEUjian/pkg/config"
	"NewBEUjian/pkg/database"
	"NewBEUjian/pkg/tools"
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func CreateUjian(c *fiber.Ctx) error {

	/*

		//Pake Role Super admin/ admin pusat
		id_admin, _ := c.Locals("id_admin").(int)
		role, _ := c.Locals("role").(string)
		names, _ := c.Locals("name").(string)

		if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
			// Jika ada kesalahan, kirim pesan kesalahan
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
	*/
	var request entity.Ujian

	if err := c.BodyParser(&request); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	filePermohonan, _ := c.FormFile("filePermohonan")
	// Memeriksa ekstensi file

	ext := filepath.Ext(filePermohonan.Filename)
	if ext != ".pdf" {
		return fmt.Errorf("file harus berupa pdf")
	}

	ujianNew := entity.Ujian{
		IdTypeUjian:   request.IdTypeUjian, // Use `request.IdTypeUjian` instead of `data`
		TypeUjian:     request.TypeUjian,   // Use `request.TypeUjian`
		NamaUjian:     request.NamaUjian,   // Use `request.NamaUjian`
		TempatUjian:   request.TempatUjian,
		IdLemdikat:    request.IdLemdikat,
		LembagaDiklat: request.LembagaDiklat, // Use `request.TempatUjian`
		// Use `request.PUKAKP`
		NamaPengawasUjian:    request.NamaPengawasUjian,    // Use `request.NamaPengawas`
		NamaVasilitatorUjian: request.NamaVasilitatorUjian, // Use `request.NamaVasilitator`
		TanggalMulaiUjian:    request.TanggalMulaiUjian,    // Use `request.TanggalMulai`
		TanggalBerakhirUjian: request.TanggalBerakhirUjian, // Use `request.TanggalBerakhir`
		WaktuUjian:           request.WaktuUjian,           // Convert string to int (WaktuUjian)
		JumlahPesertaUjian:   request.JumlahPesertaUjian,   // Convert string to int (JumlahPeserta)
		Status:               "Pending",                    // Use `request.Status`
		CreateAt:             tools.TimeNowJakarta(),
		FilePermohonan:       tools.RemoverSpaci(filePermohonan.Filename),
		IsSematkan:           "false", // Set current time in Jakarta timezone
	}

	if err := database.DB.Create(&ujianNew).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"pesan": "Gagal menyimpan data ke database", "error": err.Error()})
	}

	//Simpan file
	if filePermohonan != nil {
		if err := c.SaveFile(filePermohonan, "public/file-permohonan/"+tools.RemoverSpaci(filePermohonan.Filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": "Failed to save Buku Pelaut", "Error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Membuat Ujian Terbaru",
	})
}

func UpdateUjian(c *fiber.Ctx) error {
	// Pake Role Super admin/ admin pusat
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	id := c.Query("id")
	var data entity.Ujian

	// Start a transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Find the Ujian by its ID
	if err := tx.Where("id_ujian = ?", id).First(&data).Error; err != nil {
		tx.Rollback() // Tidak perlu rollback jika tx belum dikommit
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ujian not found"})
	}

	// Parse incoming JSON data into the request map
	var request map[string]string
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Haude"})
	}

	// Prepare the updated Ujian data
	newUpdate := entity.Ujian{

		// Use `request.TypeUjian`
		NamaUjian:   request["nama_ujian"], // Use `request.NamaUjian`
		TempatUjian: request["tempat_ujian"],

		TanggalMulaiUjian:    request["tanggal_mulai_ujian"],                     // Use `request.TanggalMulai`
		TanggalBerakhirUjian: request["tanggal_berakhir_ujian"],                  // Use `request.TanggalBerakhir`
		WaktuUjian:           request["waktu_ujian"],                             // Convert string to int (WaktuUjian)
		JumlahPesertaUjian:   tools.StringToInt(request["jumlah_peserta_ujian"]), // Convert string to int (JumlahPeserta)

		NamaPengawasUjian:    request["nama_pengawas_ujian"],
		IsSematkan:           request["is_sematkan"],
		NamaVasilitatorUjian: request["nama_vasilitator_ujian"],

		Status:   request["status"],      // Use the existing create time
		UpdateAt: tools.TimeNowJakarta(), // Set current time for update
	}

	// Update the record with new values
	if err := tx.Model(&data).Updates(newUpdate).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update Ujian data"})
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	// Return the updated data
	return c.JSON(fiber.Map{
		"Pesan": "Data berhasil diperbarui",
		"data":  newUpdate,
	})
}

func GetUjian(c *fiber.Ctx) error {

	viper := config.NewViper()
	baseUrl := viper.GetString("web.baseUrl")
	//Pake Role Super admin/ admin pusat
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwtUsersDewan(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	id := c.Query("id")
	id_users_dpkakp := c.Query("id_users_dpkakp")

	var ujian []entity.Ujian
	//penambahan filter lainnya
	baseQuery := database.DB

	if id != "" {
		baseQuery = baseQuery.Where("id_ujian = ?", id)
	}

	if id_users_dpkakp != "" {
		baseQuery = baseQuery.Where("id_users_dpkakp =?", id_users_dpkakp)
	}

	err := baseQuery.Preload("UsersUjian.CodeAksesUsersBagian").Find(&ujian)

	for i, _ := range ujian {
		ujian[i].FilePermohonan = baseUrl + "/public/file-permohonan/" + ujian[i].FilePermohonan
	}

	fmt.Println(err)
	return c.JSON(fiber.Map{
		"Pesan": "Sukses Get Ujian",
		"data":  ujian,
	})
}

func DeleteUjians(c *fiber.Ctx) error {
	// Ambil data dari JWT
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	// Validasi JWT
	if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid JWT or permissions",
			"details": err,
		})
	}

	fmt.Println("Delete")

	// Ambil parameter ID dari query string
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID parameter is required",
		})
	}

	var data entity.Ujian

	// Mulai transaksi
	database.DB.Where("id_ujian = ?", id).Preload("UsersUjian.CodeAksesUsersBagian").First(&data)

	// Hapus data ujian
	database.DB.Delete(&data)

	// Berhasil
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ujian deleted successfully",
	})
}
