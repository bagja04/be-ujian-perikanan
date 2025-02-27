package entity

type RemedialUjian struct {
	ID          uint `gorm:"primary_key;auto_increment"`
	IdUjian     uint
	IdBagian    uint
	NamaBagian  string
	IdUserUjian uint
	CreatedAt   string
}
