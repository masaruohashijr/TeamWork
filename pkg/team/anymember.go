package team

import "fmt"

type AnyMember struct {
	Member Member
}

func (a AnyMember) String() string {
	return fmt.Sprintf("Team Member: %s ", a.Member.GetName())
}

func (a *AnyMember) GetName() string {
	return a.Member.GetName()
}
func (a *AnyMember) GetAgreement() string {
	return a.Member.GetAgreement()
}
func (a *AnyMember) GetCreatedAt() int64 {
	return a.Member.GetCreatedAt()
}
func (a *AnyMember) CreatedAt(created_at int64) {
	a.Member.CreatedAt(created_at)
}
func (a *AnyMember) GetTags() []string {
	return a.Member.GetTags()
}
