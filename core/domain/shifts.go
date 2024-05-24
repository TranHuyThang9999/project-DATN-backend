package domain

import "context"

type Shift struct {
	ID         int64   `json:"id"`
	ShiftName  string  `json:"shift_name"`
	Salary     float64 `json:"salary"`
	ShiftStart int     `json:"shift_start"`
	ShiftEnd   int     `json:"shift_end"`
}

type RepositoryShift interface {
	CreateShift(ctx context.Context, shift *Shift) error
	GetShiftByID(ctx context.Context, id int64) (*Shift, error)
	UpdateShift(ctx context.Context, shift *Shift) error
	DeleteShift(ctx context.Context, id int64) error
	GetAllShifts(ctx context.Context) ([]*Shift, error)
}
