package domino

import (
	"fmt"
	"testing"
)

func TestCreateCard(t *testing.T) {
	cases := []struct {
		val      CardValue
		expected Card
	}{
		{BlankBlank, Card{SideValue{Blank, false}, SideValue{Blank, false}}},
		{BlankOne, Card{SideValue{Blank, false}, SideValue{One, false}}},
		{BlankSix, Card{SideValue{Blank, false}, SideValue{Six, false}}},
		{OneOne, Card{SideValue{One, false}, SideValue{One, false}}},
		{OneSix, Card{SideValue{One, false}, SideValue{Six, false}}},
		{TwoFour, Card{SideValue{Two, false}, SideValue{Four, false}}},
		{ThreeFour, Card{SideValue{Three, false}, SideValue{Four, false}}},
		{FiveFive, Card{SideValue{Five, false}, SideValue{Five, false}}},
		{SixSix, Card{SideValue{Six, false}, SideValue{Six, false}}},
	}
	for _, c := range cases {
		card := MakeCard(c.val)
		if card != c.expected {
			t.Errorf("Failed, result %s, expected %s", card, c.expected)
		}
	}
}

func TestStringerCard(t *testing.T) {
	cases := []struct {
		card     Card
		expected string
	}{
		{MakeCard(BlankBlank), "Card, first side: Blank and not matched, second side: Blank and not matched"},
		{MakeCard(ThreeFour), "Card, first side: 3 and not matched, second side: 4 and not matched"},
		{MakeCard(TwoFive), "Card, first side: 2 and not matched, second side: 5 and not matched"},
		{MakeCard(SixSix), "Card, first side: 6 and not matched, second side: 6 and not matched"},
	}
	for _, c := range cases {
		result := fmt.Sprintf("%s", c.card)
		if result != c.expected {
			t.Errorf("Failed, got \"%s\", expect \"%s\"", result, c.expected)
		}
	}
}

func TestGetSide(t *testing.T) {
	cases := []struct {
		card     Card
		side     Side
		expected SideValue
	}{
		{MakeCard(BlankBlank), First, SideValue{Blank, false}},
		{MakeCard(OneFive), Second, SideValue{Five, false}},
		{MakeCard(SixSix), First, SideValue{Six, false}},
	}
	for _, c := range cases {
		result := c.card.GetSide(c.side)
		if *result != c.expected {
			t.Errorf("Failed, got %s, expected %s", result, c.expected)
		}
	}
}

func TestChangeCardMatched(t *testing.T) {
	cases := []struct {
		card     Card
		side     Side
		newMatch bool
		expected Card
	}{
		{Card{SideValue{Blank, false}, SideValue{Blank, false}}, First, true, Card{SideValue{Blank, true}, SideValue{Blank, false}}},
	}
	for _, c := range cases {
		b := ChangeCardMatchValue(&c.card, c.side, c.newMatch)
		if c.card != c.expected {
			t.Errorf("Failed, got %s, expected %s", c.card, c.expected)
		}
		if !b {
			t.Errorf("Failed, boolean wrong")
		}
	}
}

func TestMatchCard(t *testing.T) {
	cases := []struct {
		firstcard, secondcard             Card
		executed                          bool
		boolexpected                      bool
		firstSide, secondSide             Side
		firstCardResult, secondCardResult Card
	}{
		{Card{SideValue{Blank, false}, SideValue{Blank, false}}, Card{SideValue{Blank, false}, SideValue{Six, false}}, true, true, First, First, Card{SideValue{Blank, true}, SideValue{Blank, false}}, Card{SideValue{Blank, true}, SideValue{Six, false}}},
	}
	for _, c := range cases {
		b, sfirst, ssecond := MatchCard(&c.firstcard, &c.secondcard, c.executed)
		if b != c.boolexpected || sfirst != c.firstSide || ssecond != c.secondSide {
			t.Errorf("Failed, got %t, %d, and %d, expected %t, %d, and %d", b, sfirst, ssecond, c.boolexpected, c.firstSide, c.secondSide)
		}
		if c.executed {
			if c.firstcard != c.firstCardResult || c.secondcard != c.secondCardResult {
				t.Errorf("Failed, got %s and %s, expected %s and %s", c.firstcard, c.secondcard, c.firstCardResult, c.secondCardResult)
			}
		}
	}
}
