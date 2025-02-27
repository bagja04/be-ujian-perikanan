package entity

//Soal Ujian bagian per fungsinya

//    select id_soal_ujain_bagian, b.id_bagian, soal, jawaban_benar, jawaban_

//Berarti Pilihnya Jawaban Doubel Per Ujiannya
//Ini Langsung Ke Soalnya
type SoalUjianTypeUjian struct {
	IdSoalUjianTypeUjian uint `gorm:"primaryKey;autoIncrement"`
	IdTypeUjian uint 
	Soal              string
	GambarSoal        string
	JawabanBenar      string
	Status            string
	CreateAt          string
	UpdateAt          string
	Materi            string
	Jawaban []Jawaban `gorm:"foreignKey:IdSoalUjianTypeUjian;constraint:OnDelete:CASCADE"`
}

type Jawaban struct {
	IdJawaban         uint `gorm:"primaryKey;autoIncrement"`
	IdSoalUjianTypeUjian  uint
	NameJawaban       string
	Status            string
	CreateAt          string
	UpdateAt          string
}
