package models

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ? AND deleted_on = ?", id, 0).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func GetArticles(pageNum, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ? AND deleted_on = ?", id).Update(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ? AND deleted_on = 0", id).Delete(&Article{})
	return true
}

// BeforeCreate hooks，每次创建一个article都会先更新created_on字段的值
//func (a *Article) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("created_on", time.Now().Unix())
//
//	return nil
//}

// BeforeUpdate hooks，每次更新一个article都会先更新modified_on字段的值
//func (a *Article) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("modified_on", time.Now().Unix())
//
//	return nil
//}
