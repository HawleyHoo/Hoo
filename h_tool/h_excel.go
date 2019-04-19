package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
	"strconv"
	"eshop/model"
	"time"
)

type Store struct {
	ID string
	Name string
}

func testReadFile() {
	//excelFileName := "./h_tool/Workbook.xlsx"
	excelFileName := "./h_tool/20181121.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//headers := []string{"序号", "门店名称", "门店编号", "品名", "PKU编码", "pku采购单位", "销售单位", "订货单价（元）", "订货数量", "需求配货时间", "备注"}
	for i, sheet := range xlFile.Sheets {
		if i == 1 {
			return
		}
		fmt.Println("name:", sheet.Name, "rows:", len(sheet.Rows), "max row:", sheet.MaxRow, "MaxCol", sheet.MaxCol)
		if len(sheet.Rows) == 1 || sheet.MaxCol < 3 {
			return // sheet 为空
		}
		// 取sku
		skuOrders := []model.WSKUOrder{}
		for i, row := range sheet.Rows {
			fmt.Printf("row index: %d :", i + 1)

			if i > 1 {
				skuid := row.Cells[0].String()
				skuname := row.Cells[1].String()
				if len(skuid) == 0 || len(skuname) == 0 {
					break
				}
				sku := model.WSKUOrder{
					ID: "123456789",
					SKUID:skuid,
					SKUName:skuname,
				}
				fmt.Println(sku)
				skuOrders = append(skuOrders, sku)
			}

		}
		fmt.Println("********************* 取门店列表 *****************************")
		// 取门店列表
		storeList := []model.WSStore{}
		row0 := sheet.Rows[0]
		row1 := sheet.Rows[1]
		length := len(row0.Cells)
		for i := 4; i < length; i++ {

			store := model.WSStore{
				StoreID:row0.Cells[i].String(),
				StoreName:row1.Cells[i].String(),
			}
			//fmt.Println(store)
			storeList = append(storeList, store)
		}
		fmt.Println(storeList)
		/*for i, row := range sheet.Rows {
			fmt.Printf("row index: %d :", i + 1)
			for j, cell := range row.Cells {
				text := cell.String()
				if j > 3 {
					store := model.WSStore{}
					switch i {
					case 0:
					case 1:
					default:
					}
					if i == 0 {
						store.StoreID = text
					} else if i == 1 {
						store.StoreName = text
					} else {
						qty, err := cell.Int()
						if err != nil {
							//fmt.Println(err)
							fmt.Printf("(%d-%d: %s)\t", i + 1, j + 1, text)
							qty = 0
						}
						store.PurchaseQTY = qty
					}
					storeList = append(storeList, store)
				}
				fmt.Printf("(%d-%d: %s)\t", i + 1, j + 1, text)
			}
			fmt.Printf("\n")

			if i > 1 {
				skuid := row.Cells[0].String()
				skuname := row.Cells[1].String()
				if len(skuid) == 0 || len(skuname) == 0 {
					break
				}
			}
		}*/

		fmt.Println("********************* 取门店订货量 *****************************")
		fmt.Println("sku list:", len(skuOrders), "store list:", len(storeList))
		// 取门店订货量
		//errbuffer := bytes.Buffer{}
		for i := 2; i < len(skuOrders) + 2; i++ {
			row := sheet.Rows[i]
			sku := skuOrders[i - 2]
			for j := 4; j < len(storeList) + 4; j++  {
				store := storeList[j - 4]
				cell := row.Cells[j]
				qty, err := cell.Int()
				if err != nil {
					//errbuffer.WriteString()
					//fmt.Printf("(第%d行-第%d列: %s)\t", i + 1, j + 1, err.Error())
					qty = 0
				}
				s := model.WSStore{
					StoreID:store.StoreID,
					StoreName:store.StoreName,
					PurchaseQTY:qty,
				}
				sku.OrderList  = append(sku.OrderList, s)
				//fmt.Println(sku.OrderList)
				//fmt.Printf("(%d-%d: %d)\t", i + 1, j + 1, qty)
			}
			skuOrders[i - 2] = sku
		}

		fmt.Println("************************ result **********************")
		fmt.Println(skuOrders)
	}
}

func exceltest()  {
	//excelFileName := "./h_tool/Workbook.xlsx"
	excelFileName := "./Book1.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//headers := []string{"序号", "门店名称", "门店编号", "品名", "PKU编码", "pku采购单位", "销售单位", "订货单价（元）", "订货数量", "需求配货时间", "备注"}
	for _, sheet := range xlFile.Sheets {
		fmt.Println("name:", sheet.Name, "rows:", len(sheet.Rows), "max row:", sheet.MaxRow, "MaxCol", sheet.MaxCol)
		if len(sheet.Rows) == 1 {
			return // sheet 为空
		}

		for i, row := range sheet.Rows {
			fmt.Printf("row: %d :", i)
			for _, cell := range row.Cells {
				text := cell.String()
				if i == 0 {
					//if text != headers[j] {
					//	return // 数据错误
					//}
				}
				fmt.Printf("%s\t", text)
			}
			fmt.Printf("\n")
		}
	}
}

func main() {

	fmt.Println(time.Now().Format("20060102150405.999"))
	fmt.Println(time.Now().Format("20060102150405000"))
	fmt.Println(time.Now().Format("20060102150405.999"))
	//a := 2
	//b := 0
	//c := a / b
	//fmt.Println(c)

	//testReadFile()
	//readFile()

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
	exceltest()
}

func writeSchPlanFile(storelist []model.StoreInfo)  {
	headerX := map[string]string{}
	headerX["A1"] = "下单日期：" + time.Now().Format("2006-01-02")
	headerX["C1"] = "到店日期：" + time.Now().AddDate(0,0, 1).Format("2006-01-02")

	//headerX2 := map[string]string{}
	headerX["A2"] = "货号"
	headerX["B2"] = "品名"
	headerX["C2"] = "订单总数"
	headerX["D2"] = "单位"

	x0 := 5
	for k,v := range storelist {
		index := convertToTitle2(x0 + k)
		headerX[index + "1"] = v.ID
		headerX[index + "2"] = v.Name + " " + v.ProductionCode

		//fmt.Println("key:", index, "value:", headerX[title])
	}

	xlsx := excelize.NewFile()
	for k, v := range headerX {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	//for k, v := range values {
	//	xlsx.SetCellValue("Sheet1", k, v)
	//}
	// xlsx.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

}


func convertToTitle2(n int) (t string) {
	for n > 0 {
		t = fmt.Sprintf("%c", (n - 1)%26 + 'A') + t
		n = (n - 1) / 26
	}
	//t = fmt.Sprintf("%c", n+'A') + t
	//return h_pkg.Reverse(t)
	return t
}

func readFile() {
	//获取当前目录，类似linux中的pwd
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(workPath)
	fmt.Println(filepath.Abs("./h_tool/"))

	xlsx, err := excelize.OpenFile("./h_tool/20181121.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get value from cell by given sheet index and axis.
	cell := xlsx.GetCellValue("hehe", "B2")
	fmt.Println(cell)
	// Get sheet index.
	index := xlsx.GetSheetIndex("hehe")
	fmt.Println("index:", index)
	// Get all the rows in a sheet.
	rows2 := xlsx.GetRows("sheet" + strconv.Itoa(index))

	fmt.Println(rows2, xlsx.GetActiveSheetIndex())

	fmt.Println("----------------------------")
	rows := xlsx.GetRows("hehe")

	//xlsx.GetSheetIndex()
	//fmt.Println(rows)
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println(row)
	}
}

func writeFile() {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	xlsx := excelize.NewFile()
	for k, v := range categories {

		xlsx.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	// xlsx.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

/*func writeToFile()  {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	xlsx := excelize.CreateFile()

	for k, v := range categories {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	xlsx.AddChart("Sheet1", "E1", `{"type":"bar3D","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Line Chart"}}`)
	// Save xlsx file by the given path.
	err := xlsx.WriteTo("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}*/

func InsertPic() {
	xlsx, err := excelize.OpenFile("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Insert a picture.
	err = xlsx.AddPicture("Sheet1", "A2", "./image1.gif", "")
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture to sheet with scaling.
	err = xlsx.AddPicture("Sheet1", "D2", "./image2.jpg", `{"x_scale": 0.5, "y_scale": 0.5}`)
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	err = xlsx.AddPicture("Sheet1", "H2", "./image3.gif", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`)
	if err != nil {
		fmt.Println(err)
	}
	// Save the xlsx file with the origin path.
	err = xlsx.Save()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
