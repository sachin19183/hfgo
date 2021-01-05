//Package keyboard abstracts all types of keyboard inputs that may be required for a program
// The input types supported are : string type.
// one variable can be read with a single function call
// Function will return with 0 and an error object if the input is invalid
package keyboard

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

//GetString reads the input from the keyboard
// It returns with 0 and an error object if the input is invalid
// It returns a float value if no error is detected
func GetString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err1 := reader.ReadString('\n')
	if err1 != nil {
		return "", errors.New("Please enter a valid value")
	}
	input = strings.TrimSpace(input)
	return input, nil
}
