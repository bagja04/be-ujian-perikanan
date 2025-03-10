package tools

import (
	//"fmt"

	"NewBEUjian/app/entity"
	//"NewBEUjian/app/models"

	//"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//Model Role model yang harus di tambahkan

func GenerateToken(entitys interface{}) string {
	var name, role, types string
	var idAdmin float64 // Mengubah tipe data menjadi string

	switch e := entitys.(type) {
	case entity.Users:
		name = e.Nama
		idAdmin = float64((int(e.IdUsers))) // Contoh penggunaan ID, sesuaikan dengan kebutuhan Anda
		role = "5"
		types = "Peserta"
	case entity.AdminPusat:
		name = e.Nama
		idAdmin = float64((int(e.IdAdminPusat))) // Contoh penggunaan ID, sesuaikan dengan kebutuhan Anda
		role = "1"
		types = "Admin Pusat"
	case entity.SuperAdmin:
		name = e.Nama
		idAdmin = float64((int(e.IdSuperAdmin))) // Konversi ID ke string
		role = "99"
		types = "SuperAdmin"

	case entity.Lemdiklat:
		name = e.NamaLemdik
		idAdmin = float64(int(e.IdLemdik))
		role = "2"
		types = "Lemdiklat"

	default:
		return ""
	}

	claims := jwt.MapClaims{
		"name":     name,
		"id_admin": idAdmin,
		"role":     role,
		"type":     types,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	return t
}
