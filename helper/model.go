package helper

import (
	"Ramdoni007/Learn-Golang-RestFullAPI/model/domain"
	"Ramdoni007/Learn-Golang-RestFullAPI/model/web"
)

func ToCategoryRespone(category domain.Category) web.CategoryRespone {
	return web.CategoryRespone{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryRespone {
	var categoryResponses []web.CategoryRespone

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryRespone(category))
	}
	return categoryResponses
}
