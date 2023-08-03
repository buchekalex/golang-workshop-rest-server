package repository

import (
	"context"
	"go-mongo-crud-rest-api/internal/entitiy"
	"go-mongo-crud-rest-api/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func (r repository) CreateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *mongo.Database) usecase.Repository {
	return &repository{db: db}
}

func (r repository) GetUser(ctx context.Context, email string) (*entitiy.User, error) {
	// TODO: implement me
	return nil, nil
}

func (r repository) UpdateUser(ctx context.Context, user *entitiy.User) (*entitiy.User, error) {
	// TODO: implement me
	return user, nil
}

func (r repository) DeleteUser(ctx context.Context, email string) error {
	// TODO: implement me
	return nil
}

type user struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

func fromModel(in *entitiy.User) user {
	return user{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}

func toModel(in user) *entitiy.User {
	return &entitiy.User{
		ID:       in.ID.String(),
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}
