package letters

import (
	"fmt"
	"strconv"
)

//Given a non-empty string, encode it using the following algorithm: for every letter, strip the consecutive repetitions, if any, and add a number with the number of times it appears consecutively, including the first time.
//
//
//Assume the following:
//1. Letters will have a maximum of 9 repetitions
//2. Letters are case sensitive (`A` and `a` are different characters)
//
//Some examples:
//
//“Goooooooooddddd” => “G1o9d5”
//“Yes” => “Y1e1s1”
//“noooo” => “n1o4”
//“Oooo” => “O1o3”
//“anagram” => “a1n1a1g1r1a1m1”
//
//Given the same assumptions, decode the string encoded. Examples:
//
//“a1n1a1g1r1a1m1” => “anagram”
//"y1e3s1" => “yeees”
//"Yuppy" => “Y1u1p2y1”

func Decode(input string) (res string) {
	l := len(input)

	for i := 0; i < l; i += 2 {
		res = fmt.Sprintf("%s%s", res, replaceNumberWithChar(string([]rune(input)[i]), string([]rune(input)[i+1])))
	}
	return res
}

func replaceNumberWithChar(char, numberString string) (res string) {
	n, err := strconv.Atoi(numberString)
	if err != nil {
		panic("wooo")
	}
	for i := 0; i < n; i++ {
		res = fmt.Sprintf("%s%s", res, char)
	}
	return res
}
