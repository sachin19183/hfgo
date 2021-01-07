//Package average2 calculates average of float data type numbers
package average2

import "fmt"

//GetAverageFloat take variable number of float arguments and calculates their sum.IT returns the average of the variable float arguments
func GetAverageFloat(floatSeries ...float64) (float64, error) {
	sum := 0.0
	for _, value := range floatSeries {
		sum += value
	}
	average := sum / float64(len(floatSeries))
	fmt.Printf("Sum is %.2f", sum)
	return average, nil
}
