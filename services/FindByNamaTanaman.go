package services

import (
	"github.com/zuhdiyazmi/alterra-mini-project/database"
	"github.com/zuhdiyazmi/alterra-mini-project/models"
)

func FindByNamaTanaman(namaTanaman string) ([]models.Tanaman, error) {
	var tanamans []models.Tanaman
	if err := database.DB.Where("nama_tanaman LIKE ?", "%"+namaTanaman+"%").Find(&tanamans).Error; err != nil {
		return nil, err
	}
	return tanamans, nil
}
