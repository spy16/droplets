package mongo

import (
	"gopkg.in/mgo.v2"
)

// Connect to a MongoDB instance located by mongo-uri using the `mgo`
// driver.
func Connect(uri string, failFast bool) (*mgo.Database, func(), error) {
	di, err := mgo.ParseURL(uri)
	if err != nil {
		return nil, doNothing, err
	}

	di.FailFast = failFast
	session, err := mgo.DialWithInfo(di)
	if err != nil {
		return nil, doNothing, err
	}

	return session.DB(di.Database), session.Close, nil
}

func doNothing() {}
