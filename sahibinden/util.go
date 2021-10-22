package sahibinden

import (
	"fmt"
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

func GetSellerPageUrl(seller string, pageSize int) string {
	return fmt.Sprintf("https://%s.sahibinden.com/vasita?pagingOffset=0&pagingSize=%d", seller, pageSize)
}
