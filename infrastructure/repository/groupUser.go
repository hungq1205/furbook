package repository

import (
	"test/entity"

	"gorm.io/gorm"
)

type GroupUserRepository struct {
	db *gorm.DB
}

func NewGroupUserRepository(db *gorm.DB) *GroupUserRepository {
	return &GroupUserRepository{db: db}
}

func (r *GroupUserRepository) GetUsersInGroup(groupID int) ([]string, error) {
	var users []string
	err := r.db.
		Model(&entity.GroupUser{}).
		Where("group_id = ?", groupID).
		Pluck("username", &users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GroupUserRepository) GetGroupsOfUser(username string) ([]*entity.Group, error) {
	var groups []*entity.Group
	err := r.db.
		Joins("join groups g on g.id = group_users.group_id").
		Where("username = ?", username).
		Select("g.*").
		Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *GroupUserRepository) GetDirectGroup(userA string, userB string) (*entity.Group, error) {
	var group entity.Group
	err := r.db.
		Model(&entity.Group{}).
		Joins("join group_users gu1 on gu1.group_id = groups.id").
		Joins("join group_users gu2 on gu2.group_id = groups.id").
		Where("g.is_direct = ? AND gu1.username = ? AND gu2.username = ?", true, userA, userB).
		Take(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupUserRepository) AddUserToGroup(groupUser *entity.GroupUser) (*entity.GroupUser, error) {
	err := r.db.Create(groupUser).Error
	if err != nil {
		return nil, err
	}
	return groupUser, nil
}

func (r *GroupUserRepository) RemoveUserFromGroup(groupUser *entity.GroupUser) error {
	err := r.db.Delete(groupUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *GroupUserRepository) RemoveUsersFromGroup(groupID int) error {
	err := r.db.Where("group_id = ?", groupID).Delete(&entity.GroupUser{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *GroupUserRepository) RemoveUserMemberships(username string) error {
	err := r.db.Where("username = ?", username).Delete(&entity.GroupUser{}).Error
	if err != nil {
		return err
	}
	return nil
}
