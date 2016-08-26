package thiago

type Subscriber struct {
	Name string `bson:"data"`
	Tags []string `bson:"tags"`
}

func (s *Session) Subscribe(name string, tags []string) error{
	session := s.Session
	subscriber := Subscriber{
		Name: name,
		Tags: tags,
	}
	coll := session.DB(s.Database).C(s.Subscribers)
	if err := coll.Insert(subscriber); err != nil {
		return err
	}
	return nil
}
