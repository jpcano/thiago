package thiago

type Publication struct {
	Data string `bson:"data"`
	Tags []string `bson:"tags"`
}

func Publish(data string, tags []string) error{
	session, err := GetSession()
	if err != nil {
		return err
	}
	publication := Publication{
		Data: data,
		Tags: tags,
	}
	coll := session.DB(Database).C(Publications)
	if err := coll.Insert(publication); err != nil {
		return err
	}
	return nil
}




