package xml

import (
	"encoding/xml"
	"golang-interview-project-masaru-ohashi/pkg/team"

	"github.com/pkg/errors"
)

type Member struct{}

func (m *Member) Decode(pogo []byte) (requestMember *team.RequestMember, err error) {
	if err = xml.Unmarshal(pogo, requestMember); err != nil {
		return nil, errors.Wrap(err, "serializer.Member.Decode")
	}
	return requestMember, nil
}

func (m *Member) Encode(pogo interface{}) ([]byte, error) {
	serialized, err := xml.MarshalIndent(pogo, "  ", "    ")
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Member.Encode")
	}
	return serialized, nil
}
