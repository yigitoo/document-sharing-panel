package lib

import "go.mongodb.org/mongo-driver/bson/primitive"

type IsAdmin bool
type Username string
type TextContent string
type DocumentID []primitive.ObjectID

type UserManagement interface {
	CreateUser()
	ValidateLogin()
	ChangeUserType(user_id string)
}

type User struct {
	UserID   primitive.ObjectID `bson:"_id"`
	Type     IsAdmin            `bson:"type"`
	Username Username           `bson:"username"`
	Password string             `bson:"password"`
}

type DocumentManager interface {
	ChangeFileContent(text_content string) error
	ClearAccessList()
	RemoveSomeone(user_id string)
	AddSomeone(user_id string)
	AddGroup(user_list []string) error
	RemoveGroup(user_list []string) error
}

type Document struct {
	ID         primitive.ObjectID   `bson:"_id"`
	Content    TextContent          `bson:"content"`
	Author     Username             `bson:"author"`
	AccessList []primitive.ObjectID `bson:"access_list"`
}
