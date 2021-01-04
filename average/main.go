package main

import "fmt"

func main() {
	var weeklyMeatConsumption = [7]float64{45.4, 39.0, 56.7, 42.2, 60.2, 66.7, 49.5}
	var sum float64 = 0
	var noOfSamples = 0.0
	for _, value := range weeklyMeatConsumption {
		sum += value
	}
	noOfSamples = float64(len(weeklyMeatConsumption))
	var avgConsumption = sum / noOfSamples
	fmt.Println("Sum of weekly meat conusmption is ", sum)
	fmt.Printf("Average weekly consumption is %.2f", avgConsumption)
}
