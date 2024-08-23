package admin

import (
	"fmt"
	"mygo_demo/models"
	good "mygo_demo/models/goods"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{})
}
func (con GoodsController) Add(c *gin.Context) {
	//获取商品分类
	goodsCateList := []good.GoodsCate{}
	models.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	//获取所有颜色信息
	goodsColer := []good.GoodsColor{}
	models.DB.Find(&goodsColer)

	//获取商品规格包装
	goodsTypeList := []good.GoodsType{}
	models.DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{
		"goodsCateList": goodsCateList,
		"goodsColer":    goodsColer,
		"goodsTypeList": goodsTypeList,
	})
}
func (con GoodsController) DoAdd(c *gin.Context) {
	// attrIdList := c.PostFormArray("attr_id_list") //获取数组
	// attrValueList := c.PostFormArray("attr_value_list")
	// goodsImageList := c.PostFormArray("goods_image_list")
	// c.JSON(http.StatusOK, gin.H{
	// "attrIdList":     attrIdList,
	// "attrValueList":  attrValueList,
	// "goodsImageList": goodsImageList,
	// })

	//1、获取表单提交过来的数据 进行判断
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	goodsVersion := c.PostForm("goods_version")
	cateId, cateIdErr := models.Int(c.PostForm("cate_id"))
	goodsNumber, _ := models.Int(c.PostForm("goods_number"))
	//注意小数点
	price, priceErr := models.Float(c.PostForm("price"))
	marketPrice, marketPriceErr := models.Float(c.PostForm("market_price"))
	goodsContent := c.PostForm("goods_content")
	isBest, isBestErr := models.Int(c.PostForm("is_best"))
	isHot, isHotErr := models.Int(c.PostForm("is_hot"))
	isNew, isNewErr := models.Int(c.PostForm("is_new"))

	//获取的是切片
	goodsColor := c.PostFormArray("goods_color")
	//
	relationGoods := c.PostForm("relation_goods")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	goodsAttr := c.PostForm("goods_attr")
	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	isDelete, _ := models.Int(c.PostForm("is_delete"))
	goodsTypeId, goodsTypeIdErr := models.Int(c.PostForm("goods_type_id"))
	sort, _ := models.Int(c.PostForm("sort"))
	status, statusErr := models.Int(c.PostForm("status"))
	addTime := int(models.GetUnix())
	//2、获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColor, ",")
	//3、上传图片   生成缩略图
	goodsImg, goodsImgErr := models.UploadImg(c, "goods_img")
	if goodsImgErr != nil {
		con.Error(c, "上传图片失败", "/admin/goods/add")
		return
	}
	if cateIdErr != nil || priceErr != nil || marketPriceErr != nil || isBestErr != nil || isHotErr != nil || isNewErr != nil || goodsTypeIdErr != nil || statusErr != nil {
		con.Error(c, "上传数据失败", "/admin/goods/add")
		return
	}

	// goodsImgObj := goods.GoodsImage{}
	//4、增加商品数据
	goods := good.Goods{
		Title:         title,
		SubTitle:      subTitle,
		GoodsSn:       goodsSn,
		CateId:        cateId,
		ClickCount:    100,
		GoodsNumber:   goodsNumber,
		ShopPrice:     price,
		MarketPrice:   marketPrice,
		RelationGoods: relationGoods,
		GoodsAttr:     goodsAttr,
		GoodsVersion:  goodsVersion,
		GoodsImg:      goodsImg,
		GoodsGift:     goodsGift,
		GoodsFitting:  goodsFitting,
		GoodsColor:    goodsColorStr,
		GoodsKeywords: goodsKeywords,
		GoodsDese:     goodsDesc,
		GoodsContent:  goodsContent,
		IsDelete:      isDelete,
		IsHot:         isHot,
		IsBest:        isBest,
		IsNew:         isNew,
		GoodsTypeId:   goodsTypeId,
		Sort:          sort,
		Status:        status,
		AddTime:       addTime,
	}

	err := models.DB.Create(&goods).Error
	if err != nil {
		con.Error(c, "增加失败", "/admin/goods/add")
	}
	//5、增加图库 信息
	wg.Add(1)
	go func() {
		goodsImageList := c.PostFormArray("goods_image_list")
		for _, v := range goodsImageList {
			goodsImgObj := good.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()

	//6、增加规格包装
	wg.Add(1)
	go func() {
		attrIdList := c.PostFormArray("attr_id_list") //获取数组
		attrValueList := c.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, attributeErr := models.Int(attrIdList[i])
			if attributeErr != nil {
				fmt.Println("获取数据失败", attributeErr.Error())
				return
			}
			// 获取商品类型属性数据
			goodsTypeAttributeObj := good.GoodsTypeAttribute{Id: goodsTypeAttributeId}
			models.DB.Find(&goodsTypeAttributeObj)
			// 给商品属性里增加数据 规格包装
			goodsAttrObj := good.GoodsAttr{}
			goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
			goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
			goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
			goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
			goodsAttrObj.AttributeVale = attrValueList[i]
			goodsAttrObj.Status = 1
			goodsAttrObj.Sort = 10
			goodsAttrObj.AttributeTime = int(models.GetUnix())
			models.DB.Create(&goodsAttrObj)
		}
		wg.Done()
	}()

	con.Success(c, "增加数据成功", "/admin/goods/add")
}

func (con GoodsController) TypeAttribute(c *gin.Context) {
	cateId, err1 := models.Int(c.Query("cateId"))
	goodsTypeAttributeList := []good.GoodsTypeAttribute{}
	err2 := models.DB.Where("cate_id=?", cateId).Find(&goodsTypeAttributeList).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  goodsTypeAttributeList,
		})
	}
}
func (con GoodsController) ImageUpload(c *gin.Context) {
	//上传图片
	imgDir, err := models.UploadImg(c, "file") //注意：可以在网络里面看到传递的参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"link": "/" + imgDir,
		})
	}
}
