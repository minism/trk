package util

import "strings"

func SplitStringUpToLastHyphen(input string) (string, string) {
	index := strings.LastIndex(input, "-")
	if index == -1 {
		return "", input
	}
	return input[:index], input[index+1:]
}
