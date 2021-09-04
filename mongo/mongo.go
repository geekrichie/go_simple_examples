package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoConn struct{

}

func Ping(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
}

func InsertInto(client *mongo.Client) {
	var collection = client.Database("testing").Collection("number")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{
		{"name","test"},
		{ "age" , 68},
		{"method", "post"},
		{"level","hard"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.InsertedID)
}

func Query(client *mongo.Client) ([]bson.D, error){
	var collection = client.Database("testing").Collection("number")
	ctx , cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{{"age" ,1}})
	res, err := collection.Find(ctx, bson.D{{"name", "test"}}, opts)
	if err != nil {
		return nil, err
	}
	var data = make([]bson.D, 0)
	for res.Next(ctx) {
		var data1 bson.D
		if err := res.Decode(&data1); err != nil {
			return nil, err
		}
		data = append(data, data1)
	}
	return data, nil
}

func Update(client *mongo.Client) {
	var collection = client.Database("testing").Collection("number")
	ctx , cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	datas, err := Query(client)
	if err != nil {
		log.Fatal(err)
	}
	 name :=  datas[0].Map()["name"]
	fmt.Println(name)
	if err != nil {
		log.Fatal(err)
	}
	result, err := collection.UpdateMany(ctx, bson.D{{"name", name}},bson.D{
		{"$set",bson.D{
			{"name","test"},
			{ "age" , 79},
			{"method", "post"},
			{"level","soft"},
		}},
	} )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
func Delete(client *mongo.Client) {
	var collection = client.Database("testing").Collection("number")
	ctx , cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	datas, err := Query(client)
	if err != nil {
		log.Fatal(err)
	}
	name :=  datas[0].Map()["name"]
	fmt.Println(name)
	if err != nil {
		log.Fatal(err)
	}
	result, err := collection.DeleteMany(ctx, bson.D{{"name", name}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func main() {
	ctx,cancel := context.WithTimeout(context.Background(), 10 *time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.33.13:27017"))
	if err != nil {
		log.Fatal(err)
	}
     defer func() {
     	if err := client.Disconnect(ctx); err != nil {
     		panic(err)
		}
	 }()
	 Ping(client)
	//[{_id ObjectID("611f6b5fcef57894b7abdbfa")} {name test} {age 56} {method get}]
	//[{_id ObjectID("611f6ced7e4580538ae75bce")} {name test} {age 68} {method post} {level hard}]

	//InsertInto(client)
	datas, err := Query(client)
	if err != nil {
		log.Fatal(err)
	}
	for _, data := range datas {
		fmt.Println(data)
	}
	//Update(client)
	//Delete(client)
}