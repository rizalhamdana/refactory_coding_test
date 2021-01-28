package repo

import "github.com/rizalhamdana/refactory_test/soal_3/model"

type Soal3Repository interface {
	GetAllItemsFromJson() (*[]model.Inventory, error)
}
