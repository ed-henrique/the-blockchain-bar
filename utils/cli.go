package utils

import (
	"errors"
	"fmt"
	"os"
)

var (
	IncorrectUsageErr = errors.New("incorrect usage")
)

func Err(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
