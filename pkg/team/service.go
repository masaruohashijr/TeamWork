package team

type MemberService interface {
	GetAll() ([]interface{}, error)
	Get(name string) (interface{}, error)
	Create(member interface{}) error
	Update(member interface{}) error
	Delete(member interface{}) error
}
