package msgpack

import (
	"golang-interview-project-masaru-ohashi/pkg/member"

	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type Member struct{}

func (r *Member) Decode(input []byte) (*member.Member, error) {
	redirect := &member.Member{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Member.Decode")
	}
	return redirect, nil
}

func (r *Member) Encode(input *member.Member) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Member.Encode")
	}
	return rawMsg, nil
}
