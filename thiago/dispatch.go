package thiago

import (
	"gopkg.in/mgo.v2/bson"
	// "fmt"
)

func (s *Session) FindSubscriberByTags(tags []string) ([]string, error){
	session := s.Session
	coll := session.DB(s.Database).C(s.Subscribers)
	var subscribers []Subscriber
	if err := coll.Find(bson.M{ "tags": bson.M{"$in": tags}}).Select(bson.M{"_id":0}).All(&subscribers); err != nil {
		return nil, err
	}
	var results []string
	for _, sub := range subscribers{
		// fmt.Println(sub.Name)
		results = append(results, sub.Name)
	}
		
	return results, nil
}
