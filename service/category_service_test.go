package service

import (
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{
	Mock: mock.Mock{},
}

var categoryService = CategoryService{
	Repository: categoryRepository,
}

func TestCategoryService_GetSuccess(t *testing.T) {
	result := entity.Category{
		Id:   "1",
		Name: "Elektronik",
	}

	categoryRepository.Mock.On("FindById", "1").Return(result)
	category, err := categoryService.Get("1")

	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, &result, category)
}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "2").Return(nil)
	category, err := categoryService.Get("2")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}
