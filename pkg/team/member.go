package team

// structs
type Member interface {
	GetName() string
	GetAgreement() string
	GetCreatedAt() int64
	CreatedAt(int64)
	GetTags() []string
	String() string
}
