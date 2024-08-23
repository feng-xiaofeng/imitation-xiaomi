package goods

type GoodsColor struct {
	Id         int
	ColorName  string
	ColorValue string
	Status     string
}

func (GoodsColor) TableName() string {
	return "models_color"
}
