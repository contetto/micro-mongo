package mongodb

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

// LH: @todo Maybe better to make all of this into a MongoSession manager where each session will hold a single collection.
type MongoSession struct {
	session *mgo.Session
}

// LH: @todo this should later get a config struct passed to be set up e.g. for production mode etc.
func New() *mgo.Session {
	fmt.Println("Initializing mongo connection...")

	var mongoUrl string

	// LH: @todo add more advanced behaviour for production environment e.g. discover which mongoDB to be used
	eUrl := os.Getenv("MONGO_URL")
	if len(eUrl) == 0 {
		mongoUrl = "localhost:27017"
	} else {
		mongoUrl = eUrl
	}

	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Println("In application.json you have specified to use mongoDB, but unfortunately mongoDB is not accessible at " + mongoUrl + ". Exiting now")
		fmt.Println("NO DB Connection to:", mongoUrl)
		panic("Terminating...")

	}

	return session
}

/*
	USAGE EXAMPLE:

	type Person struct {
		Name string
		Phone string
	}

		c := session.DB("test").C("people")
		err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
			&Person{"Cla", "+55 53 8402 8510"})
		if err != nil {
			log.Fatal(err)
		}

		result := Person{}
		err = c.Find(bson.M{"name": "Ale"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Phone:", result.Phone)
*/
