package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-interview-project-masaru-ohashi/cmd/client/errors"
	"golang-interview-project-masaru-ohashi/cmd/client/ports"
	"golang-interview-project-masaru-ohashi/pkg/team"
	"io/ioutil"
	"net/http"
	"net/url"
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
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf(string(body))
	}
	errors.CheckErrorMember(err, errHandler)
	err = json.Unmarshal(body, &newMember)
	errors.CheckErrorMember(err, errHandler)
	return newMember, nil
}

func (a *memberAPI) Put(member interface{}) error {
	// apiEndpoint := a.apiUrl + "/Member"
	return nil
}

func (a *memberAPI) Delete(m interface{}) error {
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
	return nil
}

func (a *memberAPI) GetAll() (members []interface{}, err error) {
	apiEndpoint := a.apiUrl + "/Members"
	request := &team.RequestMember{
		RepoType: a.repoType,
	}
	requestBody, err := json.Marshal(request)
	var buffer *bytes.Buffer = bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("GET", apiEndpoint, buffer)
	if err != nil {
		return members, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return members, err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf(string(body))
	}
	err = json.Unmarshal(body, &members)
	errors.CheckErrorMember(err, errHandler)
	b := string(body)
	fmt.Println("response Body:", b)
	json.Unmarshal(body, &members)
	return members, nil
}

func (a *memberAPI) Get(name string) (member interface{}, err error) {
	apiEndpoint := a.apiUrl +
		"/Member"
	values := &url.Values{}
	values.Add("Name", name)
	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("GET", apiEndpoint, buffer)
	switch v := member.(type) {
	case *team.Contractor:
		*v = team.Contractor{}
	case *team.Employee:
		*v = team.Employee{}
	}
	if err != nil {
		return member, err
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
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
	err = json.Unmarshal(body, &member)
	if err != nil {
		return member, err
	}
	return member, nil
}
