package controllers

import (
	"fmt"
	"homers-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetController struct {
	Database *gorm.DB
}

func (getPV GetController) GetPageView(c *gin.Context) {
	var viewAmount models.PageView
	var err error

	if err = getPV.Database.
		Model(&models.PageView{}).
		Where("id = ?", 1).
		First(&viewAmount).Error; err != nil {
		c.JSON(GetErrorStatusCode(err), fmt.Sprintf("Get Page View Error: %s", err.Error()))
		return
	}
	c.JSON(200, viewAmount.ViewCount)
}
