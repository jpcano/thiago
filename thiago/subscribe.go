package thiago

import (
	"gopkg.in/mgo.v2/bson"
	// "fmt"
)

type Subscriber struct {
	Name string `bson:"data"`
	Tags interface{} `bson:"tags"`
}

// type Tagger interface {
// 	Tag() string
// }

func (s *Subscriber) New(name string, tags interface{}) {
	s.Name = name
	s.Tags = tags
}

// func (s *Session) Subscribe(name string, tags interface{}) error{
// 	var subscriber Subscriber
// 	subscriber.New(name, tags)
// 	coll := s.Session.DB(s.Database).C(s.Subscribers)
	
// 	if err := coll.Insert(subscriber); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *Session) Subscribe(name string, tags interface{}) error{
	coll := s.Session.DB(s.Database).C(s.Subscribers)
	
	if _, err := coll.Upsert(bson.M{"data": name}, bson.M{"$addToSet": bson.M{"tags": bson.M{"$each": tags}}}); err != nil {
		return err
	}
	return nil
}
