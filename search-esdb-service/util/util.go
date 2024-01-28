package util

import (
	"errors"
	"strconv"
)

// create a function to find the first `-` from the back then convert the string after that `-` to int then decrease by 1 e.g. `xxxxx-45` return `xxxxx-44`
func DecreaseIndexForSearchByIndex(str string) (string, error) {
	// check is string in that format
	// e.g. `xxxxx-45`
	if len(str) < 7 {
		return str, errors.New("string is not in the format of `xxxxx-yy`")
	}

	// find the first `-` from the back
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '-' {
			// convert the string after that `-` to int
			num, err := strconv.Atoi(str[i+1:])
			if err != nil {
				return str, err
			}
			// decrease by 1
			num--
			// return `xxxxx-44`
			return str[:i+1] + strconv.Itoa(num), nil
		}
	}
	return str, errors.New("string is not in the format of `xxxxx-yy`")
}
