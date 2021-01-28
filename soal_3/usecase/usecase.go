package usecase

type Soal3Usecase interface {
	FindItemsInRooms(roomName string)
	FindItemBasedOnType(typeCategory string)
	FindItemPurchasedAtDate(date string)
	FindItemByColor(color string)
}
