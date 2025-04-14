package post

import (
	"context"
	"post/entity"
	"post/util"
	"time"
)

type UseCase interface {
	GetPost(ctx context.Context, id string) (*entity.Post, error)
	GetPostsOfUser(ctx context.Context, username string, pagination util.Pagination) ([]*entity.Post, error)
	GetPostsOfUsers(ctx context.Context, usernames []string, pagination util.Pagination) ([]*entity.Post, error)
	CheckOwnership(ctx context.Context, username, postId string) (bool, error)
	CreateBlogPost(ctx context.Context, username, content string, medias []entity.Media) (*entity.Post, error)
	CreateLostPetPost(ctx context.Context, username, content string, petId int, medias []entity.Media, area, lastSeen *entity.Location, lostAt *time.Time) (*entity.Post, error)
	PatchContent(ctx context.Context, id, content string, medias []entity.Media) (*entity.Post, error)
	PatchFound(ctx context.Context, id string, found bool) error
	DeletePost(ctx context.Context, id string) error

	CreateComment(ctx context.Context, postId, username, content string) error
	DeleteComment(ctx context.Context, postId, username string) error
	GetComments(ctx context.Context, postId string) ([]entity.Comment, error)

	UpsertInteraction(ctx context.Context, postId, username string, itype entity.InteractionType) error
	DeleteInteraction(ctx context.Context, postId, username string, itype entity.InteractionType) error
}
