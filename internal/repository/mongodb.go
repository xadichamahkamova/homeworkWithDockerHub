package repository

import (
	"context"
	"errors"
	models "service/internal/models"
	m "service/internal/pkg"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepo struct {
	MDB *m.MongoDB
}

func NewTaskRepo (db *m.MongoDB) ITaskStorage {
	return &TaskRepo{
		MDB: db,
	}
}
var ctx = context.Background()

func (db *TaskRepo) CreateTask(req *models.Task) (string, error) {

	result, err := db.MDB.Collection.InsertOne(ctx, bson.D{
		{Key: "title", Value: req.Title},
		{Key: "description", Value: req.Description},
		{Key: "status", Value: req.Status},
		{Key: "created_at", Value: time.Now().Format("2006-01-02 15:04:05")},
	})
	if err != nil {
		return "", err
	}
	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objectId.Hex(), nil
	} else {
		return "", errors.New("unexpected type for object id")
	}
}

func (db *TaskRepo) GetTask(id string) (*models.Task, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	m := models.MongoTask{}
	singleResult := db.MDB.Collection.FindOne(ctx, filter)
	if err := singleResult.Decode(&m); err != nil {
		return nil, err
	}
	resp := models.Task{
		Id:          m.Id.Hex(),
		Title:       m.Title,
		Description: m.Description,
		Status:      m.Status,
		CreatedAt:   m.CreatedAt,
	}
	return &resp, nil
}

func (db *TaskRepo) ListOfTask() ([]*models.Task, error) {

	var resp []*models.Task
	cursor, err := db.MDB.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var m models.MongoTask
		if err := cursor.Decode(&m); err != nil {
			return nil, err
		}
		item := &models.Task{
			Id:          m.Id.Hex(),
			Title:       m.Title,
			Description: m.Description,
			Status:      m.Status,
			CreatedAt:   m.CreatedAt,
		}
		resp = append(resp, item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return resp, nil
}

func (db *TaskRepo) UpdateTask(req *models.Task) (string, error) {

	objectId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"title":       req.Title,
			"description": req.Description,
			"status":      req.Status,
		},
	}
	result, err := db.MDB.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}
	if result.ModifiedCount == 0 {
		return "", errors.New("FAILED")
	}

	status := "Task updated successfully"
	return status, nil
}

func (db *TaskRepo) DeleteTask(id string) (string, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": objectId}
	deleteResult, err := db.MDB.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	if deleteResult.DeletedCount == 0 {
		return "", errors.New("FAILED")
	}

	status := "Task deleted successfully"
	return status, nil
}