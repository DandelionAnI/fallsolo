package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func input() {
	//输入n（不定）个数组
	// a, b := 0, 0
	// for {
	// 	_, err := fmt.Scan(&a, &b)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(a + b)
	// }

	// 输入n个字符串输出
	// n := 0
	// fmt.Scanln(&n)
	// str := make([]string, n)
	// for i := range str {
	//     fmt.Scan(&str[i])
	// }
	// sort.Strings(str)
	// for i,v := range str {
	//     fmt.Print(v)
	//     if i!=len(str)-1 {
	//         fmt.Print(" ")
	//     }
	// }

	str := ""
	for {
		_, err := fmt.Scan(&str)
		if err == io.EOF {
			break
		}
		strs := strings.Split(str, ",")
		sort.Strings(strs)
		for i, v := range strs {
			fmt.Print(v)
			if i != len(strs) {
				fmt.Print(",")
			} else {
				fmt.Println()
			}
		}
	}
}

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
//     "strconv"
// )

// func main() {
//     sc := bufio.NewScanner(os.Stdin)
//     for sc.Scan() {
//         str := sc.Text()
//         strs := strings.Split(str," ")
//         a,_ := strconv.Atoi(strs[0])
//         b,_ := strconv.Atoi(strs[1])
//         fmt.Println(a+b)
//     }
// }
