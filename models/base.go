package models

import (
	"context"
	"fmt"
	"github.com/iambenkay/nairacheck/services"
	"github.com/iambenkay/nairacheck/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
)

type Direction string

const (
	ASC  Direction = "ASC"
	DESC Direction = "DESC"
)

type PageSortParams struct {
	Page      int64 `query:"page"`
	Limit     int64 `query:"limit"`
	Paged     bool
	Sort      string    `query:"sort"`
	Direction Direction `query:"direction"`
}

func findOne(filter interface{}, destination interface{}, name string) (err error) {
	c := coll(name)

	var doc *bson.D
	doc, err = structOrMapToBson(filter)
	if err != nil {
		log.Println(err)
		return
	}
	var result *mongo.SingleResult
	utils.Contextualize(func(ctx context.Context) {
		result = c.FindOne(ctx, doc)
		err := result.Decode(destination)
		if err != nil {
			log.Println(err)
			return
		}
	})
	return
}

func find(filter interface{}, params *PageSortParams, destination interface{}, name string) (err error) {
	c := coll(name)

	filter, err = structOrMapToBson(filter)
	if err != nil {
		log.Println(err)
		return
	}

	var opt = options.Find()
	if params != nil {
		if params.Page <= 0 {
			params.Page = 1
		}
		if params.Limit <= 0 {
			params.Limit = 20
		}
		if params.Sort == "" {
			params.Sort = "date_created"
		}
		var dir int
		if params.Direction == DESC {
			dir = -1
		} else {
			dir = -1
		}
		if params.Paged {
			opt.SetSkip((params.Page - 1) * params.Limit).SetLimit(params.Limit)
		}
		opt.SetSort(bson.D{{params.Sort, dir}})
	}

	utils.Contextualize(func(ctx context.Context) {
		var cursor *mongo.Cursor
		cursor, err = c.Find(ctx, filter, opt)

		if err != nil {
			log.Println(err)
			return
		}
		defer func() {
			_ = cursor.Close(ctx)
		}()
		err = cursor.All(ctx, destination)
		if err != nil {
			log.Println(err)
			return
		}
		err = cursor.Err()
		if err != nil {
			log.Println(err)
			return
		}
	})
	return
}

func insertOne(document interface{}, name string) (id interface{}, err error) {
	c := coll(name)

	var doc *bson.D

	doc, err = structOrMapToBson(document)
	if err != nil {
		log.Println(err)
		return
	}
	var result *mongo.InsertOneResult
	utils.Contextualize(func(ctx context.Context) {
		result, err = c.InsertOne(ctx, doc)

		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Inserted %v into %s collection\n", result.InsertedID, name)
	})
	return result.InsertedID, nil
}

func deleteOne(document interface{}, name string) (err error) {
	c := coll(name)

	var doc *bson.D

	doc, err = structOrMapToBson(document)
	if err != nil {
		log.Println(err)
		return
	}
	var result *mongo.DeleteResult
	utils.Contextualize(func(ctx context.Context) {
		result, err = c.DeleteOne(ctx, doc)

		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Deleted %v document(s) from %s collection\n", result.DeletedCount, name)
	})
	return nil
}

func updateOne(document interface{}, update interface{}, name string) (err error) {
	c := coll(name)

	var doc *bson.D

	doc, err = structOrMapToBson(document)
	if err != nil {
		log.Println(err)
		return
	}

	update, err = structOrMapToBson(update)
	if err != nil {
		log.Println(err)
		return
	}

	var result *mongo.UpdateResult
	utils.Contextualize(func(ctx context.Context) {
		result, err = c.UpdateOne(ctx, doc, bson.D{{"$set", update}})

		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Modified %v document(s) from %s collection\n", result.ModifiedCount, name)
	})
	return nil
}

func updateMany(document interface{}, update interface{}, name string) (err error) {
	c := coll(name)

	var doc *bson.D

	doc, err = structOrMapToBson(document)
	if err != nil {
		log.Println(err)
		return
	}

	update, err = structOrMapToBson(update)
	if err != nil {
		log.Println(err)
		return
	}

	var result *mongo.UpdateResult
	utils.Contextualize(func(ctx context.Context) {
		result, err = c.UpdateMany(ctx, doc, bson.D{{"$set", update}})

		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Modified %v document(s) from %s collection\n", result.ModifiedCount, name)
	})
	return nil
}

func coll(name string) *mongo.Collection {
	return services.Bean.DatabaseClient.Database("nairacheck").Collection(name)
}

func structOrMapToBson(in interface{}) (out *bson.D, err error) {
	out = &bson.D{}
	if in == nil {
		return out, nil
	}
	if reflect.TypeOf(in).Kind() == reflect.Struct ||
		reflect.TypeOf(in).Kind() == reflect.Map {
		var data []byte
		data, err = bson.Marshal(in)
		if err != nil {
			return nil, err
		}
		err = bson.Unmarshal(data, out)
		if err != nil {
			return nil, err
		}
		return
	}
	if reflect.TypeOf(in) == reflect.TypeOf(out) {
		outRaw := in.(bson.D)
		return &outRaw, nil
	}
	return
}
