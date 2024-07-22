package blog

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, blog *Blog) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(dbClient *mongo.Client) *MongoRepository {
	return &MongoRepository {
		collection: dbClient.Database("blogs").Collection("blogs"),
	}
}

func (r *MongoRepository) Create(ctx context.Context, blog *Blog) error {
	_, err := r.collection.InsertOne(ctx, blog)
	return err
}