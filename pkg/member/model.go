package member

// structs
type Member struct {
	Code      string   `json:"code"`
	CreatedAt int64    `json:"created_at"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
}

type Employee struct {
	A_member Member
	Role     string
}

type Contractor struct {
	A_member Member
	Duration int
}
