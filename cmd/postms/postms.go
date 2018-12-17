package main

import (
	"errors"
	"log"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/willdady/postms/internal/postms/handlers"
	"github.com/willdady/postms/internal/postms/models"
	"github.com/willdady/postms/internal/postms/postgres"
	"github.com/willdady/postms/internal/rest"
	"github.com/willdady/postms/internal/utils"
)

var pgHost string = utils.Getenv("PG_HOST", "0.0.0.0")
var pgPort string = utils.Getenv("PG_PORT", "5432")
var pgUser string = utils.Getenv("PG_USER", "postgres")
var pgDB string = utils.Getenv("PG_DB", "postgres")
var pgPassword string = utils.Getenv("PG_PASSWORD", "mysecretpassword")
var pgSSLMode string = utils.Getenv("PG_SSL_MODE", "disable")
var dbConnectionString = fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v", pgHost, pgPort, pgUser, pgDB, pgPassword, pgSSLMode)

func connectToDB(retry int) (db *gorm.DB, err error) {
	if retry == 5 {
		err := errors.New("Failed to connect to database after 5 tries")
		return nil, err
	}
	db, dbErr := gorm.Open("postgres", dbConnectionString)
	if dbErr != nil {
		duration := time.Second + time.Second*time.Duration(retry)
		log.Println(dbErr)
		log.Printf("Failed to connect to database. Retrying in %v seconds.\n", duration.Seconds())
		time.Sleep(duration)
		return connectToDB(retry + 1)
	}
	return db, nil
}

var resources = rest.ResourceMap{
	"posts": rest.ActionMap{
		"create":        handlers.CreatePost,
		"detail":        handlers.GetPost,
		"list":          handlers.GetPosts,
		"update":        handlers.UpdatePost,
		"delete":        handlers.DeletePost,
		"*/comments":    handlers.GetPostCommentsForPost,
		"*/total-votes": handlers.GetPostVoteTotalForPost,
		"*/voted-users": handlers.GetPostVoteUsersForPost,
		"*/saves":       handlers.GetPostSaves,
	},
	"post-votes": rest.ActionMap{
		"create": handlers.CreatePostVote,
	},
	"post-saves": rest.ActionMap{
		"create": handlers.CreatePostSave,
		"delete": handlers.DeletePostSave,
	},
	"comments": rest.ActionMap{
		"create": handlers.CreatePostComment,
		"delete": handlers.DeletePostComment,
		"update": handlers.UpdatePostComment,
		"detail": handlers.GetPostComment,
	},
}

func main() {
	db, err := connectToDB(0)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.PostComment{})
	db.AutoMigrate(&models.PostVote{})
	db.AutoMigrate(&models.PostSave{})

	postService := postgres.NewPostService(db)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("postService", postService)
		c.Next()
	})

	rest.AttachEndpoints(resources, r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
