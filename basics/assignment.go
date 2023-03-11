package main
import "fmt"

func check(n int) {
	arr := [] int{};
	for i := 0; i < n; i++ {
		arr = append(arr, i);
	}
	 for _, val := range arr {
		if val % 2 == 0 {
			fmt.Println("Even");
		} else {
			fmt.Println("Odd");
		}
	}
}

func main() {
	check(10);
}