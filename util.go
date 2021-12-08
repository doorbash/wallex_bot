package main

import (
	"fmt"
	"strconv"
	"strings"

	thousands "github.com/floscodes/golang-thousands"
)

func ParsePrice(price string, precision int) string {
	i := strings.Index(price, ".")
	if i == -1 {
		return price
	}
	a, err := strconv.Atoi(price[0:i])
	if err != nil {
		return ""
	}
	l, err := thousands.Separate(a, "en")
	if err != nil {
		return ""
	}
	if precision == 0 {
		return l
	}
	return fmt.Sprintf("%s.%s", l, price[i+1:i+1+precision])
}
