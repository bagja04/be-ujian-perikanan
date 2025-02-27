package tools

import "github.com/gofiber/fiber/v2"

func PublicFoto(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/foto/" + params)
}

func PublicIjazah(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/ijazah-terakhir/" + params)
}

func PublicSertifikatKeahlian(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/sertifikat-keahlian/" + params)
}

func PublicSertifikatTot(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/sertifikat-6.09/" + params)
}

func PublicSertifikatToe(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/sertifikat-3.12/" + params)
}

func PublicSertifikatToeSimulator(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/sertifikat-6.10/" + params)
}

func PublicSertifikatAuditor(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/sertifikat-auditor/" + params)
}

func PublicSertifikatLainnya(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/sertifikat-keterampilan-lain/" + params)
}

func PublicBukuPelaut(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/buku-pelaut/" + params)
}

func PublicFilePermohonan(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/file-permohonan/" + params)
}

func PublicFileGambarSoal(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/soal-pelatihan/" + params)
}

func PublicFileNameSoalAnkapin(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/bank-soal/ankapin/" + params)
}

func PublicFileAtkapin(c *fiber.Ctx) error {
	params := c.Params("filename")
	return c.SendFile("public/bank-soal/atkapin/" + params)
}

//Ok
