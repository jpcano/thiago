package thiago

import (
	"gopkg.in/mgo.v2/bson"
	// "fmt"
)

func (s *Session) FindSubscriberByTags(tags interface{}) ([]string, error){
	coll := s.Session.DB(s.Database).C(s.Subscribers)
	var subscribers []Subscriber
	if err := coll.Find(bson.M{"tags": bson.M{"$in": tags}}).Select(bson.M{"_id":0}).All(&subscribers); err != nil {
		return nil, err
	}
	var results []string
	for _, sub := range subscribers {
		results = append(results, sub.Name)
	}
	return results, nil
}

func (s *Session) FindPublicationByTagsIn(tags interface{}) ([]Publication, error){
	coll := s.Session.DB(s.Database).C(s.Publications)
	var publications []Publication
	if err := coll.Find(bson.M{"tags": bson.M{"$in": tags}}).Select(bson.M{"_id":0}).All(&publications); err != nil {
		return nil, err
	}
	return publications, nil
}

func (s *Session) FindPublicationByTagsAll(tags interface{}) ([]Publication, error){
	coll := s.Session.DB(s.Database).C(s.Publications)
	var publications []Publication
	if err := coll.Find(bson.M{"tags": bson.M{"$all": tags}}).Select(bson.M{"_id":0}).All(&publications); err != nil {
		return nil, err
	}
	return publications, nil
}
