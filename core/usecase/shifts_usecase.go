package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseShift struct {
	shift domain.RepositoryShift
}

func NewUseCaseShift(shift domain.RepositoryShift) *UseCaseShift {
	return &UseCaseShift{
		shift: shift,
	}
}
func (u *UseCaseShift) AddShift(ctx context.Context, req *entities.ShiftAddReq) (*entities.ShiftAddResp, error) {
	err := u.shift.CreateShift(ctx, &domain.Shift{
		ID:         utils.GenerateUniqueKey(),
		ShiftName:  req.ShiftName,
		Salary:     req.Salary,
		ShiftStart: req.ShiftStart,
		ShiftEnd:   req.ShiftEnd,
	})
	if err != nil {
		return &entities.ShiftAddResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.ShiftAddResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (u *UseCaseShift) GetAllShifts(ctx context.Context) (*entities.ShiftGetAllResp, error) {
	resp, err := u.shift.GetAllShifts(ctx)
	if err != nil {
		return &entities.ShiftGetAllResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.ShiftGetAllResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Shifts: resp,
	}, nil
}
