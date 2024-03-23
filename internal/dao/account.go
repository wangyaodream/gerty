package dao

import (
	"fmt"
	"gerty/internal/model"

	"gorm.io/gorm"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{db: db}
}

func (a *AccountDao) IsExists(userID string) (bool, error) {
	var account model.Account
	err := a.db.Where("user_id = ?", userID).First(&account).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		fmt.Printf("AccountDao IsExists = [%v]", err)
		return false, err
	}
	return true, nil
}

func (a *AccountDao) Create(account model.Account) error {
	if err := a.db.Create(&account).Error; err != nil {
		fmt.Printf("Account create error: %v", err.Error())
		return err
	}
	return nil
}

func (a *AccountDao) FirstByUserID(userID string) (*model.Account, error) {
	var account model.Account
	err := a.db.Where("user_id = ?", userID).First(&account).Error
	if err != nil {
		fmt.Printf("FirstByUserID error = %v", err)
		return nil, err
	}
	return &account, nil
}
