package controllers

import (
	"NewBEUjian/app/entity"
	"NewBEUjian/pkg/database"
	"NewBEUjian/pkg/tools"
	"fmt"

	"strings"
	"time"


	"github.com/gofiber/fiber/v2"
)

type ResponDatas struct {
	IdTypeUjian string

	NamaSoal     string
	GambarSoal   string
	IdBagian     string
	Paket        int
	JawabanBenar string
	A            string
	B            string
	c            string
	D            string
	Kurikulum    string
}

func CdnGambarSoal(c *fiber.Ctx) error {

	//"public/foto/soal-pelatihan",

	//Sebagai Password

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Menyimpan Gambar",
		"data":  "",
	})
}

func ChekUjianAkses(c *fiber.Ctx, waktuUjian string) error {
	waktuSekarang := time.Now()

	// Format waktu yang diharapkan
	format := "2006-01-02 15:04:05 MST"

	// Parsing waktu ujian dari string ke time.Time
	parsedWaktuUjian, err := time.Parse(format, waktuUjian)
	if err != nil {
		fmt.Println("Error parsing waktu ujian:", err)
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "Format waktu ujian tidak valid. Gunakan format: " + format,
		})
	}

	// Cek apakah waktu ujian sudah lewat
	if waktuSekarang.After(parsedWaktuUjian) {
		return c.Status(403).JSON(fiber.Map{
			"Pesan": "Waktu ujian sudah lewat, akses tidak bisa dilanjutkan.",
		})
	}

	// Perhitungan selisih waktu dalam jam
	selisihWaktu := parsedWaktuUjian.Sub(waktuSekarang).Hours()

	// Cek apakah waktu ujian kurang dari 2 jam dari sekarang
	if selisihWaktu > 2 {
		return c.Status(403).JSON(fiber.Map{
			"Pesan": "Waktu ujian masih lebih dari 2 jam dari sekarang, akses ditolak.",
		})
	}

	// Jika semua cek lolos, lanjutkan
	return nil
}

func convertToWIB(input string) (string, error) {
	// Format waktu standar
	format := "2006-01-02 15:04:05"

	// Definisikan lokasi untuk WIB
	locWIB, _ := time.LoadLocation("Asia/Jakarta")

	// Ganti nama zona waktu WITA dan WIT dengan GMT offset
	if strings.Contains(input, "WITA") {
		input = strings.Replace(input, "WITA", "+08:00", 1)
	} else if strings.Contains(input, "WIT") {
		input = strings.Replace(input, "WIT", "+09:00", 1)
	} else if strings.Contains(input, "WIB") {
		input = strings.Replace(input, "WIB", "+07:00", 1)
	}

	// Parsing waktu dengan GMT offset
	parsedTime, err := time.Parse(format+" -07:00", input)
	if err != nil {
		return "", fmt.Errorf("error parsing time: %v", err)
	}

	// Konversikan ke WIB
	wibTime := parsedTime.In(locWIB)

	// Kembalikan waktu dalam format WIB
	return wibTime.Format(format), nil
}

func ambilZonaWaktu(waktu string) string {
	// Split string berdasarkan spasi
	parts := strings.Split(waktu, " ")

	// Periksa apakah elemen terakhir adalah zona waktu
	if len(parts) > 0 {
		return parts[len(parts)-1] // Mengambil elemen terakhir (zona waktu)
	}
	return ""
}

// Semoga berhasil
func AuthExam(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	codeAkses := data["code_akses"]

	//type_exam := data["type_exam"]
	if codeAkses == "" {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "Mohon Maaf Masukan Kode Akses Anda",
		})
	}

	var codeBagianUsers entity.CodeAksesUsers
	database.DB.Where("kode_akses = ? AND is_use = ?", codeAkses, "false").Find(&codeBagianUsers)
	if codeBagianUsers.KodeAkses == "" {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "Anda Tidak Terdaftar",
		})
	}

	//cek Waktu Kode Ujian dengan Waktu Mulai nya
	waktuUjianDiCode := codeBagianUsers.WaktuCodeUjian // Pastikan ini bertipe time.Time
	zonaWaktu := ambilZonaWaktu(waktuUjianDiCode)

	var waktuSekarang time.Time
	var waktuFormatted string

	switch zonaWaktu {
	case "WIT":
		waktuLokasi, _ := time.LoadLocation("Asia/Jayapura")
		// Ambil waktu sekarang dalam WIT
		waktuSekarang = time.Now().In(waktuLokasi)
		waktuFormatted = waktuSekarang.Format("2006-01-02 15:04:05")

	case "WITA":
		waktuLokasi, _ := time.LoadLocation("Asia/Makassar")
		// Ambil waktu sekarang dalam WITA
		waktuSekarang = time.Now().In(waktuLokasi)
		waktuFormatted = waktuSekarang.Format("2006-01-02 15:04:05")

	case "WIB":
		waktuLokasi, _ := time.LoadLocation("Asia/Jakarta")
		// Ambil waktu sekarang dalam WIB
		waktuSekarang = time.Now().In(waktuLokasi)
		waktuFormatted = waktuSekarang.Format("2006-01-02 15:04:05")

	default:
		fmt.Println("Zona waktu tidak valid:", zonaWaktu)
	}

	// Format waktu yang diharapkan
	format := "2006-01-02 15:04:05 -0700 MST"

	fmt.Println("Waktu Sekarang ", waktuSekarang)
	fmt.Println("Waktu Formated ", waktuFormatted)

	// Parsing waktu ujian dari string ke time.Time
	parsedWaktuUjian, err := time.Parse(format, waktuUjianDiCode)
	if err != nil {
		fmt.Println("Error parsing waktu ujian:", err)
		return c.Status(400).JSON(fiber.Map{
			"Pesan": err,
		})
	}

	fmt.Println("Waktu Parse Ujian", parsedWaktuUjian)

	// Cek apakah waktu ujian sudah lewat
	if waktuSekarang.Before(parsedWaktuUjian) {
		return c.Status(403).JSON(fiber.Map{
			"Pesan": "Maaf Waktu Ujian Belum Di Mulai",
		})
	}

	// Perhitungan selisih waktu dalam jam
	selisihWaktu := parsedWaktuUjian.Sub(waktuSekarang).Hours()

	fmt.Println("Beda waktu", selisihWaktu)

	// Cek apakah waktu ujian kurang dari 2 jam dari sekarang Ini Aftar  (Afternya)
	if selisihWaktu < -2 {
		return c.Status(403).JSON(fiber.Map{
			"Pesan": "Waktu Ujian Telah Habis Untuk Fungsi Ini",
		})
	}
	// Melanjutkan jika tidak ada error
	idUsersUjians := codeBagianUsers.IdUserUjian
	IdTypeUjian := codeBagianUsers.IdTypeUjian


	// Ambil data ujian
	var ujian entity.TypeUjian
	if result := database.DB.Where("id_type_ujian = ?", IdTypeUjian).First(&ujian); result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"Pesan": "Data ujian tidak ditemukan.",
		})
	}

	// Buat token ujian
	ujianJwt := entity.UjianJWT{
		IdUsersUjian: idUsersUjians,
		IdTypeUjian:     IdTypeUjian,
		NamaBagian: ujian.NamaTypeUjian,
		TypeUjians:   ujian.NamaTypeUjian,
		CodeAkses:    codeAkses,
	}

	token := tools.GenerateTokenExamNew(ujianJwt)

	return c.JSON(fiber.Map{
		"t": token,
	})

}

//Oke

//Fungsi Unuk mengambil soal dari uersSoal Pelatihan

func GetSoalUsers(c *fiber.Ctx) error {

	//id_users, _ := c.Locals("id_users").(int)
	id_user_ujian, _ := c.Locals("id_users_ujian").(int)  //Users Ujiannnya 
	types, _ := c.Locals("bagian").(string)

	idTypeUjian, _ := c.Locals("id_bagian").(int) //id Type Ujian
	typesUjian, _ := c.Locals("typesUjian").(string)
	namaFungsi, _ := c.Locals("namaFungsi").(string)

	//ambil soal dari table users soal berdasarkan id users pelatihan
	var soalUsers []entity.UsersSoal 
	var typeUjian entity.TypeUjian
	database.DB.Where("id_type_ujian =?", idTypeUjian).Find(&typeUjian)

	// Id users pelatihannya dan juga id_TYPE_UJIANNYA
	database.DB.Where("id_user_ujian = ? AND id_type_Ujian = ?", id_user_ujian, idTypeUjian).Find(&soalUsers)

	idSoalUjian := []int64{}

	for _, id_soal := range soalUsers {
		idSoalUjian = append(idSoalUjian, int64(id_soal.IdSoalUjianTypeUjian))
	}

	//Mengambil soal dan jawaban dari users

	var soal []entity.SoalUjianTypeUjian

	database.DB.Where("id_soal_ujian_type_ujian  IN ?", idSoalUjian).Order("id_soal_ujian_type_ujian ASC").Preload("Jawaban").Find(&soal)

	//ambil soal berdasarkan

	return c.JSON(fiber.Map{
		"Ujian":  typesUjian,
		"waktu":  typeUjian.WaktuUjian,
		"Fungsi": namaFungsi,
		"Bagian": types,
		"Soal":   soal,
		"jumlah": len(idSoalUjian),
	})

}



func Jawab(c *fiber.Ctx) error {

	//id_users, _ := c.Locals("id_users").(int)
	id_users_pelatihan, _ := c.Locals("id_users_ujian").(int)
	NamaBagian, _ := c.Locals("bagian").(string)
	IdBagian, _ := c.Locals("id_bagian").(int)
	typesUjian, _ := c.Locals("typesUjian").(string)

	kodeAkses, _ := c.Locals("CodeAkses").(string)

	if id_users_pelatihan == 0 {

	}

	if NamaBagian == "" {

	}

	var jawaban []struct {
		IdSoalBagian    string `json:"id_soal_bagian"`
		JawabanPengguna string `json:"jawaban_pengguna"`
	}

	// Parse request body
	if err := c.BodyParser(&jawaban); err != nil {
		fmt.Println("Error parsing body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal memparsing body",
		})
	}

	ScoreBenar := 0
	JawabanTerjawab := 0
	//var jawabanBenarCount int

	fmt.Println("Jawaban Masuk Total ada:", len(jawaban))
	for _, jwb := range jawaban {

		if jwb.JawabanPengguna == "" {
			continue
		}

		JawabanTerjawab++

		var soal entity.SoalUjianTypeUjian

		result := database.DB.Where("id_soal_ujian_type_ujian  = ?", jwb.IdSoalBagian).Find(&soal)

		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal mendapatkan soal",
			})
		}

		fmt.Println("Jawaban :", jwb.JawabanPengguna, "Jawaban Benar :", soal.JawabanBenar)
		
		if jwb.JawabanPengguna == soal.JawabanBenar {
			ScoreBenar++
			fmt.Println("Nambah Scoree")
		}
	}

	//Perhitungan Score Berdasarkan jumlan soal dan soal benar
	if JawabanTerjawab == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak ada jawaban yang valid",
		})
	}

	fmt.Println("Score Benar adalah:", ScoreBenar)
	fmt.Println("Jawaban yang di jawab", JawabanTerjawab)

	finalScore := (float64(ScoreBenar) / float64(len(jawaban))) * 100.0

	fmt.Println("Final Score :", finalScore)

	//Kirim data ke API E-Laut untuk Peserta Ujjian
	var usersPelatihan entity.UsersUjian
	database.DB.Where("id_user_ujian = ?", id_users_pelatihan).Find(&usersPelatihan)

	//Pengambilan users Agus
	fmt.Println(usersPelatihan)
	fmt.Println(typesUjian)
	fmt.Println(NamaBagian)
	fmt.Println(IdBagian)
	//Ubah nilainya ter gantung dia pre test dan post test
	
	usersPelatihan.Nilai = int(finalScore)

	database.DB.Model(&usersPelatihan).Updates(&usersPelatihan)

	var codeAkses entity.CodeAksesUsers
	database.DB.Where("kode_akses = ?", kodeAkses).Find(&codeAkses)

	codeAkses.IsUse = "true"
	database.DB.Model(&codeAkses).Updates(&codeAkses)

	//Update Code Akses Ujian

	//terbagu Ok

	//Jika telah mengerjakan Post test Hapus Soal yang ada di userSoal Agar data tidak numpuk

	return c.JSON(fiber.Map{
		"Pesan": "Terima Kasih Telah Megerjakan Ujian",
	})
}


