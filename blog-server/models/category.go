package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Category struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

func (category *Category) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

/**
gorm 的Callbacks

*/
func (category *Category) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func GetCategories(pageNum int, pageSize int, maps interface{}) (categories []Category) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&categories)
	return
}

func GetCategoryTotal(maps interface{}) (count int) {
	db.Model(&Category{}).Where(maps).Count(&count)
	return
}

func ExistCategoryByName(name string) bool {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)

	return category.ID > 0 
}

func AddCategory(name string, createdBy string) bool {
	db.Create(&Category{
		Name:      name,
		CreatedBy: createdBy,
	})

	return true
}

func ExistCategoryByID(id int) bool {
	var category Category
	db.Select("id").Where("id = ?", id).First(&category)

	return category.ID > 0 
}

func EditCategory(id int, data interface{}) bool {
	db.Model(&Category{}).Where("id = ?", id).Updates(data)

	return true
}

func DelCategory(id int) bool {
	db.Where("id = ?", id).Delete(&Category{})
	return true
}
