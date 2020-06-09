package db

import (
	"context"
	"github.com/tteeoo/mudcord/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var ctx context.Context
var Cancel context.CancelFunc
var users *mongo.Collection
var servers *mongo.Collection

func init() {
	ctx, Cancel = context.WithCancel(context.Background())

	mongoURI := os.Getenv("MUDCORD_MONGO_URI")
	if mongoURI == "" {
		util.Logger.Fatal("MUDCORD_MONGO_URI not in environment")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	util.CheckFatal(err)

	util.Logger.Println("Connected to db")

	users = client.Database("mudcord").Collection("users")
	servers = client.Database("mudcord").Collection("servers")
}

type ItemQuan struct {
	ID   string
	Quan int
}

// NewServer creates a new server in our db
func NewServer(id string) (*Server, error) {
	server := Server{
		ID:     id,
		Prefix: ".",
	}
	_, err := servers.InsertOne(ctx, server)

	if err != nil {
		return &Server{}, err
	}

	return &server, nil
}

// GetServer gets a server from db based on discord id
func GetServer(id string) (*Server, error) {
	data := servers.FindOne(ctx, bson.M{"id": id})

	var server Server
	err := data.Decode(&server)

	if err != nil {
		return &Server{}, err
	}

	return &server, nil
}

// SetServer changes server data in the db based on discord id
func SetServer(newServer *Server) error {
	data := servers.FindOneAndUpdate(ctx, bson.M{"id": newServer.ID}, bson.M{"$set": newServer})

	var server Server
	err := data.Decode(&server)

	if err != nil {
		return err
	}

	return nil
}

// CheckServer checks if a server with the given id exists
func CheckServer(id string) bool {
	_, err := GetServer(id)

	if err == mongo.ErrNoDocuments {
		return false
	}

	return true
}

// DeleteUser deletes a user in our db
func DeleteUser(id string) error {

	_, err := users.DeleteOne(ctx, bson.M{"id": id})

	if err != nil {
		return err
	}

	return nil
}

// NewUser creates a new user in our db
func NewUser(id string) (*User, error) {
	user := User{
		ID:      id,
		Level:   1,
		XP:      0,
		HP:      [2]int{20, 20},
		MP:      [2]int{20, 20},
		Gold:    0,
		Room:    "RoomGreatMarya",
		Hat:     "None",
		Inv:     []*ItemQuan{{ID: "ConsumableCanteen", Quan: 1}},
		Arsenal: []string{"WeaponWoodSword"},
	}
	_, err := users.InsertOne(ctx, user)

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// GetUser gets a user from db based on discord id
func GetUser(id string) (*User, error) {
	data := users.FindOne(ctx, bson.M{"id": id})

	var user User
	err := data.Decode(&user)

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// SetUser changes user data in the db based on discord id
func SetUser(newUser *User) error {
	data := users.FindOneAndUpdate(ctx, bson.M{"id": newUser.ID}, bson.M{"$set": newUser})

	var user User
	err := data.Decode(&user)

	if err != nil {
		return err
	}

	return nil
}

// CheckStarted checks if a user with the given user id exists
func CheckStarted(id string) bool {
	_, err := GetUser(id)

	if err == mongo.ErrNoDocuments {
		return false
	}

	return true
}
