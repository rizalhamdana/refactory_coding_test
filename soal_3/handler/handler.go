package handler

import (
	"fmt"

	"github.com/rizalhamdana/refactory_test/soal_3/usecase"
)

type Handler struct {
	usecase.Soal3Usecase
}

func NewHandler(usecase usecase.Soal3Usecase) *Handler {
	return &Handler{usecase}
}

func (h *Handler) RunProgram() {
	h.Soal3Usecase.FindItemsInRooms("meeting room")
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal3Usecase.FindItemBasedOnType("electronic")
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal3Usecase.FindItemBasedOnType("furniture")
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal3Usecase.FindItemPurchasedAtDate("2020-01-16")
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal3Usecase.FindItemByColor("brown")
	fmt.Println("")
	fmt.Println("======================\n")

}
