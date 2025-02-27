package tools

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ValidationJwt(c *fiber.Ctx, role string, id_admin int, names string) *fiber.Map {
	if role != "1" && role != "99" {

		fmt.Println("TerEksekusi")
		return &fiber.Map{
			"pesan": "Role Bukan Admin Pusat",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}

func ValidationJwtTenagaAhli(c *fiber.Ctx, role string, id_admin int, names string) *fiber.Map {
	if role != "1" && role != "99" && role != "55" {

		fmt.Println("TerEksekusi")
		return &fiber.Map{
			"pesan": "Role Bukan Tenaga Ahli",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}

func ValidationJwtLemdik(c *fiber.Ctx, role string, id_admin int, names string) *fiber.Map {
	if role != "2" && role != "99" && role != "1" {

		fmt.Println("terksekusi")
		return &fiber.Map{
			"pesan": "Role Bukan Lemdik",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}

func ValidationJwtMitra(c *fiber.Ctx, role string, id_admin int, names string) *fiber.Map {
	if role != "3" && role != "99" {
		return &fiber.Map{
			"pesan": "Role Bukan Admin Pusat",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}

func ValidationJwtUsersDewan(c *fiber.Ctx, role string, id_admin int, names string) *fiber.Map {
	if role != "1" && role != "88" && role != "99" {
		return &fiber.Map{
			"pesan": "Role Bukan Users Dewan",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}

func ValidationJwtBPPSDM(c *fiber.Ctx, role string, id_admin int, names string) *fiber.Map {
	if role != "4" && role != "99" {
		return &fiber.Map{
			"pesan": "Role Bukan BPPSDMKP",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}

func ValidationJwtUsers(c *fiber.Ctx, role string, id_admin int, names string) *fiber.Map {
	if role != "5" && role != "99" {
		return &fiber.Map{
			"pesan": "Role Bukan Admin Pusat",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}

func ValidationJwtSuperAdmin(c *fiber.Ctx, role string, id_admin int, names string) *fiber.Map {
	if role != "99" {
		return &fiber.Map{
			"pesan": "Role Bukan Admin Pusat",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}

func ValidationJwtExcam(c *fiber.Ctx, typess string, id_admin int, names string) *fiber.Map {
	if typess != "PostTest" {
		return &fiber.Map{
			"pesan": "Bukan Ujian PostTest !!!",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}
