package whist

import (
	"testing"

	"github.com/edstell/card"
	"github.com/stretchr/testify/assert"
)

func TestTrick_Winner(t *testing.T) {
	tcs := []struct {
		description    string
		trick          Trick
		trump          card.Suit
		expectedResult int
	}{
		{
			description:    "assert highest card wins when suit followed",
			trick:          Trick{card.CA, card.C2, card.C3, card.C4},
			trump:          card.Diamonds,
			expectedResult: 0,
		},
		{
			description:    "assert highest card wins when suit not followed",
			trick:          Trick{card.C2, card.D3, card.H4, card.H6},
			trump:          card.Spades,
			expectedResult: 0,
		},
		{
			description:    "assert trump wins",
			trick:          Trick{card.C2, card.CK, card.S2, card.CQ},
			trump:          card.Spades,
			expectedResult: 2,
		},
		{
			description:    "assert highest trump wins",
			trick:          Trick{card.C2, card.S3, card.S2, card.CQ},
			trump:          card.Spades,
			expectedResult: 1,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.trick.Winner(tc.trump))
		})
	}
}
