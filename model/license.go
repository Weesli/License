package model

type License struct {
	Owner  string `bson:"owner"`
	Name   string `bson:"name"`
	Key    string `bson:"key"`
	Status string `bson:"status"`
}
