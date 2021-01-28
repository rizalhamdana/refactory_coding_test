package main

import (
	"fmt"
	"io/ioutil"

	"github.com/rizalhamdana/refactory_test/soal_3/handler"
	"github.com/rizalhamdana/refactory_test/soal_3/repo"
	"github.com/rizalhamdana/refactory_test/soal_3/usecase"
)

func main() {
	file, err := ioutil.ReadFile("resources/inventory_list.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("SUCCESSFULLY LOAD DATA FROM INVETORY LIST JSON \n\n====================== \n\n")

	repository := repo.NewSoal3RepositoryImpl(file)
	usecase := usecase.NewUsecase(repository)
	handler := handler.NewHandler(usecase)

	handler.RunProgram()

}
