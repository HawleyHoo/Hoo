package main

import (
	"time"
	"fit"
	"fmt"
	"nursing/utils"
)

/**
datestr 入院时间
datestr2 手术时间
weekindex 第几周

## return
weeks 日期
dates1 住院日数
dates2 手术日数
weeknum 入院至今天一共多少周
 */
type hospDate struct {
	weeks   []time.Time // 日期, 这一周内有哪些天
	dates1  []string    // 住院日数
	dates2  []string    //手术日数
	weeknum int         // 入院至今天一共多少周
}

func getWeeksByOperationDateshy(hospitaldate, pid string, weekindex int) (hospdate hospDate, err error) {
	loc := time.Now().Location()
	hospitaldate = utils.Substr(hospitaldate, 0, 10)
	operationstr1 := ""
	operationstr2 := ""
	operationstr3 := ""

	//fmt.Println("hehe------", operationstr1, operationstr2)

	// 入院日期
	var t0 time.Time
	t0, err = time.ParseInLocation("2006-01-02", hospitaldate, loc)
	if err != nil {
		return
	}

	// 入院日期到今天的总周数
	weeknum := int(time.Since(t0).Hours()/24)/7 + 1
	hospdate.weeknum = weeknum

	offset := weekindex*7 - 7
	if weekindex == 0 {
		offset = weeknum*7 - 7
	}

	// 日期的第一天（第weekindex周的第一天）
	t1 := time.Date(t0.Year(), t0.Month(), t0.Day()+offset, 0, 0, 0, 0, loc)

	// 手术或产后日期
	operationTimes, err := model.FetchOperationRecordsDatehy(pid, t1)
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		return
	}
	// 手术后或产后日期
	var operatime1 time.Time
	var operatime2 time.Time
	//var operatime3 time.Time

	switch len(operationTimes) {
	case 0:
	case 1:
		operationstr1 = operationTimes[0]
	case 2:
		operationstr1 = operationTimes[0]
		operationstr2 = operationTimes[1]
	default:
		operationstr1 = operationTimes[0]
		operationstr2 = operationTimes[1]
		operationstr3 = operationTimes[2]
	}
	operationstr1 = utils.Substr(operationstr1, 0, 10)
	operationstr2 = utils.Substr(operationstr2, 0, 10)
	operationstr3 = utils.Substr(operationstr3, 0, 10)

	var weeks []time.Time
	var dates1, dates2 []string

	for i := 0; i < 7; i++ {
		t2 := t1.AddDate(0, 0, i)
		weeks = append(weeks, t2)
		dates1 = append(dates1, fmt.Sprintln(offset+i+1))

		//var operatime3 time.Time
		if operationstr1 != "" {
			operatime1, err = time.ParseInLocation("2006-01-02", operationstr1, loc)
			if err != nil {
				operatime1 = time.Now().AddDate(1, 0, 0)
				//fmt.Println("------ hos :", operatime1, err)
			}

			if operationstr2 != "" {
				operatime2, err = time.ParseInLocation("2006-01-02", operationstr2, loc)
				if err != nil {
					operatime2 = time.Now().AddDate(1, 0, 0)
					//fmt.Println("------ hos :", operatime2, err)
				}
			}
			//if operationstr3 != "" {
			//	operatime3, err = time.ParseInLocation("2006-01-02", operationstr3, loc)
			//	if err != nil {
			//		operatime3 = time.Now().AddDate(1, 0, 0)
			//		fmt.Println("------ hos :", operatime3, err)
			//	}
			//}
			//fmt.Println("success", operatime1, operatime2)

			operaoffset1 := t2.Sub(operatime1).Hours()

			difftime1 := operatime2.Sub(operatime1).Hours()
			//difftime2 := operatime3.Sub(operatime2).Hours()
			//fmt.Println("offset :", operaoffset1, difftime1)

			if operaoffset1 >= 0 && operaoffset1 < 24*10 { // 手术十天以内记录时间
				if difftime1 < 0 || difftime1 > 24*10 {
					operstr := fmt.Sprintf(" %.0f", operaoffset1/24)
					fmt.Println("--------------------", operstr)
					dates2 = append(dates2, operstr)
				} else { // 十天以内做过第二次手术 则记为 Ⅱ-1
					operaoffset2 := t2.Sub(operatime2).Hours()
					if operaoffset2 >= 0 && operaoffset2 < 24*10 {
						aaa := int(operaoffset2 / 24)
						dates2 = append(dates2, fmt.Sprintf("Ⅱ-%d", aaa))
					} else {
						dates2 = append(dates2, fmt.Sprintf(" %.0f", operaoffset1/24))
					}
				}
			} else {
				str := " "
				if difftime1 < 24*10 {
					str = "Ⅱ-"
				}
				operaoffset2 := t2.Sub(operatime2).Hours()
				if operaoffset2 >= 0 && operaoffset2 <= 24*10 {
					aaa := int(operaoffset2 / 24)
					dates2 = append(dates2, fmt.Sprintf("%s%d", str, aaa))
				} else {
					dates2 = append(dates2, "")
				}
			}
		} else {
			for ii := 0; ii < 7; ii++ {
				dates2 = append(dates2, "")
			}
			operatime1 = time.Now().AddDate(1, 0, 0)
			//fmt.Println("hos :", operatime1)
		}

	}
	fmt.Println("date2 :", dates2, len(dates2))
	hospdate.weeks = weeks
	hospdate.dates1 = dates1
	hospdate.dates2 = dates2
	return hospdate, nil
}