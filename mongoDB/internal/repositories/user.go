package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mongoBD/internal/user"
)

type UsersRepositories struct {
	client *mongo.Client
}

func NewUsersRepositories(client *mongo.Client) UsersRepositories {
	return UsersRepositories{client: client}
}

func (r UsersRepositories) GetCollections() *mongo.Collection {
	return r.client.Database("default").Collection("users")
}

func (r UsersRepositories) Create(ctx context.Context, userID int, name string) error {
	_, err := r.GetCollections().InsertOne(ctx, &user.User{
		ID:   userID,
		Name: name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r UsersRepositories) Get(ctx context.Context, userID int) (user.User, error) {

	filter := bson.D{{"id", userID}}

	var result user.User

	if err := r.GetCollections().FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (r UsersRepositories) Delete(ctx context.Context, userID int) error {

	filter := bson.D{{"id", userID}}

	_, err := r.GetCollections().DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
