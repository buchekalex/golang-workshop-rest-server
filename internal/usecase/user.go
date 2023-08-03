package usecase

import (
	"context"
	"go-mongo-crud-rest-api/internal/entitiy"
)

type UserUseCase struct {
	repository  Repository
	searchIndex SearchIndex
}

func (u UserUseCase) FindUser(ctx context.Context, email string) (*entitiy.User, error) {
	// TODO: implement me
	return nil, nil
}

func (u UserUseCase) GetUser(ctx context.Context, email string) (*entitiy.User, error) {
	// TODO: implement me
	return nil, nil
}

func (u UserUseCase) CreateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error) {
	// TODO: implement me
	return nil, nil
}

func (u UserUseCase) UpdateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error) {
	// TODO: implement me
	return nil, nil
}

func (u UserUseCase) DeleteUser(ctx context.Context, email string) error {
	// TODO: implement me
	return nil
}

func NewUserUseCase(repository Repository,
	elasticRepository SearchIndex) *UserUseCase {
	return &UserUseCase{
		repository:  repository,
		searchIndex: elasticRepository,
	}
}
