package deck

import (
	"math/rand"
	"strconv"
)

type Card struct {
	Suit string
	Name string
}

type Deck []Card

func New(n int) Deck {
	var deck []Card
	var card Card
	for k := 0; k < n; k++ {
		for i := 1; i < 5; i++ {
			switch i {
			case 1:
				card.Suit = "spade"
			case 2:
				card.Suit = "diamond"
			case 3:
				card.Suit = "club"
			case 4:
				card.Suit = "heart"
			}
			for g := 1; g < 14; g++ {
				switch g {
				case 1:
					card.Name = "A"
				case 11:
					card.Name = "J"
				case 12:
					card.Name = "Q"
				case 13:
					card.Name = "K"
				default:
					card.Name = strconv.Itoa(g)
				}
				deck = append(deck, card)
			}
		}
	}
	return deck
}

func (deck *Deck) Shuffle() {
	cards := *deck
	p := rand.Perm(len(*deck))
	for i, v := range p {
		(*deck)[i] = cards[v]
	}
}

func (hand *Deck) TakeCard(deck *Deck, n int) {
	for i := 0; i < n; i++ {
		*hand = append(*hand, (*deck)[len(*deck)-1])
		*deck = (*deck)[:len(*deck)-1]
	}
}
