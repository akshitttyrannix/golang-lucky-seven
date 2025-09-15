package deck

import (
	"math/rand"
	"time"
)

var SUITS = [4]string{"C", "D", "H", "S"}
var VALUES = [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

type Card struct {
	Suit  string
	Value int
}

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	deck := Deck{
		Cards: make([]Card, 0, 52),
	}

	for _, suit := range SUITS {
		for _, value := range VALUES {
			deck.Cards = append(deck.Cards, Card{Suit: suit, Value: value})
		}
	}

	deck.shuffle()

	return deck
}

func (d *Deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
