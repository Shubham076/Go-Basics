package main
import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 52 {
		t.Error("Expected 52 but got", len(d))
	}

	if d[0] != "Ace of Heart" {
		t.Error("Expected Ace of Heart but got", d[0])	
	}
}