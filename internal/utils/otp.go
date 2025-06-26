package utils

import "crypto/rand"

func GenerateOTP(length int) (string, error) {
	const digits = "0123456789"
	b := make([]byte, length)

	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		b[i] = digits[int(b[i])%len(digits)]
	}

	return string(b), nil
}
