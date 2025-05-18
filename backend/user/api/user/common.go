package user

import (
	"user/entity"
	"user/presenter"
	"user/usecase/friend"
)

func UserEntityToPresenter(in *entity.User, friendSerivce friend.UseCase) (*presenter.User, error) {
	friendNum, err := friendSerivce.CountFriends(in.Username)
	if err != nil {
		return nil, err
	}

	return &presenter.User{
		Username:    in.Username,
		DisplayName: in.DisplayName,
		Bio:         in.Bio,
		Avatar:      in.Avatar,
		FriendNum:   friendNum,
	}, nil
}

func ListUserEntityToPresenter(in []*entity.User, friendSerivce friend.UseCase) ([]*presenter.User, error) {
	out := make([]*presenter.User, 0)
	for _, user := range in {
		pUser, err := UserEntityToPresenter(user, friendSerivce)
		if err != nil {
			return nil, err
		}
		out = append(out, pUser)
	}
	return out, nil
}
