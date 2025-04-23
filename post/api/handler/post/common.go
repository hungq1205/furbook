package post

import (
	"post/api/presenter"
	"post/entity"
)

func ListPostEntityToPresenter(posts []*entity.Post, users map[uint]*presenter.User) []*presenter.Post {
	rPosts := make([]*presenter.Post, len(posts))
	for i, p := range posts {
		rPosts[i] = PostEntityToPresenter(p, users[p.UserID])
	}
	return rPosts
}

func PostEntityToPresenter(post *entity.Post, user *presenter.User) *presenter.Post {
	return &presenter.Post{
		ID:         post.ID,
		UserID:     post.UserID,
		Username:   user.Username,
		UserAvatar: user.Avatar,
		Content:    post.Content,
		Medias:     post.Medias,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,

		Interactions: post.Interactions,
		CommentNum:   post.CommentNum,

		LostAt:       post.LostAt,
		Area:         post.Area,
		LastSeen:     post.LastSeen,
		ContactInfo:  post.ContactInfo,
		IsResolved:   post.IsResolved,
		Participants: post.Participants,
	}
}
