package entity

// Users Akan di import datanya by excel
type Users struct {
	IdUsers      uint `gorm:"primary_key;auto_increment"`
	Nama         string
	NoTelpon     int
	Email        string
	Nik          int
	JenisKelamin string
	TempatLahir  string
	TanggalLahir string
	NomorUjian   string
	Instansi     string
	Status       string
	CreateAt     string
	UpdateAt     string
	//Pelatihan    []UsersPelatihan `gorm:"foreignKey:IdUsers"`
}

// Auto generate
type SuperAdmin struct {
	IdSuperAdmin uint `gorm:"primary_key;auto_increment"`
	Nama         string
	Email        string
	Username     string
	Password     string
	Status       string
}

// Entity Untuk Lemdiklat
// Auto Generate or Input From admin pusat
type AdminPusat struct {
	IdAdminPusat uint `gorm:"primary_key;auto_increment"`
	Nama         string
	Email        string
	Password     string
	NoTelpon     string
	Nip          string
	Status       string
	Role string
}
