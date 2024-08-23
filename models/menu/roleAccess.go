package menu

type RoleAccess struct {
	AccessId int
	RoleId   int
}

func (RoleAccess) TableName() string {
	return "role_access"
}
