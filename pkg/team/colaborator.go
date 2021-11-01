package team

import (
	"encoding/xml"
)

type Colaborator struct {
	XMLName   xml.Name `xml:"member"`
	Name      string   `json:"name" xml:"name,attr" bson:"_id" msgpack:"name" validate:"required"`
	Agreement string   `json:"agreement" xml:"agreement,attr" bson:"agreement" msgpack:"agreement" validate:"required"`
	CreatedAt int64    `json:"created_at" xml:"created_at,created_at" bson:"created_at" msgpack:"created_at"`
	Tags      []string `json:"tags" xml:"tags,tags" bson:"tags" msgpack:"tags"`
}
