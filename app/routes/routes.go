package routes

import (
	"NewBEUjian/app/controllers"
	"NewBEUjian/pkg/middleware"
	"NewBEUjian/pkg/tools"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutesFiber(app *fiber.App) {

	//Engine Terbaru

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("welcome.html", fiber.Map{})
	})
	app.Get("/metrics", monitor.New())

	app.Get("/getDataKusuka", tools.GetDataKusuka)

	lemdik := app.Group("/lemdik")

	adminPusat := app.Group("/adminPusat")


	//adminPusat := app.Group("/adminpusat")

	//AdminPusat Area
	adminPusat.Post("/login", controllers.LoginAdminPusat)
	adminPusat.Get("/getAdminPusat", middleware.JwtProtect(), controllers.GetAdminPusat)

	//Type Ujian Area
	adminPusat.Post("/createTypeUjian", middleware.JwtProtect(), controllers.CreateTypeUjian)
	adminPusat.Get("/getTypeUjian", middleware.JwtProtect(), controllers.GetTypeUjian)
	adminPusat.Delete("/deleteUjian", middleware.JwtProtect(), controllers.DeleteTypeUjian)

	//Fungsi Ujian

	//User Post Add Pelatihan

	//Users Add Ujikom

	//lemdik Area
	lemdik.Post("/login", controllers.LoginLemdik)
	lemdik.Get("/getLemdik", middleware.JwtProtect(), controllers.GetLemdik)
	lemdik.Put("/update", middleware.JwtProtect(), controllers.UpdateLemdik)

	//Uji Kom

	//Ujian
	adminPusat.Post("/createUjian", controllers.CreateUjian)
	adminPusat.Put("/updateUjian", middleware.JwtProtect(), controllers.UpdateUjian)
	adminPusat.Get("/GetUjian", middleware.JwtProtect(), controllers.GetUjian)
	adminPusat.Delete("/deleteUjians", middleware.JwtProtect(), controllers.DeleteUjians)
	//Soal Post Test dan Pretase
	adminPusat.Post("/AddSoalToUsers", middleware.JwtProtect(), controllers.AddSoalToUsersNew)                         //AddSoalToUsersNewBagianSuffle
 //AddSoalToUsersNewBagianSuffle

	//lemdik.Get("/GetSoalPelatihanById", middleware.JwtProtect(), controllers.GetSoalPelatihanByLemdik)
	//lemdik.Get("/GetPertanyaanRandom", middleware.JwtProtect(), controllers.GetPertanyaanRandom)
	//lemdik.Post("/AddSoalUsers", middleware.JwtProtect(), controllers.AddSoalToUsers)
	app.Post("/AuthExam", controllers.AuthExam)
	app.Get("/getSoalTypeUjian", middleware.JwtExamProtectNes(),controllers.GetSoalUsers)


	app.Post("/SumbitExam", middleware.JwtExamProtectNes(), controllers.Jawab)

	lemdik.Post("/LastNomorSertifBalai", middleware.JwtProtect(), controllers.LastNomorSertifBalai)

	//lemdik.Put("/updateLastSertif", middleware.JwtProtect(), controllers.AddLastSertifLowBalai)

	//Sarpras

	//Pelatihan Users Area

	//super admin
	//Create User area
	SuperAdmin := app.Group("/superadmin")

	SuperAdmin.Post("/registerAdminPusat", middleware.JwtProtect(), controllers.CreateAdminPusat)
	SuperAdmin.Post("/regiterLemdik", middleware.JwtProtect(), controllers.RegisterLemdik)
	SuperAdmin.Post("/login", controllers.SuperAdminLogin)

	app.Post("/importUjian", controllers.ImportSoalNew)   //Done


	//static file

	app.Get("/public/ijazah-terakhir/:filename", tools.PublicIjazah)
	app.Get("/public/sertifikat-keahlian/:filename", tools.PublicSertifikatKeahlian)
	app.Get("/public/sertifikat-6.09/:filename", tools.PublicSertifikatTot)
	app.Get("/public/sertifikat-3.12/:filename", tools.PublicSertifikatToe)
	app.Get("/public/sertifikat-6.10/:filename", tools.PublicSertifikatToeSimulator)
	app.Get("/public/sertifikat-auditor/:filename", tools.PublicSertifikatAuditor)
	app.Get("/public/sertifikat-keterampilan-lain/:filename", tools.PublicSertifikatLainnya)
	app.Get("/public/buku-pelaut/:filename", tools.PublicBukuPelaut)
	app.Get("/public/foto/:filename", tools.PublicFoto)
	app.Get("/public/file-permohonan/:filename", tools.PublicFilePermohonan)
	app.Get("/public/soal-pelatihan/:filename", tools.PublicFileGambarSoal)
	app.Get("/public/bank-soal/atkapin/:filename", tools.PublicFileAtkapin)
	app.Get("/public/bank-soal/ankapin/:filename", tools.PublicFileNameSoalAnkapin)

	//Form DPAKP
	app.Post("/exportPesertaPelatihan", middleware.JwtProtect(), controllers.ExportPesertaPelatihan)  // done
	app.Post("/uploadGambar", controllers.UploadGambarSoal)
	app.Get("/getGambar", controllers.GetGambar)

	adminPusat.Put("/updateSoal", middleware.JwtProtect(), controllers.UpdateSoal)
	adminPusat.Delete("/deleteSoal", middleware.JwtProtect(), controllers.DeleteSoal)

	//app.Post("/exportPesertaUjikom", controllers.ExportPesertaUjikom)

	//Tenga Ahli

	//Materi Bagian

	adminPusat.Get("/getMateriBagian", middleware.JwtProtect(), controllers.GetMateriBagianByIdBagian)

	//Yang ganda

	app.Get("/getInfoUsers", middleware.JwtExamProtectNes(), controllers.GetUserUjianByCodeAkses)


	//Matreri BAgia 
	adminPusat.Post("/createMateriBagian", middleware.JwtProtect(),controllers.CreateMateriBagian)
	adminPusat.Put("/updateMateriBagian", middleware.JwtProtect(), controllers.UpdateMateriBagian)
	adminPusat.Delete("/deleteMateriBagian", middleware.JwtProtect(), controllers.DeleteMateriBagian)

}
