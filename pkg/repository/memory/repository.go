package repository

import (
	"golang-interview-project-masaru-ohashi/pkg/team"
	"sync"
)

var (
	onceMembersRepo     sync.Once
	instanceMembersRepo *repository
)

type repository struct {
	dataMemory map[int]team.Member
	lastId     int
}

func NewMembersRepository() *repository {
	onceMembersRepo.Do(func() {
		instanceMembersRepo = &repository{
			dataMemory: map[int]team.Member{},
			lastId:     0,
		}
	})
	return instanceMembersRepo
}

func (r *repository) Get(id int) (team.Member, error) {
	var member team.Member
	member = r.dataMemory[id]
	return member, nil
}

func (r *repository) Create(member team.Member) error {
	r.dataMemory[r.lastId] = member
	r.lastId++
	return nil
}

func (r *repository) Update(member team.Member) error {
	r.dataMemory[r.lastId] = member
	r.lastId++
	return nil
}

func (r *repository) Delete(member team.Member) error {
	r.dataMemory[r.lastId] = member
	r.lastId++
	return nil
}
