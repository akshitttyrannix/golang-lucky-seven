package settings

import (
	"time"

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
