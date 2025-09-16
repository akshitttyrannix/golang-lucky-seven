package settings

import (
	"time"

	"gamesanct.com/lucky-seven/common/error"
	"gamesanct.com/lucky-seven/common/messages"
	"gamesanct.com/lucky-seven/common/success"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateSetting(ctx *gin.Context) {
	var createSettingDTO CreateSettingDTO
	if err := ctx.ShouldBindJSON(&createSettingDTO); err != nil {
		error.BadRequest(ctx, err)
		return
	}

	setting := &Setting{
		SettingID:    uuid.New().String(),
		Status:       createSettingDTO.Status,
		BettingTimer: createSettingDTO.BettingTimer,
		BonusTimer:   createSettingDTO.BonusTimer,
		ResultTimer:  createSettingDTO.ResultTimer,
		PauseTimer:   createSettingDTO.PauseTimer,
		CreatedAt:    uint32(time.Now().Unix()),
		UpdatedAt:    uint32(time.Now().Unix()),
	}

	if err := Create(setting); err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.SETTING_CREATED, setting)
}
