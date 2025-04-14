package group

import (
	"message/api/presenter"
	"message/entity"
	"message/usecase/group"
)

func groupEntityToPresenter(in *entity.Group, groupService group.UseCase) (*presenter.Group, error) {
	users, err := groupService.GetMembers(in.ID)
	if err != nil {
		return nil, err
	}
	return &presenter.Group{
		ID:        in.ID,
		Name:      in.Name,
		IsDirect:  in.IsDirect,
		OwnerName: in.OwnerName,
		Members:   users,
	}, nil
}

func groupListEntityToPresenter(in []*entity.Group, groupService group.UseCase) (out []*presenter.Group, err error) {
	for _, g := range in {
		pg, err := groupEntityToPresenter(g, groupService)
		if err != nil {
			return nil, err
		}
		out = append(out, pg)
	}
	return out, nil
}
