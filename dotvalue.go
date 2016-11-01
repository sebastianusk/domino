// To implement matching domino value
package domino

import (
	"fmt"
)

// value represent the number of dots
type DotValue int

// constant for Value, there are only 7 posibilities, from blank to six
const (
	Blank DotValue = iota
	One
	Two
	Three
	Four
	Five
	Six
)

// Side Value for implement each side of the card,
// there is the value and bool if it's already matched or not
type SideValue struct {
	dotvalue DotValue
	match    bool
}

// Stringer interface for DotValue
func (dv SideValue) String() string {
	var valued, matched string
	switch dv.dotvalue {
	case Blank:
		valued = "Blank"
	case One:
		valued = "1"
	case Two:
		valued = "2"
	case Three:
		valued = "3"
	case Four:
		valued = "4"
	case Five:
		valued = "5"
	case Six:
		valued = "6"
	}
	if dv.match {
		matched = "matched"
	} else {
		matched = "not matched"
	}
	return fmt.Sprintf("%s and %s", valued, matched)
}

// check wether two side value will be able to match or not
func CheckIfAbleToMatch(first, second SideValue) bool {
	if first.match || second.match {
		return false
	}
	if first.dotvalue != second.dotvalue {
		return false
	}
	return true
}

// change the status match of the side value
func (dv *SideValue) ChangeMatched(newMatch bool) bool {
	dv.match = newMatch
	return dv.match == newMatch
}
