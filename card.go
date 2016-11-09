// to implement domino single card
package domino

import "fmt"

// type for card value
type CardValue int

// constant of card value, there are only 28 cards posibilities
const (
	BlankBlank CardValue = iota
	BlankOne
	BlankTwo
	BlankThree
	BlankFour
	BlankFive
	BlankSix
	OneOne
	OneTwo
	OneThree
	OneFour
	OneFive
	OneSix
	TwoTwo
	TwoThree
	TwoFour
	TwoFive
	TwoSix
	ThreeThree
	ThreeFour
	ThreeFive
	ThreeSix
	FourFour
	FourFive
	FourSix
	FiveFive
	FiveSix
	SixSix
)

// type for directing the sidevalue of card
type Side int

// constant for side only two
const (
	None Side = iota
	First
	Second
)

// Card representing first and second side value
type Card struct {
	first  SideValue
	second SideValue
}

// a bit tricky here, i'm using recursive
// constant for helping recursive function
const ZeroLevel int = 7

// the recursive function
func FindCardValue(cardValue, level int) (first, second int) {
	if cardValue/(ZeroLevel-level) == 0 {
		return level, cardValue + level
	} else {
		return FindCardValue(cardValue-(ZeroLevel-level), level+1)
	}
}

// function for making a card by inputing the card value
func MakeCard(cv CardValue) Card {
	first, second := FindCardValue(int(cv), 0)
	return Card{SideValue{DotValue(first), false}, SideValue{DotValue(second), false}}
}

// function for making a card by inputing the card value and bool of the match status
func MakeCardBool(cv CardValue, firstBool, secondBool bool) Card {
	first, second := FindCardValue(int(cv), 0)
	return Card{SideValue{DotValue(first), firstBool}, SideValue{DotValue(second), secondBool}}
}

// function for stringer
func (card Card) String() string {
	return fmt.Sprintf("Card, first side: %s, second side: %s", card.first, card.second)
}

func (card Card) GetCardValue() CardValue {
	var firstValue int
	for i := 0; i < int(card.first.dotvalue); i++ {
		firstValue = firstValue + 6 - i
	}
	return CardValue(firstValue + int(card.second.dotvalue))
}

// funct for get the side, i'm using pointer so the return value can be changing in the real card
func (card *Card) GetSide(side Side) *SideValue {
	if side == First {
		return &card.first
	} else {
		return &card.second
	}
}

// funct for changing the side of the value
func ChangeCardMatchValue(card *Card, side Side, newMatch bool) bool {
	return card.GetSide(side).ChangeMatched(newMatch)
}

// function for match the first card and the second card
// if executed false, it's only used for checking,
// return bool if able to match or not
// return 2 side for get which side is matched from first and second card
func MatchCard(firstcard, secondcard *Card, executed bool) (bool, Side, Side) {
	var firstSide, secondSide Side

	switch {
	case CheckIfAbleToMatch(*(firstcard.GetSide(First)), *(secondcard.GetSide(First))):
		firstSide = First
		secondSide = First
	case CheckIfAbleToMatch(*(firstcard.GetSide(First)), *(secondcard.GetSide(Second))):
		firstSide = First
		secondSide = Second
	case CheckIfAbleToMatch(*(firstcard.GetSide(Second)), *(secondcard.GetSide(First))):
		firstSide = Second
		secondSide = First
	case CheckIfAbleToMatch(*(firstcard.GetSide(Second)), *(secondcard.GetSide(Second))):
		firstSide = Second
		secondSide = Second
	}

	b := CheckIfAbleToMatch(*firstcard.GetSide(firstSide), *secondcard.GetSide(secondSide))

	if b && executed {
		firstcard.GetSide(firstSide).ChangeMatched(true)
		secondcard.GetSide(secondSide).ChangeMatched(true)
	}
	return b, firstSide, secondSide
}
