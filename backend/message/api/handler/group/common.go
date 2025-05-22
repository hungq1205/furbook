package group

import (
	"message/api/client"
	"message/api/presenter"
	"message/entity"
	"message/usecase/group"
	"message/usecase/message"
)

func groupEntityToPresenter(in *entity.Group, username string, groupService group.UseCase, messageService message.UseCase, userService client.UserClient) (*presenter.Group, error) {
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

	var oppUsername string
	for _, u := range users {
		if u != username {
			oppUsername = u
			break
		}
	}
	if in.IsDirect {
		directUser, err := userService.FindUsers([]string{oppUsername})
		if err != nil {
			return nil, err
		}
		return &presenter.Group{
			ID:          in.ID,
			Name:        directUser[0].DisplayName,
			Avatar:      directUser[0].Avatar,
			IsDirect:    in.IsDirect,
			LastMessage: lastMessage,
		}, nil
	} else {
		return &presenter.Group{
			ID:          in.ID,
			Name:        in.Name,
			IsDirect:    in.IsDirect,
			OwnerName:   in.OwnerName,
			Members:     users,
			LastMessage: lastMessage,
		}, nil
	}
}

func groupListEntityToPresenter(in []*entity.Group, username string, groupService group.UseCase, messageService message.UseCase, userService client.UserClient) (out []*presenter.Group, err error) {
	out = make([]*presenter.Group, 0)
	for _, g := range in {
		pg, err := groupEntityToPresenter(g, username, groupService, messageService, userService)
		if err != nil {
			return nil, err
		}
		out = append(out, pg)
	}
	return out, nil
}
