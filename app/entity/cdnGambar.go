package entity

type GambarSoal struct {
	IdGambarSoal uint `gorm:"primary_key;auto_increment"`
	CodeUnik     string
	Gambar       string
	CreateAt     string
	UpdateAt     string
}

type CredentialCode struct {
	IdCredentialCode uint `gorm:"primary_key;auto_increment"`
	Code             string
	CreateAt         string
	UpdateAt         string
}
