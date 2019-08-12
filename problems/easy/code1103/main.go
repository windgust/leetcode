package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	var candy, i int
	peopleCandy := make([]int, num_people, num_people)

	for candies > 0 {
		candy++
		if candy > candies {
			candy = candies
		}
		if i >= num_people {
			i = 0
		}
		peopleCandy[i] += candy
		candies -= candy
		i++
	}
	return peopleCandy
}

func main() {
	fmt.Println(distributeCandies(60, 4))
}
