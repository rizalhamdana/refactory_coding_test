package repo

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/rizalhamdana/refactory_test/soal_2/model"
)

type Soal2RepoImpl struct {
	jsonFiles []byte
}

func NewSoal2RepositoryImpl(jsonFiles []byte) *Soal2RepoImpl {
	return &Soal2RepoImpl{jsonFiles: jsonFiles}
}

func (r *Soal2RepoImpl) GetAllUsersFromJson() (*[]model.User, error) {
	allUsers := []model.User{}

	err := json.Unmarshal(r.jsonFiles, &allUsers)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Error when get all users from json")
	}
	return &allUsers, nil
}
