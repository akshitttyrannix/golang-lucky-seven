package games

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gamesanct.com/lucky-seven/common"
	"gamesanct.com/lucky-seven/common/functions"
	"gamesanct.com/lucky-seven/config/deck"
	"gamesanct.com/lucky-seven/modules/rounds"
	"gamesanct.com/lucky-seven/modules/settings"
	"github.com/google/uuid"
)

var GameState = GAME_STATE_INIT
var IsGameRunning = false

var bettingTimer int16
var bonusTimer int16
var resultTimer int16
var pauseTimer int16

var roundID string
var displayID string

var roundCount int16

var newDeck deck.Deck
var card deck.Card

var today int64

var activeMarkets []string

var round *rounds.Round

func InitGame() {
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
	log.Println("Setting timers")

	bettingTimer = setting.BettingTimer
	bonusTimer = setting.BonusTimer
	resultTimer = setting.ResultTimer
	pauseTimer = setting.PauseTimer
}

func setRoundCount() {
	log.Println("Setting round count")

	now := functions.GetToday()
	if now == today {
		roundCount += 1
	} else {
		roundCount = 0
		today = now
	}
}

func generateIDs() {
	log.Println("Generating IDs")

	roundID = uuid.New().String()

	dateFormatted := functions.GetDateFormatted()
	paddedRoundCount := fmt.Sprintf("%05d", roundCount)

	displayID = dateFormatted + common.GAME_DISPLAY_NAME + paddedRoundCount
	log.Println("Round ID:", roundID)
	log.Println("Display ID:", displayID)
}

func getActiveMarkets() {
	activeMarkets = []string{"Main Market"}
}

func createRound() {
	round = &rounds.Round{
		RoundID:   roundID,
		DisplayID: displayID,
		GameID:    common.GAME_ID,
		GameName:  common.GAME_NAME,
		State:     GAME_STATE_INIT,
		StartTime: functions.CurrentTime(),
		CreatedAt: functions.CurrentTime(),
		UpdatedAt: functions.CurrentTime(),
	}

	rounds.Create(round)

	log.Println("Created round:", round)
}

func updateRound() {
	// round.Result = card
	round.State = GameState
	round.UpdatedAt = functions.CurrentTime()
	round.EndTime = functions.CurrentTime()
	rounds.Update(round)

	log.Println("Updated round:", round)
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

	newDeck = deck.NewDeck()

	log.Println("New deck length:", len(newDeck.Cards))

	setTimers(setting)
	setRoundCount()
	generateIDs()
	getActiveMarkets()
	createRound()

	GameState = GAME_STATE_BETTING
	updateRound()

	runBettingTimer()
}

func runBettingTimer() {
	if bettingTimer > -1 && GameState == GAME_STATE_BETTING {
		time.Sleep(time.Second)
		log.Println("Betting timer: " + strconv.Itoa(int(bettingTimer)))
		bettingTimer -= 1
		runBettingTimer()
	} else {
		GameState = GAME_STATE_BONUS
		updateRound()
		runBonusTimer()
	}
}

func runBonusTimer() {
	if bonusTimer > -1 && GameState == GAME_STATE_BONUS {
		time.Sleep(time.Second)
		log.Println("Bonus timer: " + strconv.Itoa(int(bonusTimer)))
		bonusTimer -= 1
		runBonusTimer()
	} else {
		GameState = GAME_STATE_RESULT
		updateRound()

		card = newDeck.Cards[0]
		log.Println("Card:", card)

		announceWinner()

		runResultTimer()
	}
}

func runResultTimer() {
	if resultTimer > -1 && GameState == GAME_STATE_RESULT {
		time.Sleep(time.Second)
		log.Println("Result timer: " + strconv.Itoa(int(resultTimer)))
		resultTimer -= 1
		runResultTimer()
	} else {
		GameState = GAME_STATE_PAUSE
		updateRound()
		runPauseTimer()
	}
}

func runPauseTimer() {
	if pauseTimer > -1 && GameState == GAME_STATE_PAUSE {
		time.Sleep(time.Second)
		log.Println("Pause timer: " + strconv.Itoa(int(pauseTimer)))
		pauseTimer -= 1
		runPauseTimer()
	} else {
		GameState = GAME_STATE_END
		updateRound()

		startRound()
	}
}

func announceWinner() {
	log.Println("Announcing winner")

	round.Result = card
	updateRound()

	winner := map[string]any{
		"card":    card,
		"markets": activeMarkets,
	}

	round.Result = winner
	updateRound()
}
