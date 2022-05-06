package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	//嵌入自定义model
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	//Find获取所有where子句匹配的记录，记录存入tags
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	//Select指定你从表中选择的字段，First只返回符合条件的第一行
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name, createdBy string, state int) bool {
	db.Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	})

	return true
}

// BeforeCreate hooks，每次创建一个Tag都会先更新created_on字段的值
func (t *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_on", time.Now().Unix())

	return nil
}

// BeforeUpdate hooks，每次更新一个Tag都会先更新modified_on字段的值
func (t *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("modified_on", time.Now().Unix())

	return nil
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}
