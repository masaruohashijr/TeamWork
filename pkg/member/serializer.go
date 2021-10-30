package member

type MemberSerializer interface {
	Decode(input []byte) (*Member, error)
	Encode(input *Member) ([]byte, error)
}
