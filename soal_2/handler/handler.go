package handler

import (
	"fmt"

	"github.com/rizalhamdana/refactory_test/soal_2/usecase"
)

type Handler struct {
	usecase.Soal2Usecase
}

func NewHandler(usecase usecase.Soal2Usecase) *Handler {
	return &Handler{usecase}
}

func (h *Handler) RunProgram() {
	h.Soal2Usecase.FindUsersWhoNoPhoneNumbers()
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal2Usecase.FindUsersWhoHaveArticles()
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal2Usecase.FindUsersWhoHaveAnnisInTheirName()
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal2Usecase.FindUsersWhoBornOn1996()
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal2Usecase.FindUsersWhoHaveArticlesIn2020()
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal2Usecase.FindArticlesContainTips()
	fmt.Println("")
	fmt.Println("======================\n")

	h.Soal2Usecase.FindArticlesPublishedBeforeAugust2019()
	fmt.Println("")
	fmt.Println("======================\n")
}
