package controllers

import (
	"github.com/krishpranav/golang-management/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
