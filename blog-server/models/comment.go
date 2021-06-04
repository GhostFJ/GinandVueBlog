package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// gorm:index，用于声明这个字段为索引，如果你使用了自动迁移功能则会有所影响，在不使用则无影响
type Comment struct {
	Model

	ArticleID int `json:"article_id" gorm:"index"`
	Article Article `json:"article"`

	UserID int `json:"user_id" gorm:"index"`
	User User `json:"user"`

	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

func (comment *Comment) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())

    return nil
}

func (comment *Comment) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())

    return nil
}

func ExistCommentByID(id int) bool {
	var comment Comment
	db.Select("id").Where("id = ?", id).First(&comment)

	return comment.ID > 0
}

// 查询单个
func GetComment(id int) (comment Comment) {
	db.Where("id = ?", id).First(&comment)
	// 从查询到的数据关联， 外键
	db.Model(&comment).Related(&comment.Article)
	db.Model(&comment).Related(&comment.User)

	return
}

// 批量查询
func GetComments(pageNum int, pageSize int, maps interface {}) (comments []Comment) {
	// preload 关联查询预加载,
	/*
		Preload就是一个预加载器，它会执行两条 SQL，分别是SELECT * FROM blog_comments;
		和SELECT * FROM blog_article WHERE id IN (1,2,3,4);，
		那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询
	*/

	db.Preload("Article").Preload("User").Where(maps).Offset(pageNum).Limit(pageSize).Find(&comments)
	return
}

// 查询所有
func GetCommentTotal(maps interface {}) (count int){
    db.Model(&Comment{}).Where(maps).Count(&count)
    return
}

// 修改
func EditComment(id int, data interface{}) bool {
	db.Model(&Comment{}).Where("id = ?", id).Updates(data)

	return true
}

// 新增
func AddComment(data map[string]interface{}) bool {
	// v.(I) 断言，判断一个接口值的实际类型是否为某个类型，或一个非接口值的类型是否实现了某个接口类型

	db.Create(&Comment {
		ArticleID: data["article_id"].(int),
		UserID: data["user_id"].(int),
		Content : data["content"].(string),
        CreatedBy : data["created_by"].(string),
	})

	return true
}

// 删除
func DeleteComment(id int) bool {
	db.Where("id = ?", id).Delete(Comment{})

	return true
}
