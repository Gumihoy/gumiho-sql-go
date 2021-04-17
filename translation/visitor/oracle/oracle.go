package oracle

import "strings"

func IsKeepDoubleQuote(text string) bool {
	return containLowerLetter(text) || isStartNum(text) || containSpecialCharacter(text)
}

/**
* 判断是否有小写字母
*/
func containLowerLetter(text string) bool {
	if text == "" || len(text) == 0 {
		return false
	}
	for c := range text {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}

/**
* 判断是否有小写字母
*/
func isStartNum(text string) bool {
	if text == "" || len(text) == 0 {
		return false
	}
	c := text[0]
	return c >= '0' && c <= '9'
}

var special_character = []rune{'<', '>', '（', '）', '(', ')', '{', '}', '[', ']', '!', '@', '%', '^', '&', '*', '/', '\\', '+', '-', '|', ':', '\'', '=', ' ', ';', '.', '\'', '?', '~'}
func containSpecialCharacter(text string) bool {
	if text == "" || len(text) == 0 {
		return false
	}
	for _, c := range special_character {
		contains := strings.ContainsRune(text, c)
		if contains {
			return true
		}
	}

	return false
}
