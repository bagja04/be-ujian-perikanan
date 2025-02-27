package tools

import (
	"fmt"
	"os"
)

func CreateFolder() {
	folders := []string{
		"public/ijazah-terakhir",
		"public/foto",
		"public/sertifikat-keahlian",
		"public/sertifikat-6.09",
		"public/sertifikat-3.12",
		"public/sertifikat-6.10",
		"public/sertifikat-6.09",
		"public/sertifikat-keterampilan-lain",
		"public/sertifikat-auditor",
		"public/buku-pelaut", // sertifikat
		"public/file-permohonan",
		"public/soal-pelatihan",
		//Folder Untuk Foto Sarpras

		//Untuk bank Soal
		"public/bank-soal/ankapin",
		"public/bank-soal/atkapin",
	}

	for _, folder := range folders {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", folder, err)
			return
		}

	}

}
