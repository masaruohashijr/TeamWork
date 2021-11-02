package ports

import "golang-interview-project-masaru-ohashi/pkg/team"

type MemberPort interface {
	GetAll() ([]team.Member, error)
	Get(member_name string) (team.Member, error)
	Post(member team.Member) (team.Member, error)
	Put(member team.Member) (team.Member, error)
	Delete(member team.Member) (team.Member, error)
}
