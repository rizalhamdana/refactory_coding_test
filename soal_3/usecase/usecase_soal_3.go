package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/rizalhamdana/refactory_test/soal_3/repo"
)

type Soal3UsecaseImpl struct {
	repo repo.Soal3Repository
}

func NewUsecase(repo *repo.Soal3RepoImpl) *Soal3UsecaseImpl {
	return &Soal3UsecaseImpl{repo: repo}
}

func (u *Soal3UsecaseImpl) FindItemsInRooms(roomName string) {

	allItems, err := u.repo.GetAllItemsFromJson()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	printFormat := "ID: %d - Name: %s - Type: %s\n"
	count := 0
	fmt.Printf("Items in %s: \n", strings.Title(roomName))

	for _, item := range *allItems {
		placementNameLowerCase := strings.ToLower(item.Placement.Name)
		if strings.Contains(placementNameLowerCase, roomName) {
			fmt.Printf(printFormat, item.InventoryID, item.Name, item.Type)
			count++
		}
	}

	fmt.Printf("\nCount: %d \n", count)

}

func (u *Soal3UsecaseImpl) FindItemBasedOnType(typeCategory string) {
	allItems, err := u.repo.GetAllItemsFromJson()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	printFormat := "ID: %d - Name: %s - Type: %s\n"
	count := 0

	fmt.Printf(strings.Title(typeCategory) + " Items: \n")

	for _, item := range *allItems {
		placementNameLowerCase := strings.ToLower(item.Type)
		if placementNameLowerCase == typeCategory {
			fmt.Printf(printFormat, item.InventoryID, item.Name, item.Type)
			count++
		}
	}

	fmt.Printf("\nCount: %d \n", count)
}

func (u *Soal3UsecaseImpl) FindItemPurchasedAtDate(date string) {
	allItems, err := u.repo.GetAllItemsFromJson()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	printFormat := "ID: %d - Name: %s - Type: %s - Purchased At (YYYY-mm-dd): %s\n"
	count := 0

	fmt.Printf("Items Purchased At: %s\n", date)

	for _, item := range *allItems {
		dateTimeParams, err := time.Parse("2006-01-02", date)
		paramsDay := dateTimeParams.Day()
		paramsMonth := dateTimeParams.Month()
		paramsYear := dateTimeParams.Year()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		purchasedDateItem := time.Unix(int64(item.PurchasedAt), 0)
		if purchasedDateItem.Year() == paramsYear && purchasedDateItem.Month() == paramsMonth && purchasedDateItem.Day() == paramsDay {
			datePurchasedString := fmt.Sprintf("%d-%d-%d", purchasedDateItem.Year(), purchasedDateItem.Month(), purchasedDateItem.Day())
			fmt.Printf(printFormat, item.InventoryID, item.Name, item.Type, datePurchasedString)
			count++
		}
	}
	fmt.Printf("\nCount: %d \n", count)

}

func (u *Soal3UsecaseImpl) FindItemByColor(color string) {
	allItems, err := u.repo.GetAllItemsFromJson()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	printFormat := "ID: %d - Name: %s - Type: %s\n"
	count := 0

	fmt.Printf(strings.Title(color) + " Color Items: \n")

	for _, item := range *allItems {
		itemNameLower := strings.ToLower(item.Name)
		if strings.Contains(itemNameLower, strings.ToLower(color)) {
			fmt.Printf(printFormat, item.InventoryID, item.Name, item.Type)
			count++
			continue
		}

		for _, tag := range item.Tags {
			if strings.Contains(strings.ToLower(tag), strings.ToLower(color)) {
				fmt.Printf(printFormat, item.InventoryID, item.Name, item.Type)
				count++
			}
		}
	}

	fmt.Printf("\nCount: %d \n", count)
}
