package repoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	models "go-jwt/model"
	repo "go-jwt/repository"
)

type UserRepoImpl struct {
	Db *mongo.Database
}

func NewUserRepo(db *mongo.Database) repo.UserRepo {
	return &UserRepoImpl{
		Db: db,
	}
}

func (mongo *UserRepoImpl) Insert(user models.User) error {
	bbytes, _ := bson.Marshal(user)
	_, err := mongo.Db.Collection("users").InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}
	return nil
}