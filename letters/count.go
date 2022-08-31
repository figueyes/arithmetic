package letters

import "fmt"

func CounterLetter(input string) (res string) {
	if input == "" {
		return "not-valid"
	}
	count := 0
	for i := 0; i < len(input); i++ {
		if res == "" {
			count = 1
			res = fmt.Sprintf("%s%s%d", res, string([]rune(input)[i]), count)
			continue
		}

		if string([]rune(input)[i]) == " " {
			continue
		}

		if string([]rune(input)[i]) == string([]rune(input)[i-1]) {
			count += 1
			res = fmt.Sprintf("%s%d", string([]rune(res)[:len(res)-1]), count)
		} else {
			count = 1
			res = fmt.Sprintf("%s%s%d", res, string([]rune(input)[i]), count)
		}
	}
	return res
}
