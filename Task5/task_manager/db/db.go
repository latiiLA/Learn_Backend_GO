// router/types.go
package db

import "go.mongodb.org/mongo-driver/mongo"

type DBCollections struct {
	Tasks     *mongo.Collection
}