package models

type Wishlist struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	NamaTanaman  string `gorm:"type:varchar(250);unique" json:"nama_tanaman"`
	JenisTanaman string `gorm:"type:varchar(250)" json:"jenis_tanaman"`
	Deskripsi    string `gorm:"type:text" json:"deskripsi"`
	Jumlah       int    `gorm:"type:int" json:"jumlah"`
}

func (Wishlist) TableName() string {
	return "wishlist"
}
