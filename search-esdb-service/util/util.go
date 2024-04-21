package util

import (
	"strconv"
	"strings"
)

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// Remove stopwords from token array
// arr => tokens
// slice => stopwords
func RemoveSliceFromArrays(arr []string, slice []string) []string {
	var result []string
	for _, a := range arr {
		if Contains(slice, a) {
			continue
		}
		result = append(result, a)
	}

	if len(result) == 0 {
		return arr
	}

	return result
}

func ConvertStringToFloat32Arrays(arr string) ([]float32, error) {
	if arr == "[]" {
		// Empty array case
		return nil, nil
	}
	var result []float32
	values := strings.Split(arr, ",")
	for _, v := range values {
		trimmedV := strings.TrimSpace(v)
		trimmedV = strings.Trim(trimmedV, "[]")
		floatV, err := strconv.ParseFloat(trimmedV, 32)
		if err != nil {
			return nil, err
		}
		result = append(result, float32(floatV))
	}

	return result, nil
}
