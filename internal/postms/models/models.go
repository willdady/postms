package models

import (
	"time"

	"github.com/gosimple/slug"
	"github.com/lib/pq"
	"github.com/willdady/postms/internal/utils"
)

type CommonFields struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}

type Post struct {
	CommonFields
	UserID string         `json:"userId" binding:"required"`
	Title  string         `json:"title" binding:"required"`
	Slug   string         `json:"slug"`
	Body   string         `json:"body" binding:"required"`
	Tags   pq.StringArray `json:"tags" gorm:"type:varchar(64)[]"`
}

func (p *Post) BeforeCreate() (err error) {
	p.Slug = slug.Make(p.Title)
	p.Tags = utils.ToTagSlice(p.Tags)
	return
}

func (p *Post) BeforeSave() (err error) {
	p.Slug = slug.Make(p.Title)
	p.Tags = utils.ToTagSlice(p.Tags)
	return
}

func (p *Post) BeforeUpdate() (err error) {
	p.Slug = slug.Make(p.Title)
	p.Tags = utils.ToTagSlice(p.Tags)
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
