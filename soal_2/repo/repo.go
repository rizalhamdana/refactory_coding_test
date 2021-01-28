package repo

import "github.com/rizalhamdana/refactory_test/soal_2/model"

type Soal2Repository interface {
	GetAllUsersFromJson() (*[]model.User, error)
}
