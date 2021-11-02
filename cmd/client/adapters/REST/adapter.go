package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-interview-project-masaru-ohashi/cmd/client/errors"
	"golang-interview-project-masaru-ohashi/cmd/client/ports"
	"golang-interview-project-masaru-ohashi/cmd/common"
	"golang-interview-project-masaru-ohashi/pkg/team"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var errHandler = func(e error) (team.Member, error) {
	return nil, e
}

type memberAPI struct {
	apiUrl   string
	repoType string
}

func NewMemberRESTApi(repoType, apiUrl string) ports.MemberPort {
	return &memberAPI{
		repoType: repoType,
		apiUrl:   apiUrl,
	}
}

func (a *memberAPI) Post(member team.Member) (team.Member, error) {
	apiEndpoint := a.apiUrl + "/Member"
	request := &team.RequestMember{}
	var newMember team.Member
	switch member.(type) {
	case *team.Contractor:
		request = &team.RequestMember{
			RepoType: a.repoType,
			Member:   member.(*team.Contractor),
		}
		newMember = &team.Contractor{}
	case *team.Employee:
		request = &team.RequestMember{
			RepoType: a.repoType,
			Member:   member.(*team.Employee),
		}
		newMember = &team.Employee{}
	}
	requestBody, err := json.Marshal(request)
	var buffer *bytes.Buffer = bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	errors.CheckErrorMember(err, errHandler)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	errors.CheckErrorMember(err, errHandler)
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	b := string(body)
	fmt.Println("response Body:", b)
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf(string(body))
	}
	errors.CheckErrorMember(err, errHandler)
	err = json.Unmarshal(body, &newMember)
	errors.CheckErrorMember(err, errHandler)
	return newMember, nil
}

func (a *memberAPI) Put(member team.Member) (team.Member, error) {
	apiEndpoint := a.apiUrl + "/Member"
	request := &team.RequestMember{}
	var newMember team.Member
	switch member.(type) {
	case *team.Contractor:
		request = &team.RequestMember{
			RepoType: a.repoType,
			Member:   member.(*team.Contractor),
		}
	case *team.Employee:
		request = &team.RequestMember{
			RepoType: a.repoType,
			Member:   member.(*team.Employee),
		}
	}
	requestBody, err := json.Marshal(request)
	var buffer *bytes.Buffer = bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("PUT", apiEndpoint, buffer)
	errors.CheckErrorMember(err, errHandler)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	errors.CheckErrorMember(err, errHandler)
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	b := string(body)
	fmt.Println("response Body:", b)
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf(string(body))
	}
	errors.CheckErrorMember(err, errHandler)
	err = json.Unmarshal(body, &newMember)
	errors.CheckErrorMember(err, errHandler)
	return newMember, nil
}

func (a *memberAPI) Delete(member team.Member) (team.Member, error) {
	//apiEndpoint := a.apiUrl + "/Member"
	/*request := &common.Request{
		RepoType: a.repoType,
		Member:   m,
	}
	requestBody, err := json.Marshal(request)
	var buffer *bytes.Buffer = bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("DELETE", apiEndpoint, buffer)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()*/
	return nil, nil
}

func (a *memberAPI) GetAll() ([]team.Member, error) {
	apiEndpoint := a.apiUrl + "/Members"
	request := &team.RequestMember{
		RepoType: a.repoType,
	}
	requestBody, err := json.Marshal(request)
	var buffer *bytes.Buffer = bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("GET", apiEndpoint, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	b := string(body)
	fmt.Println("response Body:", b)
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf(string(body))
	}
	results := []map[string]interface{}{}
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}
	var members []team.Member
	var agreement string
	for _, resultMap := range results {
		if cmap, ok := resultMap["Colaborator"].(map[string]interface{}); ok {
			agreement = cmap["agreement"].(string)
		}
		if agreement == common.CONTRACTOR {
			members = append(members, buildContractor(resultMap))
		} else if agreement == string(common.EMPLOYEE) {
			members = append(members, buildEmployee(resultMap))
		}
	}
	errors.CheckErrorMember(err, errHandler)
	return members, nil
}

func (a *memberAPI) Get(name string) (member team.Member, err error) {
	apiEndpoint := a.apiUrl + "/Member/" + url.QueryEscape(name)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return member, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return member, err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return member, err
	}
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(body, &resultMap)
	var agreement string
	if cmap, ok := resultMap["Colaborator"].(map[string]interface{}); ok {
		agreement = cmap["agreement"].(string)
	}
	if agreement == common.CONTRACTOR {
		member = buildContractor(resultMap)
	} else if agreement == string(common.EMPLOYEE) {
		member = buildEmployee(resultMap)
	}
	return member, nil
}

func buildContractor(resultMap map[string]interface{}) (member team.Member) {
	colaboratorMap := resultMap["Colaborator"].(map[string]interface{})
	duration := resultMap["duration"].(interface{})
	var tags []string
	tgs := colaboratorMap["tags"]
	s := reflect.ValueOf(tgs)
	for i := 0; i < s.Len(); i++ {
		tags = append(tags, s.Index(i).Elem().String())
	}
	member = &team.Contractor{
		Colaborator: team.Colaborator{
			ID:        colaboratorMap["id"].(primitive.ObjectID),
			Name:      colaboratorMap["name"].(string),
			Agreement: colaboratorMap["agreement"].(string),
			CreatedAt: int64(colaboratorMap["created_at"].(float64)),
			Tags:      tags,
		},
		Duration: int(duration.(float64)),
	}
	return
}

func buildEmployee(resultMap map[string]interface{}) (member team.Member) {
	colaboratorMap := resultMap["Colaborator"].(map[string]interface{})
	role := resultMap["role"].(interface{})
	var tags []string
	tgs := colaboratorMap["tags"]
	s := reflect.ValueOf(tgs)
	for i := 0; i < s.Len(); i++ {
		tags = append(tags, s.Index(i).Elem().String())
	}
	member = &team.Employee{
		Colaborator: team.Colaborator{
			ID:        colaboratorMap["id"].(primitive.ObjectID),
			Name:      colaboratorMap["name"].(string),
			Agreement: colaboratorMap["agreement"].(string),
			CreatedAt: int64(colaboratorMap["created_at"].(float64)),
			Tags:      tags,
		},
		Role: role.(string),
	}
	return
}
