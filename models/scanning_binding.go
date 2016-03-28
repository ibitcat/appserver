package models

// 发布扫红包
type SendScannigBinding struct {
	Barcode    string `json:"barcode" binding:"required" description:"条形码"`
	Itemname   string `json:"itemname" binding:"required" description:"商品名称"`
	Tag        string `json:"tag" binding:"required" description:"商品分类"`
	Desc       string `json:"desc" binding:"required" description:"促销信息"`
	Startdate  uint32 `json:"start_date" binding:"required" description:"红包发送日期"`
	Stopdate   uint32 `json:"stop_date" binding:"required" description:"红包结束日期"`
	TotalMoney uint32 `json:"total_money" binding:"required" description:"红包总金额"`
}
