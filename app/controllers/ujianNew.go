package controllers

import (
	"NewBEUjian/app/entity"
	"NewBEUjian/pkg/database"
	"NewBEUjian/pkg/tools"
	"fmt"


	"github.com/gofiber/fiber/v2"
)

// Ini yang di pakai saat ini
func AddSoalToUsersNew(c *fiber.Ctx) error {
	// Validasi ID Admin
	idAdmin, ok := c.Locals("id_admin").(int)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "ID Admin tidak valid",
		})
	}

	// Validasi Role
	role, ok := c.Locals("role").(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "Role tidak valid",
		})
	}

	// Validasi Nama
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

	// Parse body request
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "Gagal mengonversi form",
		})
	}

	// Validasi ID Ujian
	id_ujian, ok := data["id_ujian"]
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "ID Ujian tidak ditemukan",
		})
	}

	// Ambil data ujian dari database
	var Ujian entity.Ujian
	if err := database.DB.Where("id_ujian = ?", id_ujian).Preload("UsersUjian").First(&Ujian).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Pesan": "Gagal mengambil informasi ujian: " + err.Error(),
		})
	}

	// Ambil tipe ujian dari database
	var typeUjian entity.TypeUjian
	if err := database.DB.Where("id_type_ujian = ?", Ujian.IdTypeUjian).Preload("MateriBagian").First(&typeUjian).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Pesan": "Gagal mengambil tipe ujian: " + err.Error(),
		})
	}

	Ujian.IsSematkan = "true"

	waktuUjian := Ujian.WaktuUjian

	// Proses untuk setiap pengguna dalam ujian
	for _, user := range Ujian.UsersUjian {

		var existingCode entity.CodeAksesUsers
		if err := database.DB.Where("id_user_ujian = ? AND id_type_ujian = ?", user.IdUserUjian, typeUjian.IdTypeUjian).First(&existingCode).Error; err == nil {
			// Jika ditemukan, lanjutkan ke iterasi berikutnya
			continue
		}

		//Waktu Pelaksanaan
		newCodeAksesBagian := entity.CodeAksesUsers{
			IdUserUjian:    user.IdUserUjian,
			IdTypeUjian:      typeUjian.IdTypeUjian,
			KodeAkses:      tools.RandomString(6),
			CreteAt:        tools.TimeNowJakarta(),
			WaktuCodeUjian: waktuUjian,
			IsUse:          "false",
		}

		if err := database.DB.Create(&newCodeAksesBagian).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"Pesan": "Gagal membuat kode akses bagian: " + err.Error(),
			})
		}

		// Ambil soal untuk bagian dan paket tertent

		materiBagian := typeUjian.MateriBagian

		var listBagian []entity.SoalUjianTypeUjian
		for _, materi := range materiBagian {
			var soal []entity.SoalUjianTypeUjian
			database.DB.Where("id_type_ujian = ? AND materi = ? ", typeUjian.IdTypeUjian, materi.NamaMateri).Limit(materi.JumlahSoal).Find(&soal)
			fmt.Println(materi.NamaMateri, ":", len(soal), "Jumlah Soal Seharusnya :", materi.JumlahSoal)
			listBagian = append(listBagian, soal...)
		}

		// Tambahkan soal ke pengguna
		for _, soal := range listBagian {
			newUsersSoal := entity.UsersSoal{
				IdUserUjian:            user.IdUserUjian,
				IdCodeAksesUsers: newCodeAksesBagian.IdCodeAksesUsers,
				IdTypeUjian: int(soal.IdTypeUjian),
				IdSoalUjianTypeUjian:      soal.IdSoalUjianTypeUjian,
			}
			if err := database.DB.Create(&newUsersSoal).Error; err != nil {
				return c.Status(500).JSON(fiber.Map{
					"Pesan": "Gagal menambahkan soal untuk pengguna: " + err.Error(),
				})
			}

		}
	}

	return c.Status(200).JSON(fiber.Map{
		"Pesan": "Soal berhasil ditambahkan",
	})
}

func GetSoalUsersUjian(c *fiber.Ctx) error {

	id_users, _ := c.Locals("id_users").(int)
	codeAkses, _ := c.Locals("CodeAkses").(string)
	id_users_pelatihan, _ := c.Locals("id_users_pelatihan").(int)

	types, _ := c.Locals("types").(string)

	if id_users == 0 {

	}

	if types == "" {

	}

	fmt.Println(codeAkses)
	//ambil soal dari table users soal berdasarkan id users pelatihan
	var soalUsers []entity.UsersSoal

	// Id users pelatihannya dan juga id_bagiannya
	database.DB.Where("id_users_ujian = ? AND id_soal_ujian_type_ujian  = ?", id_users_pelatihan).Find(&soalUsers)

	idSoalUjian := []int64{}
	fmt.Println(idSoalUjian)

	for _, id_soal := range soalUsers {
		idSoalUjian = append(idSoalUjian, int64(id_soal.IdSoalUjianTypeUjian))
	}

	fmt.Println(idSoalUjian)

	//Mengambil soal dan jawaban dari users

	var soal []entity.SoalUjianTypeUjian

	database.DB.Unscoped().Where("id_soal_ujian_type_ujian  IN ?", idSoalUjian).Preload("Jawaban").Find(&soal)

	//ambil soal berdasarkan

	return c.JSON(fiber.Map{
		"Soal":   soal,
		"jumlah": len(idSoalUjian),
	})

}

func GetUserUjianByCodeAkses(c *fiber.Ctx) error {
	idUsersUjian, ok := c.Locals("id_users_ujian").(int)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"pesan": "ID pengguna ujian tidak valid",
		})
	}

	type Response struct {
		Nama        string `json:"nama"`
		Nik         string `json:"nik"`
		NomorUjian  string `json:"nomor_ujian"`
		IdUserUjian uint   `json:"id_user_ujian"`
		Instansi    string `json:"instansi"`
	}

	var response Response

	fmt.Println("id_users_ujian", idUsersUjian)

	err := database.DB.Raw(`SELECT nama, nik, nomor_ujian, id_user_ujian, instansi FROM users_ujians WHERE id_user_ujian = ?`, idUsersUjian).Scan(&response).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"pesan": "Gagal mendapatkan data",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "Sukses mendapatkan data",
		"data":  response,
	})
}



