package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
	"strconv"
)

func testReadFile() {
	//excelFileName := "./h_tool/Workbook.xlsx"
	excelFileName := "./h_tool/orderlist.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	headers := []string{"序号", "门店名称", "门店编号", "品名", "PKU编码", "pku采购单位", "销售单位", "订货单价（元）", "订货数量", "需求配货时间", "备注"}
	for _, sheet := range xlFile.Sheets {
		fmt.Println("name:", sheet.Name, "rows:", len(sheet.Rows), "max row:", sheet.MaxRow, sheet.MaxCol)
		if len(sheet.Rows) == 1 {
			return // sheet 为空
		}
		for i, row := range sheet.Rows {
			fmt.Printf("row: %d :", i)
			for j, cell := range row.Cells {
				text := cell.String()
				if i == 0 {
					if text != headers[j] {
						return // 数据错误
					}
				}
				fmt.Printf("%s\t", text)
			}
			fmt.Printf("\n")
		}
	}
}

func main() {
	testReadFile()
	//readFile()
}

func readFile() {
	//获取当前目录，类似linux中的pwd
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(workPath)
	fmt.Println(filepath.Abs("./h_tool/"))

	xlsx, err := excelize.OpenFile("./h_tool/Workbook.xlsx")
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
