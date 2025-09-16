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
	var setting Setting
	if err := ctx.ShouldBindJSON(&setting); err != nil {
		error.BadRequest(ctx, err)
		return
	}

	setting.SettingID = uuid.New().String()
	setting.CreatedAt = time.Now().Unix()
	setting.UpdatedAt = time.Now().Unix()

	if err := Create(&setting); err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.SETTING_CREATED, setting)
}
