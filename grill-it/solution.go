package main

import (
	"strconv"
	"strings"
)

func Grill(message string, code int) string {
	// convert code to binary representation
	binary := strconv.FormatInt(int64(code), 2)

	// pad with zeros on the left to match message length
	for len(binary) < len(message) {
		binary = "0" + binary
	}

	// extract characters where binary digit is '1'
	var result strings.Builder
	for i := 0; i < len(message) && i < len(binary); i++ {
		if binary[i] == '1' {
			result.WriteByte(message[i])
		}
	}

	return result.String()
}