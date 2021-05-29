package repository

import (
	"blog/infrastructure"
	"blog/models"
)

//PostRepository -> PostRepository
type PostRepository struct {
	db infrastructure.Database
}

// NewPostRepository : fetching database
func NewPostRepository(db infrastructure.Database) PostRepository {
	return PostRepository{
		db: db,
	}
}

//Save -> Method for saving post to database
func (p PostRepository) Save(post models.Post) error {
	return p.db.DB.Create(&post).Error
}

//FindAll -> Method for fetching all posts from database
func (p PostRepository) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.Post{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("post.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(post).
		Find(&posts).
		Count(&totalRows).Error
	return &posts, totalRows, err
}

//Update -> Method for updating Post
func (p PostRepository) Update(post models.Post) error {
	return p.db.DB.Save(&post).Error
}

//Find -> Method for fetching post by id
func (p PostRepository) Find(post models.Post) (models.Post, error) {
	var posts models.Post
	err := p.db.DB.
		Debug().
		Model(&models.Post{}).
		Where(&post).
		Take(&posts).Error
	return posts, err
}

//Delete Deletes Post
func (p PostRepository) Delete(post models.Post) error {
	return p.db.DB.Delete(&post).Error
}
