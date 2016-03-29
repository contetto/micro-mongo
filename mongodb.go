package mongodb

import (
	"errors"
	"fmt"
	"log"

	"github.com/micro/go-micro"

	"gopkg.in/mgo.v2"
)

/*
 TODOs:
  - Use global logger
  - Any safe way to close mongo session?
*/

type MongoSession struct {
	session *mgo.Session
	dbName  string
}

func New(service micro.Service) (*MongoSession, error) {
	log.Println("Initializing mongo connection...")

	mongoURL, found := service.Server().Options().Metadata["MONGO_URL"]
	if !found {
		fmt.Println("WARNING: MONGO_URL not set in Server's Metadata. Connect to localhost now.")
		mongoURL = "localhost:27017"
	}

	dbName, found := service.Server().Options().Metadata["MONGO_DB"]
	if !found {
		return &MongoSession{}, errors.New("MONGO_DB not set in Server's Metadata")
	}

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		fmt.Println("NO DB Connection to:", mongoURL)
		return &MongoSession{}, err
	}

	return &MongoSession{session, dbName}, nil
}

func (m *MongoSession) GetCollection(name string) *mgo.Collection {
	return m.session.DB(m.dbName).C(name)
}
