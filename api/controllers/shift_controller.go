package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerShift struct {
	shift *usecase.UseCaseShift
	*baseController
}

func NewControllerShift(shift *usecase.UseCaseShift, base *baseController) *ControllerShift {
	return &ControllerShift{
		shift:          shift,
		baseController: base,
	}
}

func (u *ControllerShift) AddShift(ctx *gin.Context) {
	var req entities.ShiftAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := u.shift.AddShift(ctx, &req)
	u.baseController.Response(ctx, resp, err)
}
func (u *ControllerShift) GetAllShifts(ctx *gin.Context) {
	resp, err := u.shift.GetAllShifts(ctx)
	u.baseController.Response(ctx, resp, err)
}
