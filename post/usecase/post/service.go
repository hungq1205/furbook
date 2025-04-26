package post

import (
	"context"
	"post/entity"
	"post/infrastructure/repository/post"
	"post/util"
	"time"
)

type Service struct {
	postRepo *post.Repository
}

func NewService(postRepo *post.Repository) *Service {
	return &Service{
		postRepo: postRepo,
	}
}

// Post

func (s *Service) GetPost(ctx context.Context, id string) (*entity.Post, error) {
	post, err := s.postRepo.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *Service) GetPostsOfUser(ctx context.Context, username string, pagination util.Pagination) ([]*entity.Post, error) {
	return s.postRepo.GetPostsOfUser(ctx, username, pagination)
}

func (s *Service) GetPostsOfUsers(ctx context.Context, usernames []string, pagination util.Pagination) ([]*entity.Post, error) {
	return s.postRepo.GetPostsOfUsers(ctx, usernames, pagination)
}

func (s *Service) CheckOwnership(ctx context.Context, username, postId string) (bool, error) {
	return s.postRepo.CheckOwnership(ctx, postId, username)
}

func (s *Service) CreateBlogPost(ctx context.Context, username, content string, medias []entity.Media) (*entity.Post, error) {
	post, err := s.postRepo.CreateBlogPost(ctx, username, content, medias)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *Service) CreateLostPetPost(ctx context.Context, username, contactInfo string, postType entity.PostType, content string, medias []entity.Media, area, lastSeen *entity.Location, lostAt *time.Time) (*entity.Post, error) {
	post, err := s.postRepo.CreateLostPetPost(ctx, username, contactInfo, postType, content, medias, area, lastSeen, lostAt)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *Service) PatchContent(ctx context.Context, id, content string, medias []entity.Media) error {
	ok, err := s.postRepo.PatchContent(ctx, id, content, medias)
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	return nil
}

func (s *Service) PatchLostFoundStatus(ctx context.Context, id string, isResolved bool) error {
	return s.postRepo.PatchFound(ctx, id, isResolved)
}

func (s *Service) DeletePost(ctx context.Context, id string) error {
	return s.postRepo.DeletePost(ctx, id)
}

// Comment

func (s *Service) GetComments(ctx context.Context, postId string) ([]entity.Comment, error) {
	comments, err := s.postRepo.GetComments(ctx, postId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *Service) CreateComment(ctx context.Context, postId, username, content string) error {
	return s.postRepo.CreateComment(ctx, postId, username, content)
}

func (s *Service) DeleteComment(ctx context.Context, postId, username string) error {
	return s.postRepo.DeleteComment(ctx, postId, username)
}

// Interaction

func (s *Service) UpsertInteraction(ctx context.Context, postId, username string, itype entity.InteractionType) error {
	return s.postRepo.UpsertInteraction(ctx, postId, username, itype)
}

func (s *Service) DeleteInteraction(ctx context.Context, postId, username string) error {
	return s.postRepo.DeleteInteraction(ctx, postId, username)
}
