package entity

//Lemdik disini adalah si Balai Pelatihan, Politeknik, Sekolah Smk

//Operator Penerbitan Ujian yang di Generate oleh
type Lemdiklat struct {
	IdLemdik     uint `gorm:"primary_key;auto_increment"`
	NamaLemdik   string
	NoTelpon     int
	Email        string
	Password     string
	Alamat       string
	Deskripsi    string
	CreateAt     string
	UpdateAt     string
}
