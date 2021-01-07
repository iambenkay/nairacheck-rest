package main

import "os"

func env(name string) string {
	return os.Getenv(name)
}

var (
	MongoURI = env("MONGO_URI")
)
