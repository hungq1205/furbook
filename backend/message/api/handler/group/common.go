package group

import (
	"message/api/presenter"
	"message/entity"
	"message/usecase/group"
	"message/usecase/message"
)

func groupEntityToPresenter(in *entity.Group, groupService group.UseCase, messageService message.UseCase) (*presenter.Group, error) {
	users, err := groupService.GetMembers(in.ID)
	if err != nil {
		return nil, err
	}
	lastMessageEntity, err := messageService.GetLastMessage(in.ID)
	if err != nil {
		return nil, err
	}
	lastMessage := presenter.Message{
		ID:        lastMessageEntity.ID,
		Content:   lastMessageEntity.Content,
		Username:  lastMessageEntity.Username,
		GroupID:   lastMessageEntity.GroupID,
		CreatedAt: lastMessageEntity.CreatedAt,
	}

	return &presenter.Group{
		ID:          in.ID,
		Name:        in.Name,
		IsDirect:    in.IsDirect,
		OwnerName:   in.OwnerName,
		Members:     users,
		LastMessage: lastMessage,
	}, nil
}

func groupListEntityToPresenter(in []*entity.Group, groupService group.UseCase, messageService message.UseCase) (out []*presenter.Group, err error) {
	out = make([]*presenter.Group, 0)
	for _, g := range in {
		pg, err := groupEntityToPresenter(g, groupService, messageService)
		if err != nil {
			return nil, err
		}
		out = append(out, pg)
	}
	return out, nil
}
