//Package keyboard abstracts all types of keyboard inputs that may be required for a program
// The input types supported are : float64 type.
// one variable can be read with a single function call
// Function will return with 0 and an error object if the input is invalid
package keyboard

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//GetFloat reads the input from the keyboard and attempts to convert it into float value
// It returns with 0 and an error object if the input is invalid
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err1 := reader.ReadString('\n')
	if err1 != nil {
		return 0, errors.New("Please enter a valid value")
	}
	input = strings.TrimSpace(input)
	floatVal, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("%s is not a valid number", input)
	}
	return floatVal, nil
}
