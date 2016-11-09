package domino

type Deck struct {
	cards []Card
}

var DominoDeck = []CardValue{BlankBlank, BlankOne, BlankTwo, BlankThree, BlankFour, BlankFive, BlankSix, OneOne, OneTwo, OneThree, OneFour, OneFive, OneSix, TwoTwo, TwoThree, TwoFour, TwoFive, TwoSix, ThreeThree, ThreeFour, ThreeFive, ThreeSix, FourFour, FourFive, FourSix, FiveFive, FiveSix, SixSix}

func NewDeck(shuffle bool) Deck {
	deck := NewDeckSpecified(shuffle, DominoDeck)
	return deck
}

func NewDeckSpecified(shuffle bool, cardValues []CardValue) Deck {
	cards := make([]Card, len(cardValues))
	for cindex, card := range cardValues {
		cards[cindex] = MakeCard(card)
	}
	deck := Deck{cards}
	return deck
}
