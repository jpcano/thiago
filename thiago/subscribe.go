package thiago

type Subscriber struct {
	Name string `bson:"data"`
	Tags []string `bson:"tags"`
}

func Subscribe(name string, tags []string) error{
	session, err := GetSession()
	if err != nil {
		return err
	}
	subscriber := Subscriber{
		Name: name,
		Tags: tags,
	}
	coll := session.DB(Database).C(Subscribers)
	if err := coll.Insert(subscriber); err != nil {
		return err
	}
	return nil
}
