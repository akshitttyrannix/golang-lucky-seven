package settings

import (
	"time"

	"gamesanct.com/lucky-seven/common/error"
	"gamesanct.com/lucky-seven/common/messages"
	"gamesanct.com/lucky-seven/common/success"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateSetting(c *gin.Context) {
	var setting Setting
	if err := c.ShouldBindJSON(&setting); err != nil {
		error.BadRequest(c, err)
		return
	}

	setting.SettingID = uuid.New().String()
	setting.CreatedAt = time.Now().Unix()
	setting.UpdatedAt = time.Now().Unix()

	if err := create(&setting); err != nil {
		error.SomethingWentWrong(c, err)
		return
	}

	success.Success(c, messages.SETTING_CREATED, setting)
}
