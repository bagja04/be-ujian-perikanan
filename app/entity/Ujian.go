package entity

//nama Ujian Ujiannya data ujiannya dai mengambil berdasarkan id type ujian  base nyaketika type ujian dulu yang ter generate
type Ujian struct {
	IdUjian           uint `gorm:"primary_key;auto_increment"`
	IdTypeUjian       uint
	TypeUjian         string
	NamaUjian         string //Judul Pelatihan
	TempatUjian       string
	LembagaDiklat            string
	NamaPengawasUjian string
	//Ankapin atau Atkapin
	NamaVasilitatorUjian string
	TanggalMulaiUjian    string
	TanggalBerakhirUjian string
	WaktuUjian           string    //Harga Pelatihan
	JumlahPesertaUjian   int    //
	Status               string //Aktif Atau Tidak
	CreateAt             string
	UpdateAt             string
	FilePermohonan       string
	IsSematkan string
	//Waktu ujian 
	UsersUjian []UsersUjian `gorm:"foreignKey:IdUjian;constraint:OnDelete:CASCADE"`
	// Explicit foreign key reference

}

//Yang Harus di isi yaitu adalah data masternya dulu

//Ini Isinya Ankapin / Atkapin
type TypeUjian struct {
	IdTypeUjian   uint `gorm:"primary_key;auto_increment"`
	NamaTypeUjian string  //Ini Isinya CPIB, HACCP, DLL
	CreateAt      string
	UpdateAt      string
	WaktuUjian int 
	Ujian         []Ujian  `gorm:"foreignKey:IdTypeUjian;references:IdTypeUjian"`
	Soal []SoalUjianTypeUjian `gorm:"foreignKey:IdTypeUjian;references:IdTypeUjian"`
	MateriBagian []MateriBagian `gorm:"foreignKey:IdTypeUjian;references:IdTypeUjian"`
}

//Oke Oke

//Materinya Seperti Misalkan BErapa Soal Soalnya
type MateriBagian struct {
	IdMateriBagian uint `gorm:"primaryKey;autoIncrement"`
	IdTypeUjian       uint
	NamaMateri     string
	JumlahSoal     int
	CreateAt       string
	UpdateAt       string
}
