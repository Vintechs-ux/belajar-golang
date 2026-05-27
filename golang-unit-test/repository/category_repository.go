package repository

import "belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindByID(id string) *entity.Category
}
