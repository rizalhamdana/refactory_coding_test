package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	restaurantName := GetRestaurantName()
	cashierName := GetCashierName()
	date := GetDate()

	isAddItem := "y"
	itemMap := make(map[string]int64)
	count := 1
	for {
		name, price := AddItem(count)
		itemMap[name] = price

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Tambah Item? (y/n): ")
		input, _ := reader.ReadString('\n')
		isAddItem = strings.TrimSuffix(input, "\n")
		if isAddItem == "n" || isAddItem == "N" {
			PrintReceipt(restaurantName, cashierName, date, itemMap)
			break
		}

	}

}
