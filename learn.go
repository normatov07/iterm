package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	msg string
}

func (e CustomError) Error() string {
	return e.msg
}

func main() {
	err := CustomError{msg: "custom error"}

	var cErr CustomError

	if errors.As(err, &cErr) {
		fmt.Println(cErr.Error())
	}

	wrappedErr := fmt.Errorf("wrpp err: %w", err)

	fmt.Println(wrappedErr)
}
