package keyboard

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
