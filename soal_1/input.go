package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func AddItem(count int) (string, int64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Masukkan Nama Item-%d: ", count)
	itemName, _ := reader.ReadString('\n')

	fmt.Print("Masukkan Harga Item (Hanya Angka): ")
	itemPriceString, _ := reader.ReadString('\n')

	itemPrice, err := strconv.ParseInt(strings.TrimSuffix(itemPriceString, "\n"), 10, 64)
	if err != nil {
		return "", 0
	}
	return strings.TrimSuffix(itemName, "\n"), itemPrice
}

func GetRestaurantName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama Warung: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}

func GetCashierName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama Kasir: ")
	input, _ := reader.ReadString('\n')

	return strings.TrimSuffix(input, "\n")
}

func GetDate() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Tanggal (dd-mm-YYYY): ")
	input, _ := reader.ReadString('\n')
	_, err := time.Parse("02-01-2006", strings.TrimSuffix(input, "\n"))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Invalid date format")
		return ""
	}

	return strings.TrimSuffix(input, "\n")
}
