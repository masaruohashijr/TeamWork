package member

type MemberService interface {
	Find(name string) (*Member, error)
	Store(member *Member) error
}
