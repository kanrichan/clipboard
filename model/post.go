package model

type Post struct {
	ID      int64  `gorm:"column:id;primary_key" json:"id"`
	UserID  int64  `gorm:"column:user_id" json:"user_id"`
	Content string `gorm:"column:content" json:"content"`
}

func NewPost(userID int64, content string) (int64, error) {
	var post = &Post{
		UserID:  userID,
		Content: content,
	}
	err := DB.Create(&post).Error
	return post.ID, err
}

func GetPostByID(id int64) (Post, error) {
	var post Post
	return post, DB.First(&post, "id = ? ", id).Error
}

func UpdPostByID(id int64, content string) error {
	var post = Post{
		ID:      id,
		Content: content,
	}
	return DB.Model(&post).Where("id = ? ", post.ID).Update(post).Error
}

func DelPostByID(id int64) error {
	var post = &Post{ID: id}
	err := DB.Delete(&post).Error
	return err
}
