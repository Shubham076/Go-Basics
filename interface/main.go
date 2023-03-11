package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() { 
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Something went wrong", err)
		os.Exit(1)
	}
	// fmt.Println(res)
	// slice := make([] byte, 99999);
	// res.Body.Read(slice)
	// fmt.Println(string(slice))

	//shortcut
	io.Copy(os.Stdout, res.Body)
}