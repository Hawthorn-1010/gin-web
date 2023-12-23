package repository

import (
	"gin-web/common"
	"gin-web/model"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	db := common.GetDB()
	return CategoryRepository{DB: db}
}

func (c CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}
	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) Delete(id int) error {
	if err := c.DB.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) Query(id int) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
