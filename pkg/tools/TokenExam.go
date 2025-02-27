package tools

import (
	"NewBEUjian/app/entity"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateTokenExam(entitys interface{}) string {
	var types string
	var id_users float64 // Mengubah tipe data menjadi string
	var id_users_pelatihan float64

	switch e := entitys.(type) {
	case entity.UsersUjian:
		id_users_pelatihan = float64((int(e.IdUserUjian)))
		//id_users = float64((int(e.IdUsers)))
		// Contoh penggunaan ID, sesuaikan dengan kebutuhan And
		types = "PreTest"
	default:
		return ""
	}

	claims := jwt.MapClaims{
		"id_users_pelatihan": id_users_pelatihan,
		"id_users":           id_users,
		"type":               types,
		"exp":                time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	return t
}
func GenerateTokenExamPostTest(entitys interface{}) string {
	var types string
	var id_users float64 // Mengubah tipe data menjadi string
	var id_users_pelatihan float64

	switch e := entitys.(type) {
	case entity.UsersUjian:
		id_users_pelatihan = float64((int(e.IdUserUjian)))
		//id_users = float64((int(e.IdUsers)))
		types = "PostTest"
	default:
		return ""
	}

	claims := jwt.MapClaims{
		"id_users_pelatihan": id_users_pelatihan,
		"id_users":           id_users,
		"type":               types,
		"exp":                time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	return t
}

func GenerateTokenExamNew(entitys interface{}) string {
	var bagian string
	//var id_users float64 // Mengubah tipe data menjadi string
	var id_bagian float64
	var types_ujian string
	var id_users_ujian float64
	var NamaFungsi string
	var CodeAkses string

	//ambil data fungsi nya

	switch e := entitys.(type) {
	case entity.UjianJWT:
		id_users_ujian = float64((int(e.IdUsersUjian)))
		id_bagian = float64((int(e.IdTypeUjian)))
		bagian = e.NamaBagian
		types_ujian = e.TypeUjians
		NamaFungsi = e.NamaFungsi
		CodeAkses = e.CodeAkses
	default:
		return ""
	}

	claims := jwt.MapClaims{
		"id_users_ujian": id_users_ujian,
		"id_bagian":      id_bagian,    //Ini Id Type Ujiannya
		"bagian":         bagian,
		"typesUjian":     types_ujian,
		"namaFungsi":     NamaFungsi,
		"CodeAkses":      CodeAkses,
		"exp":            time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	return t
}
