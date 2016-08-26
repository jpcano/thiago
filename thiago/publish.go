package thiago

type Publication struct {
	Data string `bson:"data"`
	Tags []string `bson:"tags"`
}

func (s *Session) Publish(data string, tags []string) error{
	session := s.Session
	publication := Publication{
		Data: data,
		Tags: tags,
	}
	coll := session.DB(s.Database).C(s.Publications)
	if err := coll.Insert(publication); err != nil {
		return err
	}
	return nil
}




