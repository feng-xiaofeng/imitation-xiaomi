package goods

type GoodsAttr struct {
	Id              int
	GoodsId         int
	CateId          int
	AttributeCateId int
	AttributeId     int
	AttributeType   int
	AttributeTitle  string
	AttributeVale   string
	AttributeTime   int
	Sort            int
	Status          int
}

func (GoodsAttr) TableName() string {
	return "goods_attr"
}
