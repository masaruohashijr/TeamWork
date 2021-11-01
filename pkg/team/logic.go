package team

import (
	"errors"
	"fmt"
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

func (r *memberService) GetAll() ([]interface{}, error) {
	return r.memberRepo.DbGetAll()
}

func (r *memberService) Get(name string) (interface{}, error) {
	return r.memberRepo.DbGet(name)
}

// validation rules
func (r *memberService) Create(member interface{}) error {
	switch member.(type) {
	case *Contractor:
		member = member.(*Contractor)
		if member.(*Contractor).Duration == 0 {
			return errs.Wrap(ErrMemberInvalid, fmt.Sprintf("Cannot create member: %s. Missing Duration.",
				member.(*Contractor).GetName()))
		}
		member.(*Contractor).CreatedAt(time.Now().UTC().Unix())
	case *Employee:
		member = member.(*Employee)
		if member.(*Employee).Role == "" {
			return errs.Wrap(ErrMemberInvalid, fmt.Sprintf("Cannot create member: %s. Missing Role.",
				member.(*Contractor).GetName()))
		}
		member.(*Employee).CreatedAt(time.Now().UTC().Unix())
	}
	return r.memberRepo.DbCreate(member)
}

func (r *memberService) Update(member interface{}) error {
	if err := validate.Validate(member); err != nil {
		return errs.Wrap(ErrMemberInvalid, "service.Member.Update")
	}
	//member.CreatedAt = time.Now().UTC().Unix()
	return r.memberRepo.DbCreate(member)
}
func (r *memberService) Delete(member interface{}) error {
	if err := validate.Validate(member); err != nil {
		return errs.Wrap(ErrMemberInvalid, "service.Member.Delete")
	}
	//member.CreatedAt = time.Now().UTC().Unix()
	return r.memberRepo.DbCreate(member)
}
