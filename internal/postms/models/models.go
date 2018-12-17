package models

import (
	"time"

	"github.com/gosimple/slug"
)

type CommonFields struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}

type Post struct {
	CommonFields
	UserID string `json:"userId" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Slug   string `json:"slug"`
	Body   string `json:"body" binding:"required"`
}

func (p *Post) BeforeCreate() (err error) {
	p.Slug = slug.Make(p.Title)
	return
}

func (p *Post) BeforeSave() (err error) {
	p.Slug = slug.Make(p.Title)
	return
}

func (p *Post) BeforeUpdate() (err error) {
	p.Slug = slug.Make(p.Title)
	return
}

type PostComment struct {
	CommonFields
	UserID string `json:"userId" binding:"required"`
	PostID uint   `json:"postId" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

type PostVote struct {
	CreatedAt time.Time `json:"createdAt"`
	UserID    string    `json:"userId" binding:"required" gorm:"primary_key`
	PostID    uint      `json:"postId" binding:"required" gorm:"primary_key`
	Value     int       `json:"value" binding:"required"`
}

type PostSave struct {
	CommonFields
	UserID string `json:"userId" binding:"required" gorm:"primary_key`
	PostID uint   `json:"postId" binding:"required" gorm:"primary_key`
}
