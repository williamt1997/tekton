package helpers

import (
	"homers-backend/models"
)

func CreateTestPageView() error {
	pageView := models.PageView{
		ViewCount: 50,
	}
	result := TestDB.Create(&pageView)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
