package randgenerator

import (
	"accounting_system/internal/repositories"
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
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a local random generator
	length := r.Intn(64) + 1                             // Length will be between 1 and 64
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = letters[r.Intn(len(letters))]
	}

	return string(result)
}

func GenerateUniqeCode[T any](repo *repositories.Repositories, columnName string) string {
	code := GenerateRandomCode()
	for {
		exist := repositories.FindRecord[T](repo, code, columnName)

		if exist {
			code = GenerateRandomCode()
		} else {
			break
		}
	}

	return code
}

func GenerateUniqeTitle[T any](repo *repositories.Repositories) string {
	title := GenerateRandomTitle()
	for {
		exist := repositories.FindRecord[T](repo, title, "title")

		if exist {
			title = GenerateRandomTitle()
		} else {
			break
		}
	}

	return title
}
