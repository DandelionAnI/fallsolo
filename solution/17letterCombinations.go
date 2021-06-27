package solution

//回溯
var ans []string
var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	ans = []string{}
	if digits == "" {
		return ans
	}
	letterFunc("", digits)
	return ans
}

func letterFunc(res string, digits string) {
	if digits == "" {
		ans = append(ans, res)
		return
	}

	num := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[num]); i++ {
		res += dict[num][i]
		letterFunc(res, digits)
		res = res[0 : len(res)-1]
	}
}
