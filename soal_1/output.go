package main

import (
	"bytes"
	"fmt"
	"strings"
)

func PrintReceipt(restaurantName, cashierName, date string, items map[string]int64) {

	fmt.Println("")
	fmt.Println(centerString(restaurantName, 30))

	formatOutput("Nama Kasir:", cashierName, " ")
	formatOutput("Tanggal:", date, " ")
	fmt.Println("")
	fmt.Println(strings.Repeat("=", 30))

	var totalPrice int64 = 0

	for key, element := range items {
		formatOutput(key, fmt.Sprintf("Rp%d", element), ".")
		totalPrice += element
	}

	fmt.Println("")
	formatOutput("Total:", fmt.Sprintf("Rp%d", totalPrice), ".")
}

func centerString(str string, totalField int) string {

	strLen := len(str)
	spacesToPad := totalField - strLen
	var tmpSpaces float64
	var lrSpaces int

	tmpSpaces = float64(spacesToPad) / 2
	lrSpaces = int(tmpSpaces)

	buffer := bytes.NewBufferString("")

	spacesRemaining := totalField

	for i := 0; i < lrSpaces; i++ {
		buffer.WriteString(" ")
		spacesRemaining = spacesRemaining - 1
	}
	buffer.WriteString(str)
	spacesRemaining = spacesRemaining - strLen
	for i := spacesRemaining; i > 0; i-- {
		buffer.WriteString(" ")
	}

	return buffer.String()
}

func formatOutput(item, value string, separator string) {
	lenCharItem := len(item)
	if len(item) > 15 {
		item = item[:8] + "\n" + item[8:]
		lenCharItem = len(item[8:])
	}
	outputFormat := item + leftPad(value, separator, 30-(lenCharItem+len(value))) + "\n"

	fmt.Printf(outputFormat)

}

func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}
