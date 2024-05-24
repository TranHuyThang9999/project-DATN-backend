package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseUser struct {
	user  domain.RepositoryUser
	file  domain.RepositoryFileStorages
	trans domain.RepositoryTransaction
}

func NewUseCaseUser(
	user domain.RepositoryUser,
	file domain.RepositoryFileStorages,
	trans domain.RepositoryTransaction,
) *UseCaseUser {
	return &UseCaseUser{
		user:  user,
		file:  file,
		trans: trans,
	}
}

func (u *UseCaseUser) AddUser(ctx context.Context, user []*entities.Users) (*entities.UserResp, error) {

	var userAdd = make([]*domain.Users, 0)

	for _, v := range user {
		userAdd = append(userAdd, &domain.Users{
			ID:         utils.GenerateUniqueKey(),
			UserName:   v.UserName,
			ShiftNames: v.ShiftNames,
			CinemaName: v.CinemaName,
			Age:        v.Age,
			Address:    v.Address,
			Role:       enums.ROLE_STAFF,
			CreatedAt:  utils.GenerateTimestamp(),
			UpdatedAt:  utils.GenerateTimestamp(),
		})
	}
	err := u.user.AddUser(ctx, userAdd)
	if err != nil {
		return &entities.UserResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	return &entities.UserResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Created: utils.GenerateTimestamp(),
	}, nil
}
func (u *UseCaseUser) GetAllUserById(ctx context.Context, req *entities.UserRequestFindByForm) (*entities.UserRespFindByForm, error) {
	resp, err := u.user.GetAllUserStaffs(ctx, &domain.UsersReqByForm{
		Id:        req.Id,
		UserName:  req.UserName,
		Age:       req.Age,
		Address:   req.Address,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	})
	if err != nil {
		return &entities.UserRespFindByForm{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, err
	}
	return &entities.UserRespFindByForm{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Users:   resp,
		Created: utils.GenerateTimestamp(),
	}, nil
}
