package mock

import (
	"golang-interview-project-masaru-ohashi/cmd/common"
	"golang-interview-project-masaru-ohashi/pkg/team"
)

func NewMember(name string, agreement common.Agreement, tags []string, duration int, role string) team.Member {
	if agreement == common.CONTRACTOR {
		contractor := &team.Contractor{
			Colaborator: team.Colaborator{
				Name:      name,
				Agreement: string(common.CONTRACTOR),
				Tags:      tags,
			},
			Duration: duration,
		}
		return contractor
	} else {
		employee := &team.Employee{
			Colaborator: team.Colaborator{
				Name:      name,
				Agreement: string(common.EMPLOYEE),
				Tags:      tags,
			},
			Role: role,
		}
		return employee
	}
}
