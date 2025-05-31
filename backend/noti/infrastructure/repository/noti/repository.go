package noti

import (
	"noti/entity"
	"noti/util"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *NotificationRepository {
	db.AutoMigrate(&entity.Notification{})
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) GetNotification(id int) (*entity.Notification, error) {
	var noti entity.Notification
	if err := r.db.Where("id = ?", id).First(&noti).Error; err != nil {
		return nil, err
	}
	return &noti, nil
}

func (r *NotificationRepository) GetNotificationsOfUser(username string, pagination util.Pagination) ([]*entity.Notification, error) {
	var notifications []*entity.Notification
	err := r.db.
		Model(&entity.Notification{}).
		Where("username = ?", username).
		Find(&notifications).
		Order("created_at DESC").
		Offset(pagination.Offset()).
		Limit(pagination.Size).
		Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *NotificationRepository) CreateNotification(username, icon, desc, link string) (*entity.Notification, error) {
	noti := entity.Notification{
		Username: username,
		Icon:     icon,
		Desc:     desc,
		Link:     link,
	}
	if err := r.db.Create(&noti).Error; err != nil {
		return nil, err
	}
	return &noti, nil
}

func (r *NotificationRepository) CreateNotificationToUsers(usernames []string, icon, desc, link string) ([]*entity.Notification, error) {
	var notifications []*entity.Notification
	for _, username := range usernames {
		noti := entity.Notification{
			Username: username,
			Icon:     icon,
			Desc:     desc,
			Link:     link,
		}
		if err := r.db.Create(&noti).Error; err != nil {
			return nil, err
		}
		notifications = append(notifications, &noti)
	}
	return notifications, nil
}

func (r *NotificationRepository) UpdateNotification(id int, read bool) (*entity.Notification, error) {
	noti, err := r.GetNotification(id)
	if err != nil {
		return nil, err
	}
	noti.Read = read
	if err := r.db.Save(noti).Error; err != nil {
		return nil, err
	}
	return noti, nil
}

func (r *NotificationRepository) DeleteNotification(id int) error {
	if err := r.db.Where("id = ?", id).Delete(&entity.Notification{}).Error; err != nil {
		return err
	}
	return nil
}
