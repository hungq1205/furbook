package post

import (
	"errors"
	"post/api/client"
	"post/api/presenter"
	"post/entity"
)

func ListPostEntityToPresenter(posts []*entity.Post, users map[string]*presenter.User) []*presenter.Post {
	rPosts := make([]*presenter.Post, len(posts))
	for i, p := range posts {
		rPosts[i] = PostEntityToPresenter(p, users[p.Username])
	}
	return rPosts
}

func PostEntityToPresenter(post *entity.Post, user *presenter.User) *presenter.Post {
	return &presenter.Post{
		ID:          post.ID,
		Type:        post.Type,
		Username:    post.Username,
		DisplayName: user.DisplayName,
		UserAvatar:  user.Avatar,
		Content:     post.Content,
		Medias:      post.Medias,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,

		Interactions: post.Interactions,
		CommentNum:   len(post.Comments),

		LostAt:       post.LostAt,
		Area:         presenter.LocationEntityToPresenter(post.Area),
		LastSeen:     presenter.LocationEntityToPresenter(post.LastSeen),
		ContactInfo:  post.ContactInfo,
		IsResolved:   post.IsResolved,
		Participants: post.Participants,
	}
}

func ListPostEntityToPresenterWithClient(posts []*entity.Post, userClient client.UserClient) ([]*presenter.Post, error) {
	usernames := make([]string, len(posts))
	for i, p := range posts {
		usernames[i] = p.Username
	}

	users, err := userClient.FindUsers(usernames)
	if err != nil {
		return nil, err
	}

	userDict := make(map[string]*presenter.User, len(users))
	for _, u := range users {
		userDict[u.Username] = u
	}

	rPosts := make([]*presenter.Post, len(posts))
	for i, p := range posts {
		rPosts[i] = PostEntityToPresenter(p, userDict[p.Username])
	}
	return rPosts, nil
}

func PostEntityToPresenterWithClient(post *entity.Post, userClient client.UserClient) (*presenter.Post, error) {
	users, err := userClient.FindUsers([]string{post.Username})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("user not found")
	}
	user := users[0]

	return PostEntityToPresenter(post, user), nil
}

func CommentEntityToPresenter(comment entity.Comment, user *presenter.User) *presenter.Comment {
	return &presenter.Comment{
		Username:    comment.Username,
		DisplayName: user.DisplayName,
		Avatar:      user.Avatar,
		Content:     comment.Content,
		CreatedAt:   comment.CreatedAt,
	}
}

func ListCommentEntityToPresenterWithClient(comments []entity.Comment, userClient client.UserClient) ([]*presenter.Comment, error) {
	usernames := make([]string, len(comments))
	for i, c := range comments {
		usernames[i] = c.Username
	}

	users, err := userClient.FindUsers(usernames)
	if err != nil {
		return nil, err
	}

	userDict := make(map[string]*presenter.User, len(users))
	for _, u := range users {
		userDict[u.Username] = u
	}

	pComments := make([]*presenter.Comment, len(comments))
	for i, c := range comments {
		pComments[i] = CommentEntityToPresenter(c, userDict[c.Username])
	}
	return pComments, nil
}
