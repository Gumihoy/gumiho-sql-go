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
	return "line " + strconv.Itoa(error.line) + ", col " + strconv.Itoa(error.col) + error.message
}
