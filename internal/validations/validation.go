package validations

import (
	"errors"
	"fmt"
)

func CheckEmpty(s string) error {
	if len(s) == 0 {
		return errors.New("empty field not allowed")
	} else {
		return nil
	}
}

func CheckMaxLength(s string, maxLen int) error {
	if len(s) > maxLen {
		return fmt.Errorf("field length is greater than max length witch is %d", maxLen)
	} else {
		return nil
	}
}
