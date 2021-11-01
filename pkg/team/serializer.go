package team

type MemberSerializer interface {
	Decode(input []byte) (*RequestMember, error)
	Encode(input interface{}) ([]byte, error)
}
