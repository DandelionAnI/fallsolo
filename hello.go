package main

import "fmt"

func main() {
	list := []string{"haha", "xixi"}

	pri := fmt.Sprintf("%q", list)

	fmt.Println("%s" + pri)
}
