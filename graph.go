package main
import (
	"fmt"
)
type edge struct {
	d int
	cost int
}	

type gnode []edge

func main() {
	n := 5
	edges := [][] int {
		{1, 2, 3},
		{2, 3, 3},
		{2, 4, 3},
		{0, 1, 3},
	}
	g := make([] gnode, n);
	
	for _, e := range edges {
		s := e[0]
		d := e[1]
		c := e[2]
		g[s] = append(g[s], edge{d, c});
	}

	fmt.Printf("%+v", g)
}