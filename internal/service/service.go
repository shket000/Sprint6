package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// AutoDetectAndConvert определяет, является ли введённая строка текстом или кодом Морзе,
// и выполняет соответствующую конвертацию.
func AutoDetectAndConvert(input string) (string, error) {
	if isMorse(input) {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}

// isMorse проверяет, состоит ли строка только из символов '.', '-', ' ' и пробельных символов.
func isMorse(input string) bool {
	for _, r := range input {
		if r != '.' && r != '-' && r != ' ' &&
			r != '\t' && r != '\n' && r != '\r' {
			return false
		}
	}
	return true
}
