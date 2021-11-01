package main

import (
	"flag"
	adapters "golang-interview-project-masaru-ohashi/cmd/client/adapters/REST"
	"golang-interview-project-masaru-ohashi/cmd/client/mock"
	"golang-interview-project-masaru-ohashi/cmd/client/ports"
	"golang-interview-project-masaru-ohashi/cmd/common"
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
	for _, m := range members {
		println(m)
	}
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
