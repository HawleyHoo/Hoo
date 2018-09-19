package main

import (
	"time"
	"fit"
	"fmt"
	"nursing/utils"
	"strconv"
	"database/sql"
)

type FitTime time.Time

type HAAAA struct {
	ID             int64   // 主键ID
	PatientId      int64   // 病人ID
	PatientName    string  // 姓名
	VisitId        int64   // 就诊ID
	Bed            string  // 床号
	NurseId        int     // 责任护士ID
	NurseName      string  // 责任护士ID
	DepartmentId   int     // 病区ID(科室)
	DepartmentName string  // 病区
	HospNum        string  // 住院号
	Age            string  // 年龄
	Category       string  // 病人类别
	Gender         string  // 性别
	Physician      string  // 住院医师
	NursingDegree  string  // 护理级别
	NursingDegreeV string  // 护理级别
	Diagnosis      string  // 诊断
	HospTime       FitTime // 入院时间
	DateTime       FitTime // 测量时间
	NewOrder       int     // 新医嘱 ？
	StoppedOrder   int     // 已停医嘱 ？
	IsFever        bool    // 是否发热（最后一次测量体温>37.5°）
	IsOperation    bool    // 是否待手术 ？
	IsArrearage    bool    // 是否欠费
	IsNewPatient   bool    // 是否是新病人
}


var timeChina = []string{"", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
var dateChina = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}

func sinicizingDate(intval int) {
	for index := len(strconv.Itoa(intval)); index >= 0; index++ {
		//ii := intval % 10
	}
}

func initEngine() (*xorm.Engine, error) {
	db, err := xorm.NewEngine("mysql", "phpgroup:fitcome_meal1!qw2@tcp(39.108.133.131:1714)/nursing?charset=utf8")
	//db, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3307)/nuring?charset=utf8")
	db.TZLocation = time.Now().Location()
	db.DatabaseTZ = time.Now().Location() // Now().Location()
	//SnakeMapper 支持struct为驼峰式命名，表结构为下划线命名之间的转换，这个是默认的Maper
	//映射同名设置默认
	db.SetMapper(core.SameMapper{})
	fmt.Println("------hahaha")
	if err == nil {
		return db, err
	}
	return nil, err
}

func initmysqlEngine() (*sql.DB, error) {
	db, err := sql.Open("mysql", "phpgroup:fitcome_meal1!qw2@tcp(39.108.133.131:1714)/nursing?charset=utf8")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sinicizingTime(intval int) (datestr string) {
	if intval >= 0 && intval <= 60 {
		if intval == 0 {
			datestr = "零"
		} else if intval < 10 {
			datestr = timeChina[intval]
		} else {
			datestr = timeChina[intval/10] + "十" + timeChina[intval%10]
		}
	} else {
		datestr = ""
	}
	return
}

func testttt()  {
	for index := 0; index < len(operationTimes)-1; index++ {
		t1 = operationTimes[index]
		t2 = operationTimes[index+1]

		offset := int(t2.Sub(t1).Hours() / 24)
		//fmt.Println("offset:", t1, t2, "offset day:", offset, t2.Sub(t1).Hours())
		//if t2.Sub(t1).Hours() <= 24*10 {

		//fmt.Println("index:", index, "flag:",flag, "lable:", dateLables[flag - 1])

		for ii := 0; ii < offset; ii++ {
			lable = dateLables[flag]
			str := fmt.Sprintf("%s%d", lable, ii)
			if ii > 10 {
				str = ""
			}
			results = append(results, str)
			//fmt.Println("flag:", flag, lable, "str:", str, t1.AddDate(0, 0, ii))
		}
		if offset <= 10 {
			flag++
		} else {
			flag = 1
		}
		/*if index == len(operationTimes) - 2 {
			flag++
		} else {
		}*/
	}
}

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

func GetOperationRecords(operationTimes []time.Time)  {

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
	// 向前数10天，向后数7天
	operationTimes, err := model.FetchOperationRecordsDatehy(pid, t1)
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		return
	}
	// 手术后或产后日期
	var operatime1 time.Time
	var operatime2 time.Time
	var operatime3 time.Time

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

		if len(operationTimes) == 1 {
			operatime1, err = time.ParseInLocation("2006-01-02", operationstr1, loc)
			if err != nil {
				operatime1 = time.Now().AddDate(1, 0, 0)
				//fmt.Println("------ hos :", operatime1, err)
			}
			operaoffset1 := t2.Sub(operatime1).Hours()
			if operaoffset1 >= 0 && operaoffset1 < 24*10 { // 手术十天以内记录时间
				operStr := fmt.Sprintf(" %.0f", operaoffset1/24)
				dates2 = append(dates2, operStr)
			} else {
				dates2 = append(dates2, "")
			}

		} else if len(operationTimes) == 2 {
			operatime1, err = time.ParseInLocation("2006-01-02", operationstr1, loc)
			if err != nil {
				operatime1 = time.Now().AddDate(1, 0, 0)
				//fmt.Println("------ hos :", operatime1, err)
			}
			operatime2, err = time.ParseInLocation("2006-01-02", operationstr2, loc)
			if err != nil {
				operatime2 = time.Now().AddDate(1, 0, 0)
				//fmt.Println("------ hos :", operatime2, err)
			}

			operaoffset1 := t2.Sub(operatime1).Hours()
			difftime1 := operatime2.Sub(operatime1).Hours()
			//difftime2 := operatime3.Sub(operatime2).Hours()
			fmt.Println("offset :", operaoffset1, difftime1)

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
			operatime1, err = time.ParseInLocation("2006-01-02", operationstr1, loc)
			if err != nil {
				operatime1 = time.Now().AddDate(1, 0, 0)
				//fmt.Println("------ hos :", operatime1, err)
			}
			operatime2, err = time.ParseInLocation("2006-01-02", operationstr2, loc)
			if err != nil {
				operatime2 = time.Now().AddDate(1, 0, 0)
				//fmt.Println("------ hos :", operatime2, err)
			}
			operatime3, err = time.ParseInLocation("2006-01-02", operationstr3, loc)
			if err != nil {
				operatime3 = time.Now().AddDate(1, 0, 0)
				fmt.Println("------ hos :", operatime3, err)
			}
			operaoffset1 := t2.Sub(operatime1).Hours()
			difftime1 := operatime2.Sub(operatime1).Hours()
			difftime2 := operatime3.Sub(operatime2).Hours()
			fmt.Println("offset :", operaoffset1, difftime1, difftime2)

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
		}
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

				if operationstr3 != "" {
					operatime3, err = time.ParseInLocation("2006-01-02", operationstr3, loc)
					if err != nil {
						operatime3 = time.Now().AddDate(1, 0, 0)
						fmt.Println("------ hos :", operatime3, err)
					}
				}
				fmt.Println("success", operatime1, operatime2)

				operaoffset1 := t2.Sub(operatime1).Hours()
				difftime1 := operatime2.Sub(operatime1).Hours()
				difftime2 := operatime3.Sub(operatime2).Hours()
				fmt.Println("offset :", operaoffset1, difftime1)

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
