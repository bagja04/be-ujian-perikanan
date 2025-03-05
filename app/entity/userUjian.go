package entity

type UsersSoal struct {
	IdUsersSoal            uint `gorm:"primaryKey;autoIncrement"`
	IdSoalUjianTypeUjian      uint
	IdCodeAksesUsers uint
	IdUserUjian            uint
	IdTypeUjian int 

	CreateAt               string
	UpdateAt               string
}

type CodeAksesUsers struct {
	IdCodeAksesUsers uint `gorm:"primary_key;auto_increment"`
		
	IdUserUjian            uint
	IdTypeUjian uint 
	NamaBagian             string
	KodeAkses              string
	CreteAt                string
	UpdateAt               string
	WaktuCodeUjian         string
	IsUse                  string
	//UsersUjian            [] UsersUjian `gorm:"foreignKey:IdUserUjian;references:IdUserUjian"`
}

type UsersUjian struct {
	IdUserUjian            uint `gorm:"primary_key;auto_increment"`
	Nama                   string
	Nik                    string
	TempatLahir            string
	TanggalLahir           string
	NomorUjian             string
	JenisKelamin           string
	Instansi               string
	IdUjian                uint
	IdCodeAksesUsersBagian uint
	Nilai              int
	NilaiPraktek int
	StatusLulus            string
	CreteAt                string
	Status                 string
	UpdateAt               string
	CodeAksesUsersBagian   CodeAksesUsers `gorm:"foreignKey:IdUserUjian;references:IdUserUjian"`
}

type UjianJWT struct {
	IdUsersUjian uint
	IdTypeUjian     uint
	NamaBagian   string
	TypeUjians   string
	NamaFungsi   string
	CodeAkses    string
}
