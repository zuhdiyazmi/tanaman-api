package services

import (
	"github.com/zuhdiyazmi/go-tanaman/database"
	"github.com/zuhdiyazmi/go-tanaman/models"
)

func FindByNamaTanaman(namaTanaman string) ([]models.Tanaman, error) {
	var tanamans []models.Tanaman
	if err := database.DB.Where("nama_tanaman LIKE ?", "%"+namaTanaman+"%").Find(&tanamans).Error; err != nil {
		return nil, err
	}
	return tanamans, nil
}
