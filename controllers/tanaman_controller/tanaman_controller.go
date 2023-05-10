package tanaman_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zuhdiyazmi/alterra-mini-project/database"
	"github.com/zuhdiyazmi/alterra-mini-project/models"
	"github.com/zuhdiyazmi/alterra-mini-project/services"
	"gorm.io/gorm"
)

// index: mengambil semua tanaman
func Index(c *gin.Context) {
	var tanamans []models.Tanaman
	if err := database.DB.Find(&tanamans).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tanamans,
	})
}

// show: mengambil tanaman spesifik dengan id
func Show(c *gin.Context) {
	var tanaman models.Tanaman
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id tanaman tidak valid",
		})
		return
	}

	if err := database.DB.First(&tanaman, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "tanaman tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "terjadi kesalahan saat mengambil data tanaman",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tanaman,
	})
}

// create: membuat produk baru
func Create(c *gin.Context) {
	var tanaman models.Tanaman

	if err := c.BindJSON(&tanaman); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// memeriksa apakah produk dengan nama yang sama sudah ada
	if tanamanExists(tanaman.NamaTanaman) {
		c.JSON(http.StatusConflict, gin.H{"error": "tanaman dengan nama yang sama sudah ada"})
		return
	}

	tx := database.DB.Begin()

	if err := tx.Create(&tanaman).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"data": tanaman})
}

func tanamanExists(name string) bool {
	var tanaman models.Tanaman
	if err := database.DB.Where("nama_tanaman = ?", name).First(&tanaman).Error; err != nil {
		return false
	}
	return true
}

// update: memperbarui tanaman yang sudah ada dengan id
func Update(c *gin.Context) {
	var tanaman models.Tanaman
	// parsing id dari parameter url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id tanaman tidak valid",
		})
		return
	}
	// mencari tanaman dengan id yang sesuai di database
	if err := database.DB.First(&tanaman, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "tanaman tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// memperbarui atribut dari tanaman berdasarkan json request body
	if err := c.ShouldBindJSON(&tanaman); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// menyimpan perubahan ke database
	if err := database.DB.Save(&tanaman).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// mengirimkan response dengan data tanaman yang barusan diperbarui
	c.JSON(http.StatusOK, gin.H{
		"data": tanaman,
	})
}

func Delete(c *gin.Context) {
	// mengambil id dari URL endpoint
	id := c.Param("id")

	// melakukan konversi tipe data dari string ke int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id tanaman tidak valid"})
		return
	}

	// membuat objek tanaman untuk dihapus
	tanaman := models.Tanaman{Id: idInt}

	// melakukan penghapusan data dari database
	result := database.DB.Delete(&tanaman)
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
		return
	}

	if err := result.Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}

// search: mengambil semua tanaman berdasarkan nama tanaman
func Search(c *gin.Context) {
	namaTanaman := c.Query("nama_tanaman")

	// memanggil service untuk mencari tanaman berdasarkan nama tanaman
	tanamans, err := services.FindByNamaTanaman(namaTanaman)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Terjadi kesalahan saat mengambil data tanaman",
		})
		return
	}

	if len(tanamans) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tanaman tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tanamans": tanamans,
	})
}
