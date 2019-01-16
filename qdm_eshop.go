package main

import (
	"time"
	"eshop/model"
	"github.com/Luxurioust/excelize"
	"fmt"
	"eshop/util"
	"strconv"
	//"github.com/tealeg/xlsx"
	"os"
	"eshop/service"
)

type Hoo struct {
	SkuID string `json:"sku_id"`
	skuName string `json:"sku_name"`
	Qty int `json:"qty"`
}

func groupby(arr1, arr2 []model.ScheOrder) (res []model.ScheOrder)  {
	g1 := make(map[string]model.ScheOrder)
	for _, v := range arr1 {
		g1[v.StoreID] = v
	}
	g2 := make(map[string]model.ScheOrder)
	for _, v := range arr2 {
		g2[v.StoreID] = v
	}

	for k,v := range g1 {
		v2, ok := g2[k]
		if ok {
			v.PurchaseQTY += v2.PurchaseQTY
		} else {
			util.LogInfo(" no match storeid %s %s", v.StoreID, v.StoreName)
		}
		res = append(res, v)
	}
	return
}

func accumulatePurchaseQty(source []model.ScheduleSKU) (result []model.ScheduleSKU) {
	g1 := make(map[string]model.ScheduleSKU)

	for _ , v := range source {
		id := v.SkuID[:8]
		fmt.Println("id", id)
		h, ok := g1[id]
		if ok {
			if h.SkuID == id {
				h.TotalQTY += v.TotalQTY
				h.OrderList = groupby(h.OrderList, v.OrderList)
				g1[id] = h

			} else {
				v.TotalQTY += h.TotalQTY
				v.OrderList = groupby(v.OrderList, h.OrderList)
				g1[id] = v
			}
		} else {
			g1[id] = v
		}
	}

	for _, v := range g1 {
		//fmt.Println("group:", v.SkuID, v.SkuName, v.TotalQTY, v.OrderList)
		result = append(result, v)
	}
	return result
}

func main() {
	var n int64 = 500
	fmt.Println("n", int(n))
	return
	//for  {
	//	fmt.Println("1")
	//	fmt.Println("2")
	//	fmt.Println("--------")
	//	time.Sleep(time.Second)
	//}
	var sum int64 = 0
	var ii uint
	for ii = 0; ii <= 36;ii++ {
		v := 2 << ii
		sum += int64(v)
	}
	fmt.Printf("2^365: %d .\n", sum)
	fmt.Println("2^365: ", 0xffff)

	for i := 0;i < 10 ; i ++ {
		//defer func(i int) {
		//	fmt.Println("hehe", i)
		//}(i)
		switch i {
		case 1, 6, 8,9:
			{
				switch i {
				case 8:
					fmt.Println("aaaa", 8)
				case 9:
					fmt.Println("aaaa", 9)

				}
			}

			//fmt.Println("switch", i)
		default:
			continue
		}


		fmt.Println("-----------", i)
	}


	//h1 := Hoo{"202", "文昌蛋", 10}
	//h3 := Hoo{"203", "武昌蛋", 12}
	//h2 := Hoo{"202_2", "文昌蛋", 24}
	//arr := []Hoo{h1, h2, h3}
	orglist := service.GetSchSkuList(time.Now())
	for i, v := range orglist {
		fmt.Println(i," org:", v.SkuID, v.SkuName, v.TotalQTY, v.OrderList)
	}
	//g1 := make(map[string]Hoo)

	res := accumulatePurchaseQty(orglist)
	for _, v := range res {
		fmt.Println("org:", v.SkuID, v.SkuName, v.TotalQTY, v.OrderList)
		//fmt.Printf("group %+v \n", v)
	}

	return
	//jsonstr, err  := json.Marshal(arr)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//data := []Hoo{}
	//err = json.Unmarshal([]byte(jsonstr), &data)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//fmt.Println("res", data)
	//mm := make(map[int]string)
	//mm[20] = "hehe"
	//fmt.Println("hehe", mm[20])
	//fmt.Println("hehe", mm[30])


	//ticker:=time.NewTicker(time.Second*1)
	//go func() {
	//	for _ = range ticker.C {
	//		fmt.Println("test")
	//	}
	//}()
	//
	//time.Sleep(time.Minute)

	//model.SchAddWarningErrorRecord("test", "teset")
	//service.FetchCurStoreTask()
	//service.FetchCurStoreStock()
	//res, err := service.YouzanItemsGet("440046657")
	//if res.IsListing == true && res.SoldNum > 950 {
	//	service.PushSMSOnce()
	//err := service.PushSMSOnce("18588726982", "三文鱼头线上销量已超过950，请速下线处理。")
	//if err != nil {
	//
	//}
	//}
	//fmt.Printf("res :%+v", res)
	//fmt.Println(err)

}

func schplan() {
	strs := []string{"1", "2", "3", "4"}
	for _, v := range strs {
		fmt.Sprintln(v)
	}
	fmt.Println("hehe", len(" "))

	pkuqty, err := util.ConvertSkuQtyToPkuQTY(16, 200, "kg", "kg", "n")
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("pku qty", pkuqty)
	return
	//model.SchAddWarningErrorRecord("getSkuInfoErr", "logstr")
	//fmt.Println(time.Now())
	//fmt.Println(time.Now().Local())
	//return
	// 查询 总仓库数量
	whList, count, err := model.GetWarehouseList("")
	//for _,wh := range whList {
	//	//util.LogInfo("%+v",wh)
	//	storeList, err := model.SchGetStoreInfo(wh.Code, "")
	//	//util.LogInfo("%+v, %v",storeList, err)
	//	fmt.Println(len(storeList), err)
	//	writeSchPlanFile(storeList)
	//}
	fmt.Println("count", count, err)
	storeList, err := model.SchGetStoreInfo(whList[0].Code, "")
	fmt.Println(len(storeList), err)
	writeSchPlanFile(storeList)
}

func writeSchPlanFile(storelist []model.StoreInfo) {
	headerX := map[string]string{}
	content := map[string]string{}
	t0 := time.Now().AddDate(0, 0, 1)

	headerX["A1"] = "下单日期：" + t0.Format("2006-01-02")
	headerX["C1"] = "到店日期：" + t0.AddDate(0, 0, 1).Format("2006-01-02")

	//headerX2 := map[string]string{}
	headerX["A2"] = "货号"
	headerX["B2"] = "品名"
	headerX["C2"] = "订单总数"
	headerX["D2"] = "单位"

	storeMap := map[string]string{}

	// 第一行  第二行， 门店id，门店名称
	x0 := 5
	for k, v := range storelist {
		index := convertToTitle2(x0 + k)
		headerX[index+"1"] = v.ID
		headerX[index+"2"] = v.Name + " " + v.ProductionCode

		//fmt.Println("key:", index, "value:", headerX[title])
		storeMap[v.ID] = index
	}

	// 列  sku表头
	skuList := getSchSkuList(t0)
	for row, v := range skuList {
		rowStr := strconv.Itoa(row + 3)
		//f1 :=  "A" + rowStr
		headerX["A"+rowStr] = v.SkuID
		headerX["B"+rowStr] = v.SkuName
		headerX["C"+rowStr] = strconv.Itoa(v.TotalQTY)
		headerX["D"+rowStr] = "盒"

		for _, s := range v.OrderList {
			key := storeMap[s.StoreID]
			content[key+rowStr] = strconv.Itoa(s.PurchaseQTY)
		}
	}

	excelFileName := "./test.xlsx"
	xlsx, err := excelize.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sheet := "Sheet1"
	//xlsx := excelize.NewFile()
	for k, v := range headerX {
		xlsx.SetCellValue(sheet, k, v)
	}
	for k, v := range content {
		xlsx.SetCellValue(sheet, k, v)
	}
	xlsx.MergeCell(sheet, "A1", "B1")
	xlsx.MergeCell(sheet, "C1", "D1")
	//xlsx.SetCellStyle()
	// xlsx.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`)
	// Save xlsx file by the given path.
	err = xlsx.SaveAs("./test.xlsx")
	if err != nil {
		fmt.Println(err)
	}

}

func getSchSkuList(t0 time.Time) []model.ScheduleSKU {
	//  GetScheduleSkuOrderList
	//  查询上线的SKU
	skuList, count, err := model.ScheduleGetLiveSkuList("", 0, 0, 0)
	if err != nil {
		util.LogError("ScheduleGetLiveSkuList: %s", err.Error())
		return nil
	}
	util.LogInfo("schedule plan push to pos,store count:%d  %d", count, len(skuList))
	//now := time.Now()
	today := t0.Format("20060102")

	result := []model.ScheduleSKU{}
	// 根据sku和date 获取每个门店的订货数
	for _, v := range skuList {
		if v.IsDaoFenHuoBiao == false {
			continue
		}
		//targetdate := now.AddDate(0, 0, 1).Format("20060102")
		res, err := model.GetScheduleSkuOrderList(v.ID, today, v.SkuID)
		if err != nil {
			//util.LogInfo("GetScheduleSkuOrderList: %s", err.Error())
			continue
		}
		result = append(result, res)
		//util.LogInfo("row %d %+v, %s", row, res, targetdate)

	}
	return result
}

func convertToTitle2(n int) (t string) {
	for n > 0 {
		t = fmt.Sprintf("%c", (n-1)%26+'A') + t
		n = (n - 1) / 26
	}
	//t = fmt.Sprintf("%c", n+'A') + t
	//return h_pkg.Reverse(t)
	return t
}

// letterOnlyMapF is used in conjunction with strings.Map to return only the
// characters A-Z and a-z in a string.
func letterOnlyMapF(rune rune) rune {
	switch {
	case 'A' <= rune && rune <= 'Z':
		return rune
	case 'a' <= rune && rune <= 'z':
		return rune - 32
	}
	return -1
}

// intOnlyMapF is used in conjunction with strings.Map to return only the
// numeric portions of a string.
func intOnlyMapF(rune rune) rune {
	if rune >= 48 && rune < 58 {
		return rune
	}
	return -1
}
