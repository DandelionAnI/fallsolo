package main

// import (
// 	"fmt"
// 	"os"
// 	//. "solution/solution"
// )

// func main() {
// 	num := os.Args[1]
// 	topic := os.Args[2]

// 	f, err := os.Create("./solution/" + num + topic + ".go")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		f.Write([]byte("package solution\n"))
// 		f.Write([]byte("\n"))
// 		f.Write([]byte("func " + topic + "() {\n"))
// 		f.Write([]byte("}\n"))
// 	}
// 	defer f.Close()
// }

import (
	"fmt"
)

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)

	nums := []int{1, 2, 2, 3, 3, 5, 5, 7}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum ^= nums[i]
	}
	p := 0
	fmt.Println(sum)
	for (sum & (1 << p)) == 0 {
		fmt.Println(p)
		p++
	}
	nums1, nums2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if (nums[i] & (1 << p)) != 0 {
			nums1 ^= nums[i]
		} else {
			nums2 ^= nums[i]
		}
	}
	fmt.Println(nums1, nums2)
}
