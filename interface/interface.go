package main
import "fmt"

/*
with this we are telling go that any type that defines or 
matching the description of the function inside bot interface
is also  a type of bot
*/
type bot interface {
	generateString() string
}

type englishbot struct{}
type spanishbot struct{}

func (e englishbot) generateString() string {
	return "Hello"
}

func (s spanishbot) generateString() string {
	return "Hola"
}

func printString(b bot){
	fmt.Println(b.generateString())
}

func main() {
	b := englishbot{};
	s := spanishbot{}

	printString(b);
	printString(s);
}