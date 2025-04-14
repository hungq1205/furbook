package message

import (
	"test/api/presenter"
	"test/entity"
)

func messageEntityToPresenter(in *entity.Message) (out *presenter.Message) {
	return &presenter.Message{
		ID:        in.ID,
		Content:   in.Content,
		Username:  in.Username,
		GroupID:   in.GroupID,
		CreatedAt: in.CreatedAt,
	}
}

func messageListEntityToPresenter(in []*entity.Message) (out []*presenter.Message) {
	for _, msg := range in {
		out = append(out, messageEntityToPresenter(msg))
	}
	return out
}
