package team

type RequestMember struct {
	RepoType string `json:"RepoType"`
	Member   Member `json:"Member"`
}
