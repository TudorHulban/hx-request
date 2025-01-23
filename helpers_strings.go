package hxrequest

import (
	"errors"
	"strconv"
	"strings"

	goerrors "github.com/TudorHulban/go-errors"
)

func getNumberFromParentheses(text string) (int64, error) {
	_, after, existsLeftP := strings.Cut(text, "(")
	if existsLeftP {
		before, _, existsRightP := strings.Cut(after, ")")
		if existsRightP {
			return strconv.ParseInt(before, 10, 64)
		}
	}

	return 0,
		goerrors.ErrInvalidInput{
			Caller:     "GetNumberFromParentheses",
			InputName:  "number",
			InputValue: text,
			Issue:      errors.New("invalid input"),
		}
}

func parseBool(value string) (bool, error) {
	switch value {
	case "1", "t", "T", "true", "TRUE", "True", "yes":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "no":
		return false, nil
	}

	return false,
		goerrors.ErrUnknownValue
}
