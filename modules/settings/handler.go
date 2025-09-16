package settings

import (
	"time"

	"gamesanct.com/lucky-seven/common/error"
	"gamesanct.com/lucky-seven/common/messages"
	"gamesanct.com/lucky-seven/common/success"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateSettingEntity(dto CreateSettingDTO) *Setting {
	return &Setting{
		SettingID:    uuid.New().String(),
		Status:       dto.Status,
		BettingTimer: dto.BettingTimer,
		BonusTimer:   dto.BonusTimer,
		ResultTimer:  dto.ResultTimer,
		PauseTimer:   dto.PauseTimer,
		CreatedAt:    uint32(time.Now().Unix()),
		UpdatedAt:    uint32(time.Now().Unix()),
	}
}

func CreateSetting(ctx *gin.Context) {
	var dto CreateSettingDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		error.BadRequest(ctx, err)
		return
	}

	setting := CreateSettingEntity(dto)

	if err := Create(setting); err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.SETTING_CREATED, setting)
}
