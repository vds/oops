package mongodb_test

import "labix.org/v2/mgo"

type MongoDBStorage struct{}

func foo() {
	_ = MongoDBStorage{}
	session, err := mgo.Dial("localhost:27017")
}
