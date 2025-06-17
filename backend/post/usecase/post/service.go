package post

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"post/api/client"
	"post/entity"
	"post/infrastructure/repository/post"
	"post/util"
	"time"
)

type Service struct {
	postRepo    *post.Repository
	notiClient  client.NotiClient
	userService client.UserClient
}

func NewService(postRepo *post.Repository, notiClient client.NotiClient) *Service {
	return &Service{
		postRepo:   postRepo,
		notiClient: notiClient,
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

func (s *Service) GetNearLostPosts(ctx context.Context, latitude float64, longitude float64, pagination util.Pagination) ([]*entity.Post, error) {
	return s.postRepo.GetNearLostPosts(ctx, latitude, longitude, pagination)
}

func (s *Service) GetParticipatedPostsOfUser(ctx context.Context, username string, pagination util.Pagination) ([]*entity.Post, error) {
	return s.postRepo.GetParticipatedPostsOfUser(ctx, username, pagination)
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
	if postType != entity.Found && area != nil && len(area.Location.Coordinates) > 0 {
		address, err := fetchAddress(area.Location.Coordinates[1], area.Location.Coordinates[0])
		if err != nil {
			return nil, fmt.Errorf("failed to fetch address for area: %w", err)
		}
		area.Address = address
	}

	if lastSeen != nil && len(lastSeen.Location.Coordinates) > 0 {
		address, err := fetchAddress(lastSeen.Location.Coordinates[1], lastSeen.Location.Coordinates[0])
		if err != nil {
			return nil, fmt.Errorf("failed to fetch address for lastSeen: %w", err)
		}
		lastSeen.Address = address
	}

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
	if err := s.postRepo.PatchFound(ctx, id, isResolved); err != nil {
		return err
	}
	post, err := s.postRepo.GetPost(ctx, id)
	if err != nil {
		return err
	}
	return s.notiClient.CreateNotiToUsers(post.Participants, "post", "post:resolved:"+post.Username, id)
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
	if err := s.postRepo.CreateComment(ctx, postId, username, content); err != nil {
		return err
	}
	post, err := s.postRepo.GetPost(ctx, postId)
	if err != nil {
		return err
	}
	if post.Username != username {
		_, err = s.notiClient.CreateNoti(post.Username, "post", "post:comment:"+username, postId)
	}
	return err
}

func (s *Service) DeleteComment(ctx context.Context, postId, username string) error {
	return s.postRepo.DeleteComment(ctx, postId, username)
}

// Interaction

func (s *Service) UpsertInteraction(ctx context.Context, postId, username string, itype entity.InteractionType) error {
	if err := s.postRepo.UpsertInteraction(ctx, postId, username, itype); err != nil {
		return err
	}
	post, err := s.postRepo.GetPost(ctx, postId)
	if err != nil {
		return err
	}
	if post.Username != username {
		_, err = s.notiClient.CreateNoti(post.Username, "post", "post:interaction:"+username, postId)
	}
	return err
}

func (s *Service) DeleteInteraction(ctx context.Context, postId, username string) error {
	return s.postRepo.DeleteInteraction(ctx, postId, username)
}

// Participation

func (s *Service) Participate(ctx context.Context, postId, username string) error {
	if err := s.postRepo.Participate(ctx, postId, username); err != nil {
		return err
	}
	post, err := s.postRepo.GetPost(ctx, postId)
	if err != nil {
		return err
	}
	_, err = s.notiClient.CreateNoti(post.Username, "post", "post:participate:"+username, postId)
	return err
}

func (s *Service) Unparticipate(ctx context.Context, postId, username string) error {
	if err := s.postRepo.Unparticipate(ctx, postId, username); err != nil {
		return err
	}
	post, err := s.postRepo.GetPost(ctx, postId)
	if err != nil {
		return err
	}
	_, err = s.notiClient.CreateNoti(post.Username, "post", "post:unparticipate:"+username, postId)
	return err
}

func fetchAddress(lat, lon float64) (string, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%f&lon=%f&format=json", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		DisplayName string `json:"display_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.DisplayName, nil
}
