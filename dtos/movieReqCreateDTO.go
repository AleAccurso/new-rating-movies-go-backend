package dtos

type MovieReqCreateDTO struct {
	MovieDbId int32 `bson:"movie_db_id,omitempty" json:"movie_db_id"`
}
