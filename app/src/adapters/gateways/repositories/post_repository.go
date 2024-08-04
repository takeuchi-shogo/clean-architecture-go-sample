package repositories

import (
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
	"gorm.io/gorm"
)

type PostRepository struct{}

func (r *PostRepository) FindByID(db *gorm.DB, id string) (post *entities.Posts, err error) {
	post = &entities.Posts{}
	db.Where("id = ?", id).First(post)
	if post.ID == "" {
		return &entities.Posts{}, errors.New("post is not found")
	}
	return post, nil
}

func (r *PostRepository) FindByUserID(db *gorm.DB, userID string) (posts []*entities.Posts, err error) {
	posts = []*entities.Posts{}
	db.Where("user_id = ?", userID).Find(&posts)
	if len(posts) <= 0 {
		return []*entities.Posts{}, errors.New("posts is not found")
	}
	return posts, nil
}

func (r *PostRepository) Create(db *gorm.DB, post *entities.Posts) (newPost *entities.Posts, err error) {
	newPost = &entities.Posts{}

	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	newPost.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	newPost.UserID = post.UserID
	newPost.Title = post.Title
	newPost.Content = post.Content
	newPost.CreatedAt = t.Unix()
	newPost.UpdatedAt = t.Unix()

	if err = newPost.Validate(); err != nil {
		return &entities.Posts{}, err
	}

	err = db.Create(newPost).Error

	return newPost, err
}
