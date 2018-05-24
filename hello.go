package main

import "fmt"

func main()  {
	fmt.Println("hello world")
	var name  = []string{"123","456"}
	for s:=range name {
		fmt.Println(name[s])
	}
}