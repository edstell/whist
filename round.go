package whist

import "github.com/edstell/card"

type Round struct {
	Lead  int
	Trick Trick
}

func (r Round) Winner(trumps card.Suit) int {
	return (r.Trick.Winner(trumps) + r.Lead) % 4
}
