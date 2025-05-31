package noti

import (
	"noti/entity"
	"noti/infrastructure/repository/noti"
	"noti/util"
)

type Service struct {
	notiRepo *noti.NotificationRepository
}

func NewService(notiRepo *noti.NotificationRepository) *Service {
	return &Service{
		notiRepo: notiRepo,
	}
}

func (s *Service) GetNoti(id int) (*entity.Notification, error) {
	usr, err := s.notiRepo.GetNotification(id)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) GetNotisOfUser(username string, pagination util.Pagination) ([]*entity.Notification, error) {
	notis, err := s.notiRepo.GetNotificationsOfUser(username, pagination)
	if err != nil {
		return nil, err
	}
	return notis, nil
}

func (s *Service) CreateNoti(username, icon, desc, link string) (*entity.Notification, error) {
	usr, err := s.notiRepo.CreateNotification(username, icon, desc, link)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) CreateNotiToUsers(usernames []string, icon, desc, link string) ([]*entity.Notification, error) {
	return s.notiRepo.CreateNotificationToUsers(usernames, icon, desc, link)
}

func (s *Service) UpdateNoti(id int, read bool) (*entity.Notification, error) {
	usr, err := s.notiRepo.UpdateNotification(id, read)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) DeleteNoti(id int) error {
	if err := s.notiRepo.DeleteNotification(id); err != nil {
		return err
	}
	return nil
}
