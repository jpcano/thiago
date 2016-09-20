package thiago

// import (
// 	"gopkg.in/mgo.v2/bson"
// )

type Publication struct {
	Data string `bson:"data"`
	Tags interface{} `bson:"tags"`
}

func (p *Publication) New(data string, tags interface{}) {
	p.Data = data
	p.Tags = tags
}

func (s *Session) Publish(data string, tags interface{}) error{
	var publication Publication
	publication.New(data, tags)

	coll := s.Session.DB(s.Database).C(s.Publications)
	if err := coll.Insert(publication); err != nil {
		return err
	}
	return nil
}
