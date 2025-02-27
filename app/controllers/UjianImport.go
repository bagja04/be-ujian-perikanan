package controllers

import (
	"NewBEUjian/app/entity"
	"NewBEUjian/pkg/database"
	"NewBEUjian/pkg/tools"

	//"NewBEUjian/pkg/tools"
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func ImportSoalNew(c *fiber.Ctx) error {
	// Membaca file Excel dari request
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Pesan": "Gagal membaca file"})
	}

	// Ambil id_type_ujian dari query
	idTypeUjian := c.Query("id_type_ujian")
	if idTypeUjian == "" {
		return c.Status(400).JSON(fiber.Map{"Pesan": "Parameter id_type_ujian diperlukan"})
	}

	// Validasi ekstensi file
	ext := filepath.Ext(file.Filename)
	if ext != ".xlsx" && ext != ".xls" {
		return c.Status(400).JSON(fiber.Map{"Pesan": "File harus berupa Excel (.xlsx atau .xls)"})
	}

	// Membuka file Excel
	excelFile, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Pesan": "Gagal membuka file Excel"})
	}
	defer excelFile.Close()

	// Membaca file Excel menggunakan excelize
	f, err := excelize.OpenReader(excelFile)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Pesan": "Gagal membaca file Excel"})
	}

	// Mendapatkan nama sheet pertama
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return c.Status(400).JSON(fiber.Map{"Pesan": "File Excel tidak memiliki sheet"})
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Pesan": "Gagal membaca baris di sheet: " + sheets[0]})
	}

	// Ambil data ujian terkait
	var dataUjian entity.TypeUjian
	if err := database.DB.Where("id_type_ujian = ?", idTypeUjian).Preload("Soal").First(&dataUjian).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"Pesan": "Data ujian tidak ditemukan"})
	}

	// Parsing Excel
	var listSoal []ResponDatas
	var soalGanda []entity.SoalUjianTypeUjian

	header := rows[0] // Baris pertama sebagai header

	for _, rowData := range rows[1:] {
		soal := ResponDatas{}
		for i, columnName := range header {
			if i >= len(rowData) {
				continue
			}

			// Mapping data dari Excel ke Struct
			switch columnName {
			case "soal":
				soal.NamaSoal = rowData[i]
			case "gambar_soal":
				soal.GambarSoal = rowData[i]
			case "jawaban_benar":
				soal.JawabanBenar = rowData[i]
			case "A":
				soal.A = rowData[i]
			case "B":
				soal.B = rowData[i]
			case "C":
				soal.c = rowData[i] // FIX: Menggunakan 'C' yang benar
			case "D":
				soal.D = rowData[i]
			case "kode_materi":
				switch rowData[i] {
				case "AT1":
					soal.Kurikulum = "Teori Dasar Dan Prinsip-prinsip Dasar Pengoperasian Mesin Kapal Penangkap Ikan"
				}
			}
		}

		// Cek apakah soal sudah ada di database
		var existingSoal entity.SoalUjianTypeUjian
		database.DB.Where("soal LIKE ? AND id_type_ujian = ?", "%"+soal.NamaSoal+"%", idTypeUjian).First(&existingSoal)
		if existingSoal.IdSoalUjianTypeUjian != 0 {
			soalGanda = append(soalGanda, existingSoal)
			fmt.Println("Soal sudah ada, lewati:", soal.NamaSoal)
			continue
		}

		listSoal = append(listSoal, soal)
	}

	// Insert data soal baru ke database
	for _, data := range listSoal {
		pertanyaan := entity.SoalUjianTypeUjian{
			IdTypeUjian:    uint(tools.StringToInt(idTypeUjian)),
			Soal:           data.NamaSoal,
			GambarSoal:     data.GambarSoal,
			Materi:         data.Kurikulum,
			Status:         "Aktif",
			CreateAt:       tools.TimeNowJakarta(),
		}

		if err := database.DB.Create(&pertanyaan).Error; err != nil {
			fmt.Println("Gagal menyimpan pertanyaan:", data.NamaSoal, "Error:", err)
			continue
		}

		// Insert jawaban
		dataJawaban := []entity.Jawaban{
			{IdSoalUjianTypeUjian: pertanyaan.IdSoalUjianTypeUjian, NameJawaban: data.JawabanBenar, CreateAt: tools.TimeNowJakarta()},
			{IdSoalUjianTypeUjian: pertanyaan.IdSoalUjianTypeUjian, NameJawaban: data.A, CreateAt: tools.TimeNowJakarta()},
			{IdSoalUjianTypeUjian: pertanyaan.IdSoalUjianTypeUjian, NameJawaban: data.B, CreateAt: tools.TimeNowJakarta()},
			{IdSoalUjianTypeUjian: pertanyaan.IdSoalUjianTypeUjian, NameJawaban: data.c,CreateAt: tools.TimeNowJakarta()},
			{IdSoalUjianTypeUjian: pertanyaan.IdSoalUjianTypeUjian, NameJawaban: data.D, CreateAt: tools.TimeNowJakarta()},
		}

		if err := database.DB.Create(&dataJawaban).Error; err != nil {
			fmt.Println("Gagal menyimpan jawaban untuk pertanyaan:", data.NamaSoal, "Error:", err)
			continue
		}

		// Update jawaban benar di SoalUjianTypeUjian
		database.DB.Model(&pertanyaan).Update("jawaban_benar", data.JawabanBenar)
	}

	return c.JSON(fiber.Map{
		"Pesan":        "Data soal berhasil diimpor",
		"Soal Ganda":   soalGanda,
		"Soal Tersimpan": len(listSoal),
	})
}
