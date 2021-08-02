package main

import (
	"fmt"
	"os"
	//. "solution/solution"
)

func main() {
	num := os.Args[1]
	topic := os.Args[2]

	f, err := os.Create("./solution/" + num + topic + ".go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		f.Write([]byte("package solution\n"))
		f.Write([]byte("\n"))
		f.Write([]byte("func " + topic + "() {\n"))
		f.Write([]byte("}\n"))
	}
	defer f.Close()
}
