package utils

import (
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// NewArangoCollection is ...
func NewArangoCollection(endpoints []string, dbName string, collectionName string) driver.Collection {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: endpoints,
	})
	if err != nil {
		panic(err)
	}

	credentials := driver.BasicAuthentication("testuser", "testuser")
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: credentials,
	})
	if err != nil {
		panic(err)
	}

	db, err := client.Database(nil, dbName)
	if err != nil {
		panic(err)
	}

	collection, err := db.Collection(nil, collectionName)
	if err != nil {
		panic(err)
	}

	return collection
}
