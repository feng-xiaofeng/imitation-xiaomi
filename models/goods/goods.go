package goods

type Goods struct {
	Id            int
	Title         string
	SubTitle      string
	GoodsSn       string
	CateId        int
	ClickCount    int
	GoodsNumber   int
	ShopPrice     float64
	MarketPrice   float64
	RelationGoods string
	GoodsAttr     string
	GoodsVersion  string
	GoodsImg      string
	GoodsGift     string
	GoodsFitting  string
	GoodsColor    string
	GoodsKeywords string
	GoodsDese     string
	GoodsContent  string
	IsDelete      int
	IsHot         int
	IsBest        int
	IsNew         int
	GoodsTypeId   int
	Sort          int
	Status        int
	AddTime       int
}

func (Goods) TableName() string {
	return "goods"
}
