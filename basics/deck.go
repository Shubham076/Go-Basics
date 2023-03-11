package main
import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"math/rand"
	"time"
)


type deck []string;		

func newDeck() deck {
	cards := deck{}
	names := []string{"Heart", "Diamond", "Club", "Spade"};
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", 
	"Nine", "Ten", "King", "Queen", "Jack"};

	for _, c := range names {
		for _, v := range values {
			cards = append(cards, v + " of " + c)
		}
	}
	return cards;
}
func (d deck) print() {
	for idx, card := range d {
		fmt.Println(idx, card);
	}
}

func deal(d deck, size int) (deck, deck){
	return d[:size], d[size:];	
}

func (d deck) toString() string {
	return strings.Join([]string (d), ",")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename);
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",");
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		new_idx := r.Intn(len(d) - 1);
		d[i], d[new_idx] = d[new_idx], d[i]
	}
}