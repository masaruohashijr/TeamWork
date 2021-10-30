package json

import (
	"encoding/json"
	"golang-interview-project-masaru-ohashi/pkg/member"

	"github.com/pkg/errors"
)

type Member struct{}

func (r *Member) Decode(input []byte) (*member.Member, error) {
	member := &member.Member{}
	if err := json.Unmarshal(input, member); err != nil {
		return nil, errors.Wrap(err, "serializer.Member.Decode")
	}
	return member, nil
}

func (r *Member) Encode(input *member.Member) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Member.Encode")
	}
	return rawMsg, nil
}
