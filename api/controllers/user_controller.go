package controllers

import (
	"flick_tickets/core/entities"
	"flick_tickets/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllersUser struct {
	user *usecase.UseCaseUser
	*baseController
}

func NewControllersUser(
	user *usecase.UseCaseUser,
	baseController *baseController,
) *ControllersUser {
	return &ControllersUser{
		user:           user,
		baseController: baseController,
	}
}
func (u *ControllersUser) AddUser(ctx *gin.Context) {
	var req []*entities.Users
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := u.user.AddUser(ctx, req)
	u.baseController.Response(ctx, resp, err)
}
func (u *ControllersUser) GetListStaff(ctx *gin.Context) {

	var req entities.UserRequestFindByForm

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := u.user.GetAllUserById(ctx, &req)
	u.baseController.Response(ctx, resp, err)

}
