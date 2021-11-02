package main

import (
	"flag"
	adapters "golang-interview-project-masaru-ohashi/cmd/client/adapters/REST"
	"golang-interview-project-masaru-ohashi/cmd/client/mock"
	"golang-interview-project-masaru-ohashi/cmd/client/ports"
	"golang-interview-project-masaru-ohashi/cmd/common"
	"golang-interview-project-masaru-ohashi/pkg/team"
	"math/rand"
)

var repoTypePtr *string
var apiTypePtr *string

func main() {
	apiTypePtr = flag.String("api", "", "domain:{GraphQL,gRPC,REST,RSocket}")
	repoTypePtr = flag.String("repo", "", "domain:{elastic,maria,memory,mongo,mysql,postgres,redis,sqlite}")
	flag.Parse()
	newMemberAPI := settleApi(apiTypePtr, repoTypePtr)
	newMember1 := mock.NewMember("Masaru", common.CONTRACTOR, []string{"GO", "Scala"}, 10, "")
	newMemberAPI.Post(newMember1)
	newMember2 := mock.NewMember("Ária", common.EMPLOYEE, []string{"Scala", "C#"}, 0, "Software Engineer")
	newMemberAPI.Post(newMember2)
	newMember3 := mock.NewMember("Mariana", common.CONTRACTOR, []string{"Ruby", "Python"}, 10, "")
	newMemberAPI.Post(newMember3)
	newMember4 := mock.NewMember("Cristina", common.CONTRACTOR, []string{"R", "Scala"}, 10, "")
	newMemberAPI.Post(newMember4)
	newMember5 := mock.NewMember("Daniel", common.EMPLOYEE, []string{"GO", "C"}, 0, "Software Developer")
	newMemberAPI.Post(newMember5)
	members, _ := newMemberAPI.GetAll()
	forEachPrintln(members)
	ms := toInterfaces(members)
	choosen := toMember(selectRandom(ms))
	name := choosen.GetName()
	member, _ := newMemberAPI.Get(name)
	agreement := member.GetAgreement()
	if agreement == common.CONTRACTOR {
		member.(*team.Contractor).Duration = 5
	} else {
		member.(*team.Employee).Role = "Scala Developer"
	}
	member, _ = newMemberAPI.Put(member)
}

func settleApi(apiType, repoType *string) ports.MemberPort {
	switch *apiType {
	case "GraphQL":
		repo := adapters.NewMemberRESTApi(*repoType, common.REST_URL)
		return repo
	case "gRPC":
		repo := adapters.NewMemberRESTApi(*repoType, common.REST_URL)
		return repo
	case "REST":
		repo := adapters.NewMemberRESTApi(*repoType, common.REST_URL)
		return repo
	case "RSocket":
		repo := adapters.NewMemberRESTApi(*repoType, common.REST_URL)
		return repo
	case "SOAP":
		repo := adapters.NewMemberRESTApi(*repoType, common.REST_URL)
		return repo
	}
	return nil
}

func forEachPrintln(members []team.Member) {
	for _, v := range members {
		println(v.GetName())
	}
}

func selectRandom(a []interface{}) interface{} {
	pos := rand.Intn(len(a))
	return a[pos]
}

func toInterfaces(x []team.Member) []interface{} {
	y := make([]interface{}, len(x))
	for i, v := range x {
		y[i] = v
	}
	return y
}

func toMember(x interface{}) team.Member {
	var y team.Member
	y = x.(team.Member)
	return y
}
