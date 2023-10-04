package blog

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description Description        `json:"description" bson:"description"`
	Category    string             `json:"category" bson:"category"`
	Created     Created            `json:"created" bson:"created"`
	Likes       int64              `json:"likes" bson:"likes"`
	Dislikes    int64              `json:"dislikes" bson:"dislikes"`
	Comments    Comments           `json:"comments" bson:"comments"`
	Views       int64              `json:"views" bson:"views"`
	Monetized   Monetized          `json:"monetized" bson:"monetized"`
	Updated_at  time.Time          `json:"updated" bson:"updated"`
}

type Description struct {
	Short string `json:"short" bson:"short"`
	Long  string `json:"long" bson:"long"`
}

type Created struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Time time.Time          `json:"time" bson:"time"`
}

type Likes struct {
	Count int64        `json:"count" bson:"count"`
	Users []LikedUsers `json:"likes" bson:"likes"`
}

type LikedUsers struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Time time.Time          `json:"time" bson:"time"`
	Name string             `json:"name" bson:"name"`
}

type Comments struct {
	Count      int64       `json:"count" bson:"count"`
	Commentors []Commentor `json:"commentors" bson:"commentors"`
}

type Commentor struct {
	ID      primitive.ObjectID `json:"id" bson:"id"`
	Name    string             `json:"name" bson:"name"`
	Comment string             `json:"comment" bson:"comment"`
}

type Monetized struct {
	Ok          bool    `json:"ok" bson:"ok"`
	TotalEarned float64 `json:"total_earned" bson:"total_earned"`
}

// type Updated struct {
// 	ID   primitive.ObjectID `json:"id" bson:"id"`
// 	Time time.Time          `json:"time" bson:"time"`
// }
