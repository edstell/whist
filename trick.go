package whist

import "github.com/edstell/card"

// Trick of four cards in the order they were played.
type Trick [4]card.Card

// Winner (index) given the trump suit.
func (t Trick) Winner(trump card.Suit) int {
	norm := normalize(t[0].Suit(), trump)
	w := 0
	for i, c := range t {
		if norm(c) > norm(t[w]) {
			w = i
		}
	}
	return w
}

// normalize the value of the card given the current and trump suit.
func normalize(s, t card.Suit) func(card.Card) int {
	return func(c card.Card) int {
		switch c.Suit() {
		case s:
			return aceHigh(c.Rank())
		case t:
			return aceHigh(c.Rank()) + 13
		default:
			return 0
		}
	}
}

func aceHigh(r card.Rank) int {
	return (int(r) + 12) % 13
}
