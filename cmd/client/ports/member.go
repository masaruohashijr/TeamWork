package ports

import "golang-interview-project-masaru-ohashi/pkg/team"

type MemberPort interface {
	GetAll() ([]interface{}, error)
	Get(member_name string) (interface{}, error)
	Post(member team.Member) (team.Member, error)
	Put(interface{}) error
	Delete(interface{}) error
}
