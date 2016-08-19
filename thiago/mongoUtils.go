package thiago

import (
	"gopkg.in/mgo.v2"
	"time"
)

var session *mgo.Session

const (
	Host       = "127.0.0.1:27017"
	Username   = ""
	Password   = ""
	Database   = "thiago"
	Publications = "publications"
	Subscribers = "subscribers"
)

func GetSession() (*mgo.Session, error) {
	if session == nil {
		var err error

		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{Host},
			Username: Username,
			Password: Password,
			Database: Database,
			Timeout:  60 * time.Second,
		})

		if err != nil {
			return nil, err
		}
	}

	return session, nil
}

func Close() {
	session.Close()
}
