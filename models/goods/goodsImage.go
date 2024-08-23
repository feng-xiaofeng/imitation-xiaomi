package goods

type GoodsImage struct {
	Id      int
	GoodsId int
	ImgUrl  string
	Sort    int
	ColorId int
	AddTime int
	Status  int
}

func (GoodsImage) TableName() string {
	return "models_image"
}
