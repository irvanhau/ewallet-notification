package repository

import (
	"context"
	"ewallet-notification/internal/models"

	"gorm.io/gorm"
)

type EmailRepository struct {
	DB *gorm.DB
}

func (r *EmailRepository) GetTemplate(ctx context.Context, templateName string) (models.NotificationTemplate, error) {
	var (
		resp models.NotificationTemplate
	)
	err := r.DB.Where("template_name = ?", templateName).Last(&resp).Error
	return resp, err
}

func (r *EmailRepository) InsertNotificationHistory(ctx context.Context, notif *models.NotificationHistory) error {
	return r.DB.Create(notif).Error
}
