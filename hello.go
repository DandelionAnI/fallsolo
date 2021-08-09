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

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	//a := 0
// 	//fmt.Scan(&a)
// 	//fmt.Printf("%d\n", a)

// 	a,b := []int{1,1,2,3,4,5,6,7,8,9},[]int{4,4,4,5,5,5,6,6,6,10}
// 	// a, b := []int{1, 2, 3}, []int{1, 2, 3}
// 	l, r := len(a)-1, len(b)-1
// 	sort.Ints(a)
// 	sort.Ints(b)
// 	fmt.Println(a, b)
// 	cnt := 0
// 	for cnt <= 0 && l != -1 {
// 		if a[l] > b[r] {
// 			cnt++
// 		}  else if a[l] < b[r] {
// 			cnt--
// 		}
// 		l--
// 		r--
// 	}
// 	if l == -1 && cnt <= 0 {
// 		fmt.Println(-1)
// 	} else {
// 		ans := len(a) - l -1
// 		fmt.Println(ans)
// 	}
// }
