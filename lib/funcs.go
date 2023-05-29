package lib

import "go.mongodb.org/mongo-driver/bson/primitive"

func StringToObjectID(str string) primitive.ObjectID {
	object_id, err := primitive.ObjectIDFromHex(str)
	Logger(err)
	return object_id
}
