package parser

import "strconv"

type Error struct {
	line    int
	col     int
	message string
}

func NewError(line int, col int, message string) *Error {
	return &Error{line, col, message}
}

func (error *Error) Error() string {
	return "Line " + strconv.Itoa(error.line) + ", Col " + strconv.Itoa(error.col) + error.message
}
