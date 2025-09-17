package settings

import (
	"gamesanct.com/lucky-seven/common/customerror"
	"gamesanct.com/lucky-seven/common/messages"
	"gamesanct.com/lucky-seven/common/success"
	"github.com/gin-gonic/gin"
)

func CreateSetting(ctx *gin.Context) {
	var dto CreateSettingDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		customerror.BadRequest(ctx, err)
		return
	}

	setting := CreateSettingEntity(dto)

	if err := Create(setting); err != nil {
		customerror.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.SETTING_CREATED, setting)
}

func StartGame(ctx *gin.Context) {
	setting := &Setting{
		SettingID: SETTING_ID,
		Status:    SETTING_STATUS_ACTIVE,
	}

	if err := UpdateOne(setting); err != nil {
		customerror.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.SETTING_STARTED, nil)
}

func StopGame(ctx *gin.Context) {
	setting := &Setting{
		SettingID: SETTING_ID,
		Status:    SETTING_STATUS_INACTIVE,
	}

	if err := UpdateOne(setting); err != nil {
		customerror.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.SETTING_STOPPED, nil)
}
