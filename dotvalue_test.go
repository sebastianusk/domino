// for testing the dotvalue
package domino

import (
	"fmt"
	"testing"
)

func TestValueToString(t *testing.T) {
	cases := []struct {
		sidevalue SideValue
		expected  string
	}{
		{SideValue{Blank, true}, "Blank and matched"},
		{SideValue{One, false}, "1 and not matched"},
		{SideValue{Six, false}, "6 and not matched"},
		{SideValue{Four, true}, "4 and matched"},
	}
	for _, c := range cases {
		result := fmt.Sprintf("%s", c.sidevalue)
		if result != c.expected {
			t.Errorf("Result wrong, wanted %s, got %s", c.expected, result)
		}
	}
}

func TestCheckIfMatched(t *testing.T) {
	cases := []struct {
		first, second SideValue
		expected      bool
	}{
		{SideValue{Blank, true}, SideValue{Blank, true}, false},
		{SideValue{Blank, false}, SideValue{Blank, true}, false},
		{SideValue{Blank, true}, SideValue{Blank, false}, false},
		{SideValue{Blank, false}, SideValue{Blank, false}, true},
		{SideValue{One, false}, SideValue{Five, false}, false},
		{SideValue{Three, false}, SideValue{Three, false}, true},
	}
	for _, c := range cases {
		result := CheckIfAbleToMatch(c.first, c.second)
		if result != c.expected {
			t.Errorf("Result wrong, first: %s, second: %s, result: %t, expected: %t", c.first, c.second, result, c.expected)
		}
	}
}

func TestChangeMatched(t *testing.T) {
	cases := []struct {
		test     SideValue
		newMatch bool
		expected SideValue
	}{
		{SideValue{One, true}, true, SideValue{One, true}},
		{SideValue{One, true}, false, SideValue{One, false}},
		{SideValue{One, false}, true, SideValue{One, true}},
		{SideValue{One, false}, false, SideValue{One, false}},
	}
	for _, c := range cases {
		b := c.test.ChangeMatched(c.newMatch)
		if c.test != c.expected {
			t.Errorf("Result Wrong, tested: %s, new value: %t, expected: %s", c.test, c.newMatch, c.expected)
		}
		if !b {
			t.Errorf("Result boolean failed, got %t, expected %t", b, true)
		}
	}
}
