package games

import (
	"fmt"
	"log"
	"os"

	"gamesanct.com/lucky-seven/common"
	"gamesanct.com/lucky-seven/common/functions"
	"gamesanct.com/lucky-seven/config/deck"
	"gamesanct.com/lucky-seven/modules/rounds"
	"gamesanct.com/lucky-seven/modules/settings"
	"github.com/google/uuid"
)

var GameState = GAME_STATE_INIT
var IsGameRunning = false

// var setting *settings.Setting

var bettingTimer int16
var bonusTimer int16
var resultTimer int16
var pauseTimer int16

var roundID string
var displayID string

var roundCount int64

var card deck.Card

var today int64

var activeMarkets []string

func InitGame() {
	// deck := deck.NewDeck()

	IsGameRunning = false
	GameState = GAME_STATE_INIT
	today = functions.GetToday()

	roundCount = rounds.GetTodaysRoundCount()

	log.Println("Round count:", roundCount)

	nodeEnv := os.Getenv(common.DOT_ENV_NODE_ENV)
	if !IsGameRunning && nodeEnv != common.NODE_ENV_LOCAL {
		IsGameRunning = true
		startRound()
	}
}

func setTimers(setting *settings.Setting) {
	bettingTimer = setting.BettingTimer
	bonusTimer = setting.BonusTimer
	resultTimer = setting.ResultTimer
	pauseTimer = setting.PauseTimer
}

func setRoundCount() {
	now := functions.GetToday()
	if now == today {
		roundCount += 1
	} else {
		roundCount = 0
		today = now
	}
}

func generateIDs() {
	roundID = uuid.New().String()

	dateFormatted := functions.GetDateFormatted()
	paddedRoundCount := fmt.Sprintf("%05d", roundCount)

	displayID = dateFormatted + common.GAME_DISPLAY_NAME + paddedRoundCount
}

func getActiveMarkets() {
	activeMarkets = []string{"Main Market"}
}

func createRound() {
	round := &rounds.Round{
		RoundID:   roundID,
		DisplayID: displayID,
		GameID:    common.GAME_ID,
		GameName:  common.GAME_NAME,
		Status:    GAME_STATE_INIT,
		StartTime: functions.CurrentTime(),
		CreatedAt: functions.CurrentTime(),
		UpdatedAt: functions.CurrentTime(),
	}

	log.Println("Creating round:", round)

	rounds.Create(round)
}

func startRound() {

	setting, err := settings.GetSettingByID()
	if err != nil {
		log.Fatal("Failed to get setting:", err)
	}

	if !((GameState == GAME_STATE_INIT || GameState == GAME_STATE_END) && IsGameRunning && setting.Status == settings.SETTING_STATUS_ACTIVE) {
		IsGameRunning = false
		log.Fatal("Game is not running or active")
	}

	setTimers(setting)
	setRoundCount()
	generateIDs()
	getActiveMarkets()
	createRound()
}
