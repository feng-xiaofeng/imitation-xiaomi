package menu

type Focus struct {
	Id        int
	Title     string
	FocusType int
	FocusImg  string
	Link      string
	Sort      int
	Status    int
	AddTime   int
}

func (Focus) TableName() string {
	return "focus"
}
