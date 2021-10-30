package member

type MemberRepository interface {
	Find(name string) (*Member, error)
	Store(member *Member) error
}
