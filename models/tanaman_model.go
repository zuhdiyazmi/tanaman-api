package models

type Tanaman struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	NamaTanaman  string `gorm:"type:varchar(250);unique" json:"nama_tanaman"`
	JenisTanaman string `gorm:"type:varchar(250)" json:"jenis_tanaman"`
	Deskripsi    string `gorm:"type:text" json:"deskripsi"`
	Lokasi       string `gorm:"type:varchar(250)" json:"lokasi"`
	Catatan      string `gorm:"type:text" json:"catatan"`
	Jumlah       int    `gorm:"type:int" json:"jumlah"`
}

func (Tanaman) TableName() string {
	return "tanaman"
}
