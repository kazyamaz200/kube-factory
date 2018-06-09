package store

import (
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// ArangoCollection as driver.CollectionDocuments
type ArangoCollection = driver.CollectionDocuments

// ConnectArangoCollection is ...
func ConnectArangoCollection(endpoints []string, dbName string, userName string, password string, collectionName string) ArangoCollection {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: endpoints,
	})
	if err != nil {
		panic(err)
	}

	credentials := driver.BasicAuthentication(userName, password)
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

	connector, err := db.Collection(nil, collectionName)
	if err != nil {
		panic(err)
	}

	return connector
}
