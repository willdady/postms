package postgres

import (
	"encoding/base64"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/willdady/postms/internal/errors"
	"github.com/willdady/postms/internal/postms/models"
	"github.com/willdady/postms/internal/utils"
)

type PostService struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{DB: db}
}

func (service *PostService) CreatePost(post *models.Post) error {
	service.DB.Create(post)
	return nil
}

func (service *PostService) UpdatePost(post *models.Post) error {
	service.DB.Save(post)
	return nil
}

func (service *PostService) DeletePost(post *models.Post) error {
	if post.ID == 0 {
		return &errors.DeleteIsMissingID{}
	}
	service.DB.Delete(post)
	return nil
}

func (service *PostService) GetPost(postID uint64) (models.Post, error) {
	p := models.Post{}
	if service.DB.Where("id = ?", postID).First(&p).RecordNotFound() {
		return p, &errors.NotFound{}
	}
	return p, nil
}

func (service *PostService) GetPosts(cursor string, userID string, tag string) ([]models.Post, string, error) {
	posts := []models.Post{}
	query := service.DB.Order("id desc")
	if cursor != "" {
		id, err := base64.StdEncoding.DecodeString(cursor)
		if err != nil {
			return posts, "", &errors.CursorDecodingError{}
		}
		query = query.Where("id <= ?", id)
	}
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if tag != "" {
		query = query.Where("? = ANY(tags)", tag)
	}
	// Note we over-fetch by 1 so we can check if there are more items
	limit := 101
	query = query.Limit(limit)
	query.Find(&posts)
	nextCursor := ""
	if len(posts) == limit {
		lastItem := posts[len(posts)-1]
		nextCursor = utils.UintToBase64(lastItem.ID)
		posts = posts[:len(posts)-1]
	}
	return posts, nextCursor, nil
}

func (service *PostService) CreatePostComment(postComment *models.PostComment) error {
	service.DB.Create(postComment)
	return nil
}

func (service *PostService) UpdatePostComment(postComment *models.PostComment) error {
	service.DB.Save(postComment)
	return nil
}

func (service *PostService) DeletePostComment(postComment *models.PostComment) error {
	if postComment.ID == 0 {
		return &errors.DeleteIsMissingID{}
	}
	service.DB.Delete(postComment)
	return nil
}

func (service *PostService) GetPostComment(postCommentID uint64) (models.PostComment, error) {
	p := models.PostComment{}
	if service.DB.Where("id = ?", postCommentID).First(&p).RecordNotFound() {
		return p, &errors.NotFound{}
	}
	return p, nil
}

func (service *PostService) GetPostCommentsForPost(postID uint64) ([]models.PostComment, error) {
	postComments := make([]models.PostComment, 0)
	service.DB.Where("post_id = ?", postID).Order("id desc").Find(&postComments)
	return postComments, nil
}

func (service *PostService) GetPostVoteTotalForPost(postID uint64) int64 {
	result := struct {
		Total int64
	}{}
	service.DB.Raw("SELECT SUM(value) as total FROM post_votes WHERE post_id = ?", postID).Scan(&result)
	return result.Total
}

func (service *PostService) GetPostVote(postID uint64, userID string) (models.PostVote, error) {
	pV := models.PostVote{}
	if service.DB.Where("post_id = ?", postID).Where("user_id = ?", userID).First(&pV).RecordNotFound() {
		return pV, &errors.NotFound{}
	}
	return pV, nil
}

func (service *PostService) GetPostVoteUsersForPost(postID uint64) []string {
	postVotes := []models.PostVote{}
	service.DB.Select("DISTINCT user_id").Where("post_id = ?", postID).Find(&postVotes)
	userIDs := make([]string, 0)
	for _, pV := range postVotes {
		userIDs = append(userIDs, pV.UserID)
	}
	return userIDs
}

func (service *PostService) CreatePostVote(postVote *models.PostVote) error {
	service.DB.Create(postVote)
	return nil
}

func (service *PostService) CreatePostSave(postSave *models.PostSave) (models.PostSave, bool, error) {
	existingPostSave := models.PostSave{}
	service.DB.Unscoped().Where("post_id = ?", postSave.PostID).Where("user_id = ?", postSave.UserID).First(&existingPostSave)
	if existingPostSave.ID > 0 {
		service.DB.Unscoped().Model(&existingPostSave).Update("deleted_at", nil)
		return existingPostSave, false, nil
	}
	service.DB.Create(postSave)
	return *postSave, true, nil
}

func (service *PostService) GetPostSave(postSaveID uint64) (models.PostSave, error) {
	results := models.PostSave{}
	if service.DB.Where("id = ?", postSaveID).First(&results).RecordNotFound() {
		return results, &errors.NotFound{}
	}
	return results, nil
}

func (service *PostService) GetPostSaves(postID uint64, userID string) ([]models.PostSave, error) {
	results := []models.PostSave{}
	var query *gorm.DB
	if postID > 0 {
		query = service.DB.Where("post_id = ?", postID)
	}
	if userID != "" {
		query = service.DB.Where("user_id = ?", userID)
	}
	query.Find(&results)
	return results, nil
}

func (service *PostService) DeletePostSave(postSave *models.PostSave) error {
	if postSave.ID == 0 {
		return &errors.DeleteIsMissingID{}
	}
	service.DB.Delete(postSave)
	return nil
}

func (service *PostService) GetTags() ([]string, error) {
	tags := pq.StringArray{}
	service.DB.Raw("SELECT array_agg(DISTINCT flattags) FROM posts, unnest(tags) as flattags").Row().Scan(&tags)
	return tags, nil
}
