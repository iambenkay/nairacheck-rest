package models

import (
	"context"
	"github.com/iambenkay/nairacheck/services"
	"github.com/iambenkay/nairacheck/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	Sort      string `query:"sort"`
	Direction Direction `query:"direction"`
}

func find(filter interface{}, params *PageSortParams, destination interface{}, name string) (err error) {
	c := coll(name)

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
			opt.SetSkip(params.Page * params.Limit).SetLimit(params.Limit)
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

func coll(name string) *mongo.Collection {
	return services.Bean.DatabaseClient.Database("nairacheck").Collection(name)
}
