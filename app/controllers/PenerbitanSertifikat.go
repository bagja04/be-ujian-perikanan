package controllers

import (
	"NewBEUjian/app/entity"
	"NewBEUjian/pkg/database"
	"NewBEUjian/pkg/tools"
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)


func CretePengajuanSertifikat(c *fiber.Ctx)error{

	//Jwt Dulu 
	id_admin, _ := c.Locals("id_admin").(int)
    role, _ := c.Locals("role").(string)
    names, _ := c.Locals("name").(string)

    tools.ValidationJwtUsers(c, role, id_admin, names)

    //Ambil Infomasi berdasarkan Id Ujian, Jadi Paramter Utama di Form nya itu adalah IdUjian SAja
	id_ujian := c.Query("id_ujian")
	var requet entity.Pengajuan
	if err := c.BodyParser(&requet); err != nil {
		return c.Status(400).JSON(fiber.Map{
            "message": err.Error(),
        })
	}

    var ujian entity.Ujian

	database.DB.Where("id_ujian = ?", id_ujian).Preload("UsersUjian").First(&ujian)
	if ujian.IdTypeUjian == 0{
		return c.Status(404).JSON(fiber.Map{
            "message": "Data Ijian Tidak Ditemukan",
        })
	}

	filePermohonan, _ := c.FormFile("filePermohonan")
	// Memeriksa ekstensi file

	ext := filepath.Ext(filePermohonan.Filename)
	if ext != ".pdf" {
		return fmt.Errorf("file harus berupa pdf")
	}

    //Validasi Jumlah Lulus
	//

	//Buat Ppengajuan Baru 
	newPengajuan := entity.Pengajuan{
		IdLemdikat : ujian.IdLemdikat,
		IdUjian : ujian.IdUjian ,
		JumlahLulus : requet.JumlahLulus,
		TglPengajuan: requet.TglPengajuan,
		FilePermohonan :tools.RemoverSpaci(filePermohonan.Filename),
		Note: requet.Note,
		Status: "Pending Verifikator",
        CreateAt: tools.TimeNowJakarta(),
	}

	if err := database.DB.Create(&newPengajuan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"pesan": "Gagal menyimpan data ke database", "error": err.Error()})
	}


	if filePermohonan != nil {
		if err := c.SaveFile(filePermohonan, "public/file-permohonan/"+tools.RemoverSpaci(filePermohonan.Filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": "Failed to save Buku Pelaut", "Error": err.Error()})
		}
	}

	// Pengajuan Terlah Di Save Sekaranng ambahkan untuk users Itu nya 
	usersSertifikat:= ujian.UsersUjian
	
	var PesertaLulus []entity.UsersUjian
	//Cek dan Gunakan Komponen 60:40  untuk mengambil Orang yang Lulus untuk di ajukan 
	for _, users := range usersSertifikat{
		NilaiPraktek:= float64(users.NilaiPraktek)
		NilaiCat := float64(users.Nilai)
		TotalNilai := (NilaiPraktek * 0.60) + (NilaiCat * 0.40)
		if TotalNilai >= 65 {
            PesertaLulus = append(PesertaLulus, users)
        }
	}

	for _, PesertaLulus := range PesertaLulus{
		//Buat Untuk Create Itu 
		newUserSertifikat :=entity.SertifikatUser{
			IdPengajuan : newPengajuan.IdPengajuan,
			IdUsers : PesertaLulus.IdUserUjian,
			Nama : PesertaLulus.Nama,
			TempatLahir : PesertaLulus.TempatLahir,
			TanggalLahir: PesertaLulus.TanggalLahir,
			Nik : PesertaLulus.Nik,
			NomorSertifikat: "Masih jadi Pertimbangan",
			CreateAt: tools.TimeNowJakarta(),
		}
		if err := database.DB.Create(&newUserSertifikat).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"pesan": "Gagal menyimpan data ke database", "error": err.Error()})
        }
	}




	return c.JSON(fiber.Map{
		"message": "Berhasil membuat pengajuan sertifikat",
	})
}