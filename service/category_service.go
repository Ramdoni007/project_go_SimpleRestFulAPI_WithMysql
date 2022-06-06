package service

import (
	"Ramdoni007/Learn-Golang-RestFullAPI/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context,request web.CategoryCreateRequest)web.CategoryRespone
	Update(ctx context.Context,request web.CategoryUpdateRequest)web.CategoryRespone
	Delete(ctx context.Context,categoryId int)
	FindById(ctx context.Context,categoryId int)web.CategoryRespone
	FindAll(ctx context.Context)[]web.CategoryRespone
}
