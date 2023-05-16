package service

import (
	"context"
	"golangeko/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, CategoryId int)
	FindById(ctx context.Context, CategoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
