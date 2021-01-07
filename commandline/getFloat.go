//Package commandline contains method to parse a command line argument and convert it into a float64 slice
package commandline

import (
	"strconv"
	"strings"
)

//Function GetFloatCmd takes command line input and converts it into a float64 slcie and retunrs it for further usage
func GetFloatCmd(argList []string) ([]float64, error) {
	var floatSlice []float64
	for _, value := range argList {
		value = strings.TrimSpace(value)
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return floatSlice, err
		}
		floatSlice = append(floatSlice, floatVal)

	}
	return floatSlice, nil
}
