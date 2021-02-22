package main

import "testing"

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMillerRabin(t *testing.T) {
	n1 := int64(592701729979) // true
	if !MillerRabin(n1) {
		t.Errorf("got false, expected true for %d", n1)
	}
	n2 := int64(100) // true
	if !MillerRabin(n2) {
		t.Errorf("got false, expected true for %d", n2)
	}

	n3 := int64(101) // false
	if MillerRabin(n3) {
		t.Errorf("got true, expected false for %d", n3)
	}
}

// Note that as tests go, this is especially wonky! It will sometimes fail since it has a random number
func TestMillerRabin_OddComposites(t *testing.T) {
	n2 := int64(55) // true
	if !MillerRabin(n2) {
		t.Errorf("got false, expected true for %d", n2)
	}
	n3 := int64(49) // true
	if !MillerRabin(n3) {
		t.Errorf("got false, expected true for %d", n3)
	}
	n4 := int64(75) // true
	if !MillerRabin(n4) {
		t.Errorf("got false, expected true for %d", n4)
	}
	n5 := int64(87) // true
	if !MillerRabin(n5) {
		t.Errorf("got false, expected true for %d", n5)
	}
	n6 := int64(91) // true
	if !MillerRabin(n6) {
		t.Errorf("got false, expected true for %d", n6)
	}
	n7 := int64(93) // true
	if !MillerRabin(n7) {
		t.Errorf("got false, expected true for %d", n7)
	}
}

