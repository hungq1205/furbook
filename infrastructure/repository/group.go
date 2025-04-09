package repository

import (
	"test/entity"

	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) GetGroup(groupID int) (*entity.Group, error) {
	var group entity.Group
	err := r.db.First(&group, groupID).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRepository) CreateGroup(group *entity.Group) (*entity.Group, error) {
	err := r.db.Create(group).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (r *GroupRepository) UpdateGroup(group *entity.Group) (*entity.Group, error) {
	err := r.db.Save(group).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (r *GroupRepository) DeleteGroup(groupID int) error {
	err := r.db.Delete(&entity.Group{}, groupID).Error
	if err != nil {
		return err
	}
	return nil
}
