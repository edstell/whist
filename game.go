package whist

import "github.com/edstell/card"

type Game struct {
	trump  card.Suit
	Rounds []Round
}
