//package datafiles read data from files

package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloat function reads float data from a file
func GetFloat(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(file)
	//i := 0
	var tmpNum float64
	for scanner.Scan() {
		tmpNum, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		} else {
			numbers = append(numbers, tmpNum)
		}
		//i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
