package errors

import "golang-interview-project-masaru-ohashi/pkg/team"

func CheckErrorMember(err error, handler func(e error) (team.Member, error)) {
	if err != nil {
		handler(err)
	}
}
