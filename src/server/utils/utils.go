package utils

import (
	"log"
	"strconv"
)

// Atoi converts string to int.
func Atoi(v string) int {
	s, err := strconv.Atoi(v)
	if err != nil {
		log.Printf(err.Error())
	}
	return s
}

// Itoa converts int to string.
func Itoa(v int) string {
	return strconv.Itoa(v)
}
