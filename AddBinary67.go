package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {

	var result []string

	var add uint8 = 0
	i, j:= len(a) -1, len(b) - 1
	for ; i >=0 && j >= 0; i, j = i-1,j-1 {

		x := a[i] - '0'
		y := b[j] - '0'
		r := (x + y + add) % 2
		add = (x + y + add) / 2
		result = append([]string{strconv.Itoa(int(r))}, result...)
	}
	for ;i >= 0; i--{
		x := a[i] - '0'
		r := (x + add) % 2
		add = (x  + add) / 2
		result = append([]string{strconv.Itoa(int(r))}, result...)
	}
	for ; j >= 0; j-- {
		y := b[j] - '0'
		r := ( y + add) % 2
		add = (y + add) / 2
		result = append([]string{strconv.Itoa(int(r))}, result...)
	}
	if add > 0 {
		result = append([]string{"1"}, result...)
	}
	return strings.Join(result, "")
}

func main() {
	fmt.Println(addBinary("1010","1011"))
}
