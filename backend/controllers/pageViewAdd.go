package controllers

import (
	"fmt"
	"homers-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddController struct {
	Database *gorm.DB
}

func (addPV AddController) AddPageView(c *gin.Context) {
	var err error
	if err = addPV.Database.
		Model(&models.PageView{}).
		Where("id = ?", 1).
		Update("view_count", gorm.Expr("view_count + ?", 1)).Error; err != nil {
		c.JSON(GetErrorStatusCode(err), fmt.Sprintf("Add Page View Error: %s", err.Error()))
		return
	}
	c.JSON(200, "One Soul Collected Successfully")
}
