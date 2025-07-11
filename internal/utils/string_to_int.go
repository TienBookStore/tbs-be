package utils

import (
	"errors"
	"strconv"
)

func StringToUint(s string) (uint, error) {
	if s == "" {
		return 0, errors.New("input string is empty")
	}

	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	if value > uint64(^uint(0)) {
		return 0, errors.New("value exceeds maximum uint size")
	}

	return uint(value), nil
}
