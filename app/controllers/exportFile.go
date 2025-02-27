package controllers

import (
	"NewBEUjian/app/entity"
	"NewBEUjian/pkg/database"

	"NewBEUjian/pkg/tools"
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

//Seharusnya sudah done

func ExportPesertaPelatihan(c *fiber.Ctx) error {

	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	names, _ := c.Locals("name").(string)

	if err := tools.ValidationJwt(c, role, id_admin, names); err != nil {
		// Jika ada kesalahan, kirim pesan kesalahan
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Membaca file Excel dari request
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// Memeriksa ekstensi file
	ext := filepath.Ext(file.Filename)
	if ext != ".xlsx" && ext != ".xls" {
		return fmt.Errorf("file harus berupa file Excel (.xlsx atau .xls)")
	}

	// Membuka file Excel
	excelFile, err := file.Open()
	if err != nil {
		return err
	}
	defer excelFile.Close()

	// Membaca file Excel menggunakan excelize
	f, err := excelize.OpenReader(excelFile)
	if err != nil {
		return err
	}

	// Mendapatkan nama semua sheet dalam file Excel
	sheets := f.GetSheetList()

	// Membaca data dari sheet pertama
	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return err
	}

	var models entity.Ujian

	if err := c.BodyParser(&models); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Failed to parse request body", "Error": err.Error()})
	}

	idPelatihan := models.IdUjian

	var ujian entity.Ujian
	database.DB.Where("id_ujian =?", idPelatihan).Find(&ujian)

	//var userPelatihanList []entity.UsersPelatihan
	var users []entity.Users

	for _, rowUsers := range rows[1:] {
		user := entity.Users{}

		for i, columnName := range rows[0] {
			if i >= len(rowUsers) {
				// Jika indeks i melebihi panjang rowData, lanjutkan ke baris berikutnya
				continue
			}
			switch columnName {
			case "nama":
				user.Nama = rowUsers[i]
			case "no_telpon":
				user.NoTelpon = tools.StringToInt(rowUsers[i])
			case "email":
				user.Email = rowUsers[i]
			case "lemdiklat":
				user.Instansi = rowUsers[i]
			case "no_ujian":
				user.NomorUjian = rowUsers[i]

			case "nik":
				user.Nik = tools.StringToInt(rowUsers[i])
			case "tempat_lahir":
				user.TempatLahir = rowUsers[i]
			case "tanggal_lahir":
				user.TanggalLahir = rowUsers[i]
			case "jenis_kelamin":
				user.JenisKelamin = rowUsers[i]
			}

			//Masukan List Ke dalam Users List  List

		}
		users = append(users, user)

	}

	//Bandingkan
	if len(users) >= ujian.JumlahPesertaUjian {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Jumlah peserta melebihi batas peserta ujian",
		})
	}

	//

	for _, AllUsers := range users {

		//Buat Id Users terlebih Dahulu
		err := database.DB.Create(&AllUsers)
		fmt.Println(err)
		var Pelatihan entity.Ujian
		database.DB.Where("id_ujian =?", idPelatihan).Find(&Pelatihan)

		dataUsers := entity.UsersUjian{

			IdUjian:      idPelatihan,
			Nama:         AllUsers.Nama,
			Nik:          tools.IntToString(AllUsers.Nik),
			NomorUjian:   AllUsers.NomorUjian,
			JenisKelamin: AllUsers.JenisKelamin,
			Instansi:     AllUsers.Instansi,
			TempatLahir:  AllUsers.TempatLahir,
			TanggalLahir: AllUsers.TanggalLahir,
			CreteAt:      tools.TimeNowJakarta(),
		}

		database.DB.Create(&dataUsers)

	}

	return c.JSON(fiber.Map{
		"Pesan":      "Sukses Upload Data Peserta Pelatihan ",
		"data":       users,
		"Total Data": len(users),
	})
}

func ExportMateriPelatihan(c *fiber.Ctx) error {

	//Komentar Ini Aja

	//Yang akan di save itu ini

	return nil
}
