package user

type User struct {
	ID   int    `bson:"id"`
	Name string `bson:"name"`
}
