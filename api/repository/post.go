package repository

import (
	"blog/infrastructure"
	"blog/models"
)

type PostRepository struct {
	db infrastructure.Database
}

func NewRepository(db infrastructure.Database) PostRepository {
	return PostRepository{
		db: db,
	}
}
func (p PostRepository) Save(post models.Post) error {
	return p.db.DB.Create(&post).Error
}

func (p PostRepository) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0
	queryBuilder := p.db.DB.Order("created_at desc").Model(&models.Post{})
	if keyword != "" {

		querykeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			p.db.DB.Where("post.title LIKE ? ", querykeyword))
	}
	err := queryBuilder.Where(post).Find(&posts).Count(&totalRows).Error
	return &posts, totalRows, err
}
func (p PostRepository) Update(post models.Post) error {
	return p.db.DB.Save(&post).Error
}
func (p PostRepository) Find(post models.Post) (models.Post, error) {
	var posts models.Post
	err := p.db.DB.Debug().
		Model(&models.Post{}).Where(&post).Take(&posts).Error
	return posts, err
}
func (p PostRepository) Delete(post models.Post) error {
	return p.db.DB.Delete(&post).Error
}
