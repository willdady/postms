package services

import (
	"github.com/willdady/postms/internal/postms/models"
)

type PostService interface {
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(post *models.Post) error
	GetPost(postID uint64) (models.Post, error)
	GetPosts(cursor string, userID string, tag string) ([]models.Post, string, error)
	CreatePostComment(postComment *models.PostComment) error
	UpdatePostComment(postComment *models.PostComment) error
	DeletePostComment(postComment *models.PostComment) error
	GetPostComment(postCommentID uint64) (models.PostComment, error)
	GetPostCommentsForPost(postID uint64) ([]models.PostComment, error)
	GetPostVoteTotalForPost(postID uint64) int64
	GetPostVote(postID uint64, userID string) (models.PostVote, error)
	GetPostVoteUsersForPost(postID uint64) []string
	CreatePostVote(postVote *models.PostVote) error
	CreatePostSave(postSave *models.PostSave) (models.PostSave, bool, error)
	GetPostSave(postSaveID uint64) (models.PostSave, error)
	GetPostSaves(postID uint64, userID string) ([]models.PostSave, error)
	DeletePostSave(postSave *models.PostSave) error
	GetTags() ([]string, error)
}
