package dao

import (
	"fmt"
	"gerty/internal/model"
	"log"

	"gorm.io/gorm"
)

type ContentDao struct {
	db *gorm.DB
}

func NewContentDao(db *gorm.DB) *ContentDao {
	return &ContentDao{db: db}
}

func (c *ContentDao) Create(detail model.ContentDetail) error {
	if err := c.db.Create(&detail).Error; err != nil {
		log.Printf("content create error = %v", err)
		return err
	}
	return nil
}

func (c *ContentDao) IsExists(contentID int) (bool, error) {
	var detail model.ContentDetail
	err := c.db.Where("id = ?", contentID).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		fmt.Printf("ContentDAO IsExists = [%v]", err)
		return false, err
	}
	return true, nil
}

func (c *ContentDao) Update(id int, detail model.ContentDetail) error {
	if err := c.db.Where("id = ?", id).Updates(&detail).Error; err != nil {
		log.Printf("content update error = %v", err)
	}
	return nil
}
