package repository

import (
	"context"

	"device-control-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceRepository interface {
	AddDevice(device models.Device) (string, error)
	GetDevices() ([]models.Device, error)
	GetDevice(id string) (models.Device, error)
	UpdateDevice(device models.Device) error
	DeleteDevice(id string) error
}

type deviceRepository struct {
	collection *mongo.Collection
}

func NewDeviceRepository(db *mongo.Database) DeviceRepository {
	return &deviceRepository{
		collection: db.Collection("devices"),
	}
}

func (r *deviceRepository) AddDevice(device models.Device) (string, error) {
	res, err := r.collection.InsertOne(context.Background(), device)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(string)
	return id, nil
}

func (r *deviceRepository) GetDevices() ([]models.Device, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var devices []models.Device
	if err = cursor.All(context.Background(), &devices); err != nil {
		return nil, err
	}
	return devices, nil
}

func (r *deviceRepository) GetDevice(id string) (models.Device, error) {
	var device models.Device
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&device)
	if err != nil {
		return models.Device{}, err
	}
	return device, nil
}

func (r *deviceRepository) UpdateDevice(device models.Device) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": device.ID}, bson.M{"$set": device})
	return err
}

func (r *deviceRepository) DeleteDevice(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
