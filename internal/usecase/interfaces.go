package usecase

import (
	"context"
	"go-mongo-crud-rest-api/internal/entitiy"
)

type SearchIndex interface {
	GetUser(ctx context.Context, email string) (*entitiy.User, error)
	CreateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error)
	UpdateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error)
	DeleteUser(ctx context.Context, email string) error
	FindUser(ctx context.Context, email string) (*entitiy.User, error)
}

type Repository interface {
	GetUser(ctx context.Context, email string) (*entitiy.User, error)
	CreateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error)
	UpdateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error)
	DeleteUser(ctx context.Context, email string) error
}

type User interface {
	GetUser(ctx context.Context, email string) (*entitiy.User, error)
	CreateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error)
	UpdateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error)
	DeleteUser(ctx context.Context, email string) error
	FindUser(ctx context.Context, email string) (*entitiy.User, error)
}
