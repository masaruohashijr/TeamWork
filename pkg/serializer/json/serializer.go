package json

import (
	"encoding/json"
	"golang-interview-project-masaru-ohashi/pkg/team"
	"strings"

	"github.com/pkg/errors"
)

type Member struct{}

func (m *Member) Decode(pogo []byte) (*team.RequestMember, error) {
	str := string(pogo)
	var rm *team.RequestMember
	if strings.Contains(str, "\"agreement\":\"EMPLOYEE\"") {
		rm = &team.RequestMember{
			RepoType: "",
			Member:   &team.Employee{},
		}
	} else {
		rm = &team.RequestMember{
			RepoType: "",
			Member:   &team.Contractor{},
		}
	}
	err := json.Unmarshal(pogo, rm)
	if err == nil {
		return rm, nil
	}
	if _, ok := err.(*json.UnmarshalTypeError); err != nil && !ok {
		return nil, errors.Wrap(err, "serializer.Member.Decode")
	}
	return nil, errors.Wrap(err, "serializer.Member.Decode")
}

func (m *Member) Encode(pogo interface{}) ([]byte, error) {
	serialized, err := json.Marshal(pogo)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Member.Encode")
	}
	return serialized, nil
}
