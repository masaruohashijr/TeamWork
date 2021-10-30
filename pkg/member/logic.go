package member

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrMemberNotFound = errors.New("Member Not Found")
	ErrMemberInvalid  = errors.New("Member Invalid")
)

type memberService struct {
	memberRepo MemberRepository
}

func NewMemberService(memberRepo MemberRepository) MemberService {
	return &memberService{
		memberRepo,
	}
}

func (r *memberService) Find(code string) (*Member, error) {
	return r.memberRepo.Find(code)
}

func (r *memberService) Store(member *Member) error {
	if err := validate.Validate(member); err != nil {
		return errs.Wrap(ErrMemberInvalid, "service.Member.Store")
	}
	member.CreatedAt = time.Now().UTC().Unix()
	return r.memberRepo.Store(member)
}
