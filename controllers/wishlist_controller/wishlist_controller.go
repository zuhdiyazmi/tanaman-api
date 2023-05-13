package wishlist_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zuhdiyazmi/go-tanaman/database"
	"github.com/zuhdiyazmi/go-tanaman/models"
	"gorm.io/gorm"
)

// Index menampilkan semua wishlist tanaman
func Index(c *gin.Context) {
	var wishlists []models.Wishlist
	if err := database.DB.Find(&wishlists).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wishlists,
	})
}

// Show menampilkan satu wishlist tanaman
func Show(c *gin.Context) {
	var wishlist models.Wishlist
	if err := database.DB.Where("id = ?", c.Param("id")).First(&wishlist).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "wishlist not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wishlist,
	})
}

// Create membuat wishlist tanaman baru
func Create(c *gin.Context) {
	var input models.Wishlist
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// memeriksa apakah wishlist dengan nama yang sama sudah ada
	if wishlistExists(input.NamaTanaman) {
		c.JSON(http.StatusConflict, gin.H{"error": "wishlist dengan nama yang sama sudah ada"})
		return
	}

	tx := database.DB.Begin()

	if err := tx.Create(&input).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func wishlistExists(name string) bool {
	var wishlist models.Wishlist
	if err := database.DB.Where("nama_tanaman = ?", name).First(&wishlist).Error; err != nil {
		return false
	}
	return true
}

// Update mengupdate wishlist yang sudah ada
func Update(c *gin.Context) {
	var wishlist models.Wishlist
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id wishlist tidak valid",
		})
		return
	}
	if err := database.DB.First(&wishlist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "wishlist tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.ShouldBindJSON(&wishlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := database.DB.Save(&wishlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": wishlist,
	})
}

// Delete menghapus wishlist yang sudah ada
func Delete(c *gin.Context) {
	var wishlist models.Wishlist
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id wishlist tidak valid",
		})
		return
	}
	if err := database.DB.First(&wishlist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "wishlist tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := database.DB.Delete(&wishlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "wishlist berhasil dihapus",
	})
}

// Search mencari produk berdasarkan kata kunci tertentu
func Search(c *gin.Context) {
	db := database.DB

	// Validasi input query string
	namaTanaman := c.Query("nama_tanaman")
	if namaTanaman == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama tanaman harus diisi",
		})
		return
	}

	// Lakukan pencarian
	var wishlist []models.Wishlist
	if err := db.Where("nama_tanaman LIKE ?", "%"+namaTanaman+"%").Find(&wishlist).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Tampilkan hasil pencarian atau pesan error
	if len(wishlist) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Pencarian tidak ditemukan",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"wishlist": wishlist,
		})
	}
}
