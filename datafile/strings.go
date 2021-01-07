//package datafiles read data from files

package datafile

import (
	"bufio"
	"os"
)

// GetString function reads float data from a file
func GetString(fileName string) ([]string, error) {
	var words []string
	file, err := os.Open(fileName)
	if err != nil {
		return words, err
	}
	scanner := bufio.NewScanner(file)
	//i := 0
	var tmpStr string
	for scanner.Scan() {
		tmpStr = scanner.Text()
		words = append(words, tmpStr)

	}
	err = file.Close()
	if err != nil {
		return words, err
	}
	if scanner.Err() != nil {
		return words, scanner.Err()
	}
	return words, nil
}
