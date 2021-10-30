package service

type NewMemberRequest struct {
	MemberType string `json:"member_type"`
	Duration   int    `json:"duration"`
}

type MemberResponse struct {
	MemberID   int      `json:"member_id"`
	MemberType string   `json:"member_type"`
	Duration   int      `json:"duration"`
	Role       string   `json:"role"`
	Tags       []string `json:"tags"`
}

type MemberService interface {
	NewMember(int, NewMemberRequest) (*MemberResponse, error)
	GetMember(int) ([]MemberResponse, error)
}
