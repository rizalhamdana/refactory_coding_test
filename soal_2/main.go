package main

import (
	"fmt"
	"io/ioutil"

	"github.com/rizalhamdana/refactory_test/soal_2/handler"
	"github.com/rizalhamdana/refactory_test/soal_2/repo"
	"github.com/rizalhamdana/refactory_test/soal_2/usecase"
)

func main() {
	file, err := ioutil.ReadFile("resources/profile_list.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("SUCCESSFULLY LOAD DATA FROM PROFILE LIST JSON \n\n====================== \n\n")

	repository := repo.NewSoal2RepositoryImpl(file)
	usecase := usecase.NewUsecase(repository)
	handler := handler.NewHandler(usecase)

	handler.RunProgram()

}
