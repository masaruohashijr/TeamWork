package team

import "fmt"

type Employee struct {
	Colaborator Colaborator
	Role        string `json:"role" xml:"role,attr" bson:"role" msgpack:"role"`
}

func (e Employee) String() string {
	return fmt.Sprintf("Team Member: %s ", e.GetName())
}

func (e *Employee) GetName() string {
	return e.Colaborator.Name
}
func (e *Employee) GetAgreement() string {
	return e.Colaborator.Agreement
}
func (e *Employee) GetCreatedAt() int64 {
	return e.Colaborator.CreatedAt
}
func (e *Employee) CreatedAt(created_at int64) {
	e.Colaborator.CreatedAt = created_at
}
func (e *Employee) GetTags() []string {
	return e.Colaborator.Tags
}
