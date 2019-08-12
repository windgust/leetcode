package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	d := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	dateArr := strings.Split(date, "-")
	year, _ := strconv.Atoi(dateArr[0])
	month, _ := strconv.Atoi(dateArr[1])
	day, _ := strconv.Atoi(dateArr[2])

	if isLeapYear(year) {
		d[2]++
	}

	var res int
	for i := 1; i < month; i++ {
		res += d[i]
	}
	res += day
	return res

}

func isLeapYear(y int) bool {
	return ((y%4 == 0 && y%100 != 0) || (y%400 == 0 && y%3200 != 0))
}

func main() {
	fmt.Println(dayOfYear("2004-03-01"))
}