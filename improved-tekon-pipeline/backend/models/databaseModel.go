package models

import "gorm.io/gorm"

type PageView struct {
	gorm.Model
	ViewCount int64
}
