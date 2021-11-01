package team

import "fmt"

type Contractor struct {
	Colaborator Colaborator
	Duration    int `json:"duration" xml:"duration,attr" bson:"duration" msgpack:"duration"`
}

func (c Contractor) String() string {
	return fmt.Sprintf("Team Member: %s ", c.GetName())
}

func (c *Contractor) GetName() string {
	return c.Colaborator.Name
}
func (c *Contractor) GetAgreement() string {
	return c.Colaborator.Agreement
}
func (c *Contractor) GetCreatedAt() int64 {
	return c.Colaborator.CreatedAt
}
func (c *Contractor) CreatedAt(created_at int64) {
	c.Colaborator.CreatedAt = created_at
}
func (c *Contractor) GetTags() []string {
	return c.Colaborator.Tags
}
