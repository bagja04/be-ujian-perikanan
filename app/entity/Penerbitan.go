package entity


//Penerbitan Sertifkat
type Pengajuan  struct{
	IdPengajuan             uint `gorm:"primary_key;auto_increment"`
	IdLemdikat uint 
	IdUjian uint 
	JumlahLulus uint 
	TglPengajuan string
	
	FilePermohonan string
	Note string
	Status string
	CreateAt string
	UpdateAt string
}

type SertifikatUser struct{
	IdSertifikatUser uint `gorm:"primary_key;auto_increment"`
	IdPengajuan uint
	IdUsers uint 
	Nama string
	TempatLahir string
	TanggalLahir string
	Nik string
	NomorSertifikat string
	NomorBalnko string
	CreateAt string
	UpdateAt string

}

