package repository

import (
	"context"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type shiftRepository struct {
	db *gorm.DB
}

func NewCollectionShift(shif *adapter.PostGresql) domain.RepositoryShift {
	return &shiftRepository{
		db: shif.CreateCollection(),
	}
}

func (r *shiftRepository) CreateShift(ctx context.Context, shift *domain.Shift) error {
	return r.db.WithContext(ctx).Create(shift).Error
}

func (r *shiftRepository) GetShiftByID(ctx context.Context, id int64) (*domain.Shift, error) {
	var shift domain.Shift
	if err := r.db.WithContext(ctx).First(&shift, id).Error; err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *shiftRepository) UpdateShift(ctx context.Context, shift *domain.Shift) error {
	return r.db.WithContext(ctx).Save(shift).Error
}

func (r *shiftRepository) DeleteShift(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&domain.Shift{}, id).Error
}

func (r *shiftRepository) GetAllShifts(ctx context.Context) ([]*domain.Shift, error) {
	var shifts []*domain.Shift
	if err := r.db.WithContext(ctx).Find(&shifts).Error; err != nil {
		return nil, err
	}
	return shifts, nil
}
