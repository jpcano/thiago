package thiago

import (
	"gopkg.in/mgo.v2"
	"time"
)

type Session struct {
	Session *mgo.Session
	Database string
	Subscribers string
	Publications string
}

type SessionInfo struct {
	Host string
	Username string
	Password string
	Database string
	Publications string
	Subscribers string
}

func GetSession(info SessionInfo) (*Session, error) {
	mgosession, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{info.Host},
		Username: info.Username,
		Password: info.Password,
		Database: info.Database,
		Timeout:  60 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	session := &Session{
		Session: mgosession,
		Database: info.Database,
		Subscribers: info.Subscribers,
		Publications: info.Publications,
	}
	
	return session, nil
}

// func GetSession() (*mgo.Session, error) {
// 	if session == nil {
// 		var err error

// 		session, err = mgo.DialWithInfo(&mgo.DialInfo{
// 			Addrs:    []string{Host},
// 			Username: Username,
// 			Password: Password,
// 			Database: Database,
// 			Timeout:  60 * time.Second,
// 		})

// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return session, nil
// }

func (p *Session) Close() {
	p.Session.Close()
}
