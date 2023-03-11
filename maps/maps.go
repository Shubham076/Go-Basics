package main
import "fmt"
func main() {
	// var m map[string]string;
	m := make(map[string]string)

	x := map[string]string{
		"name": "Shubham",
		"pin": "abc",
	}

	//adding the key
	m["white"] = "#ffff"
	color := "white";
	m[color] = "#212121"
	fmt.Println(x)
	fmt.Println(m)

	//deleting the key
	delete(m, "white")

	//iterating over map
	z := map[int]string {
		10: "abc",
		12: "def",
		13: "efg",
	}

	for k,v := range z {
		fmt.Println(k, v);
	}
	fmt.Println(m)
	
	//check the existence of key
	_, isPresent := z[100];
	fmt.Println("Key present or not: ", isPresent);
}