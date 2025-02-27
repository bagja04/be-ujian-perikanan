package main

import (
	"NewBEUjian/app/routes"
	"NewBEUjian/pkg/config"
	"NewBEUjian/pkg/database"
	"NewBEUjian/pkg/tools"

	//"backend-elaut/pkg/config"
	"log"
	"os"
	//"gorm.io/gorm"
	//"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2"
)

func main() {
	tools.CreateFolder()
	viperConfig := config.NewViper()

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Gagal membuka file log:", err)
	}
	defer file.Close()

	database.Connect()
	//Bagja
	// Set output log ke file yang telah dibuka
	log.SetOutput(file)

	app := config.NewFiber(viperConfig)

	routes.SetupRoutesFiber(app)

	log.Fatal(app.Listen(config.NewViper().GetString("web.port")))
}
