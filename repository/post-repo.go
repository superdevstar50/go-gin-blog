package repository

import (
	"github.com/abdurraufraihan/golang-blog-api/model"
	"gorm.io/gorm"
)

type PostRepo interface {
	Insert(post model.Post) model.Post
	AllPost() []model.Post
}

type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *postRepo {
	return &postRepo{db: db}
}

func (p postRepo) Insert(post model.Post) model.Post {
	p.db.Create(&post)
	return post
}

func (p postRepo) AllPost() []model.Post {
	var posts []model.Post
	p.db.Preload("Post").Find(&posts)
	return posts
}
