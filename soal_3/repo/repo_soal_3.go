package repo

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/rizalhamdana/refactory_test/soal_3/model"
)

type Soal3RepoImpl struct {
	jsonFiles []byte
}

func NewSoal3RepositoryImpl(jsonFiles []byte) *Soal3RepoImpl {
	return &Soal3RepoImpl{jsonFiles: jsonFiles}
}

func (r *Soal3RepoImpl) GetAllItemsFromJson() (*[]model.Inventory, error) {
	allItems := []model.Inventory{}

	err := json.Unmarshal(r.jsonFiles, &allItems)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Error when get all items from json")
	}
	return &allItems, nil
}
