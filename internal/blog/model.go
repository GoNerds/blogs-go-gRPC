package blog

type Blog struct {
	ID      string `bson:"_id, omitempty"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
	Author  string `bson:"author"`
}