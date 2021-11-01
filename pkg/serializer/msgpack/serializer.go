package msgpack

import (
	"golang-interview-project-masaru-ohashi/pkg/team"

	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type Member struct{}

func (m *Member) Decode(pogo []byte) (requestMember *team.RequestMember, err error) {
	if err = msgpack.Unmarshal(pogo, requestMember.Member); err != nil {
		return nil, errors.Wrap(err, "serializer.Member.Decode")
	}
	return requestMember, nil
}

func (m *Member) Encode(pogo interface{}) ([]byte, error) {
	serialized, err := msgpack.Marshal(pogo)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Member.Encode")
	}
	return serialized, nil
}
