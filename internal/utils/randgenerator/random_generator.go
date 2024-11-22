package randgenerator

import (
	"math/rand"
	"time"
)

func GenerateRandomCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := r.Intn(64) + 1
	numbers := "0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = numbers[r.Intn(len(numbers))]
	}

	return string(result)
}

func GenerateRandomTitle() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) 
	length := r.Intn(64) + 1                             
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = letters[r.Intn(len(letters))]
	}

	return string(result)
}
