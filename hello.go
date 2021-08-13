package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	Code()
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

func main() {
	Code()
}

func Code() {
	nums := []int{6, 3, 8, 4, 5, 23, 6, 8, 3, 2, 7}
	//正序
	sort.Ints(nums)
	fmt.Println(nums)
	//SearchInts(二分查找)，要求已经升序排序
	pos := sort.SearchInts(nums, 5)
	fmt.Println(pos)
	//升序条件自定义查找条件
	pos = sort.Search(len(nums), func(i int) bool { return nums[i] > 7 })
	fmt.Println(pos)
	//降序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)

	// sort.Ints()
	// sort.Float64s()
	// sort.Strings()
	// sort.Slice()
	// sort.SearchInts()
	// sort.SearchFloat64s()
	// sort.SearchStrings()
}
