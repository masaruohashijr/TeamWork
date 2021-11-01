package team

type MemberRepository interface {
	DbGetAll() ([]interface{}, error)
	DbGet(name string) (interface{}, error)
	DbCreate(member interface{}) error
	DbUpdate(member interface{}) error
	DbDelete(member interface{}) error
}
