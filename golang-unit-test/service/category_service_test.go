package service

import (
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
	categoryService    = CategoryService{Repository: categoryRepository}
)

func TestCategoryService_Get(t *testing.T) {
	categoryRepository.Mock.On("FindByID", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
