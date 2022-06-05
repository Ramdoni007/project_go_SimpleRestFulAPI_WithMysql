package service

import (
	"Ramdoni007/Learn-Golang-RestFullAPI/helper"
	"Ramdoni007/Learn-Golang-RestFullAPI/model/web "
	"Ramdoni007/Learn-Golang-RestFullAPI/repository"
	"context"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryRespone {
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
		}else {
			tx.Commit()
		}
	}()

}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryRespone {

}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	panic("implement me")
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryRespone {
	panic("implement me")
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryRespone {
	panic("implement me")
}

