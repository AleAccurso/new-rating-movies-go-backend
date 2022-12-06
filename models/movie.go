package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
	MovieDbId   int32              `bson:"movie_db_id,omitempty" json:"movie_db_id"`
	ReleaseDate string             `bson:"release_date,omitempty" json:"release_date"`
	Director    string             `bson:"director,omitempty" json:"director"`
	Casting     string             `bson:"casting,omitempty" json:"casting"`
	VoteAverage int8               `bson:"vote_average,omitempty" json:"vote_average"`
	VoteCount   int32              `bson:"vote_count,omitempty" json:"vote_count"`
	Genre       []string           `bson:"genre,omitempty" json:"genre"`
	En          MovieLocalInfo     `bson:"en,omitempty" json:"en"`
	Fr          MovieLocalInfo     `bson:"fr,omitempty" json:"fr"`
	It          MovieLocalInfo     `bson:"it,omitempty" json:"it"`
	Nl          MovieLocalInfo     `bson:"nl,omitempty" json:"nl"`
}
