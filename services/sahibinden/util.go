package sahibinden

import (
	"strconv"
	"strings"
)

func ParseInteger(price string) int {
	integerPrice, err := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(ClearText(price),
		" ", ""), ".", ""), "TL", ""))
	if err != nil {
		panic(err)
	}
	return integerPrice
}

func ClearText(text string) string {
	return strings.TrimSpace(strings.ReplaceAll(text, "\n", ""))
}
