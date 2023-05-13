package database

import (
	"github.com/zuhdiyazmi/go-tanaman/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase digunakan untuk membuat koneksi ke database mysql dan menginisialisasi objek db
func ConnectDatabase() {
	dsn := "root:waduh@tcp(localhost:3306)/go_tanaman"    // konfigurasi koneksi ke database mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // membuat koneksi database
	if err != nil {
		panic("failed to connect database")
	}

	// migrasi otomatis untuk tabel tanaman
	err = db.AutoMigrate(&models.Tanaman{}, &models.Wishlist{})
	if err != nil {
		panic("failed to migrate table")
	}

	// Assign objek DB ke variabel global
	DB = db
}
