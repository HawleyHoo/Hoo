package main

import (
	"time"
	"fit"
	"strconv"
	"fmt"
	"encoding/json"
	"strings"
)

type FirstNursingRecordController struct {
	fit.Controller
}

/*添加/修改首次护理单数据，传nrlid是编辑修改，无nrlid则是新增*/
func (c FirstNursingRecordController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	// 病人ID
	VAA01 := r.FormValue("pid")
	// 护士ID
	BCE01A := r.FormValue("nid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	NRL38 := r.FormValue("datetime")
	// 科室ID
	BCK01 := r.FormValue("did")
	// 文书ID（修改时需要）
	recordId := r.FormValue("nrlid")

	if VAA01 == "" || BCE01A == "" || BCE03A == "" || NRL38 == "" || BCK01 == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	_, err_BCK01 := strconv.Atoi(BCK01)
	if err_BCK01 != nil {
		c.RenderingJsonAutomatically(2, "参数错误  department_id: ")
		return
	}
	VAA01_i, err_VAA01 := strconv.Atoi(VAA01)
	if err_VAA01 != nil {
		c.RenderingJsonAutomatically(2, "参数错误 VAA01: ")
		return
	}
	_, err_BCE01A := strconv.Atoi(BCE01A)
	if err_BCE01A != nil {
		c.RenderingJsonAutomatically(2, "参数错误 BCE01A: ")
		return
	}
	_, err_NRL38 := time.Parse("2006-01-02 15:04:05", NRL38)
	if err_NRL38 != nil {
		c.RenderingJsonAutomatically(2, "参数错误 NRL38")
		return
	}
	NRL02 := r.FormValue("NRL02")
	NRL03 := r.FormValue("NRL03")
	NRL04 := r.FormValue("NRL04")
	NRL05 := r.FormValue("NRL05")
	NRL06 := r.FormValue("NRL06")
	NRL06A := r.FormValue("NRL06A")
	NRL07 := r.FormValue("NRL07")
	NRL07A := r.FormValue("NRL07A")
	NRL08 := r.FormValue("NRL08")
	NRL09 := r.FormValue("NRL09")

	NRL10 := r.FormValue("NRL10")
	NRL11 := r.FormValue("NRL11")
	NRL12 := r.FormValue("NRL12")
	NRL12A := r.FormValue("NRL12A")
	NRL13 := r.FormValue("NRL13")
	NRL14 := r.FormValue("NRL14")
	NRL14A := r.FormValue("NRL14A")
	NRL15 := r.FormValue("NRL15")
	NRL16 := r.FormValue("NRL16")
	NRL16A := r.FormValue("NRL16A")
	NRL17 := r.FormValue("NRL17")
	NRL17A := r.FormValue("NRL17A")
	NRL18 := r.FormValue("NRL18")
	NRL19 := r.FormValue("NRL19")
	LimbsId := r.FormValue("LimbsId")

	NRL20 := r.FormValue("NRL20")
	NRL20A := r.FormValue("NRL20A")
	NRL21 := r.FormValue("NRL21")
	NRL22 := r.FormValue("NRL22")
	NRL23 := r.FormValue("NRL23")
	NRL24 := r.FormValue("NRL24")
	NRL25 := r.FormValue("NRL25")
	NRL26 := r.FormValue("NRL26")
	NRL27 := r.FormValue("NRL27")
	NRL28 := r.FormValue("NRL28")
	NRL29 := r.FormValue("NRL29")

	NRL30 := r.FormValue("NRL30")
	NRL30A := r.FormValue("NRL30A")
	NRL31 := r.FormValue("NRL31")
	NRL32 := r.FormValue("NRL32")
	NRL33 := r.FormValue("NRL33")
	NRL34 := r.FormValue("NRL34")
	NRL35 := r.FormValue("NRL35")
	NRL36 := r.FormValue("NRL36")
	NRL37 := r.FormValue("NRL37")

	nrl2 := model.NRL2{
		VAA01:   VAA01_i,
		NRL02:   NRL02,
		NRL03:   NRL03,
		NRL04:   NRL04,
		NRL05:   NRL05,
		NRL06:   NRL06,
		NRL06A:  NRL06A,
		NRL07:   NRL07,
		NRL07A:  NRL07A,
		NRL08:   NRL08,
		NRL09:   NRL09,
		NRL10:   NRL10,
		NRL11:   NRL11,
		NRL12:   NRL12,
		NRL12A:  NRL12A,
		NRL13:   NRL13,
		NRL14:   NRL14,
		NRL14A:  NRL14A,
		NRL15:   NRL15,
		NRL16:   NRL16,
		NRL16A:  NRL16A,
		NRL17:   NRL17,
		NRL17A:  NRL17A,
		NRL18:   NRL18,
		NRL19:   NRL19,
		LimbsId: LimbsId,
		NRL20:   NRL20,
		NRL20A:  NRL20A,
		NRL21:   NRL21,
		NRL22:   NRL22,
		NRL23:   NRL23,
		NRL24:   NRL24,
		NRL25:   NRL25,
		NRL26:   NRL26,
		NRL27:   NRL27,
		NRL28:   NRL28,
		NRL29:   NRL29,
		NRL30:   NRL30,
		NRL30A:  NRL30A,
		NRL31:   NRL31,
		NRL32:   NRL32,
		NRL33:   NRL33,
		NRL34:   NRL34,
		NRL35:   NRL35,
		NRL36:   NRL36,
		NRL37:   NRL37,
		NRL38:   NRL38,
		BCE01A:  BCE01A,
		BCE03A:  BCE03A,
	}
	if recordId == "" {
		isExist, _ := model.IsExistNRL2(VAA01)
		if isExist == true {
			c.RenderingJsonAutomatically(4, "已存在该病人的首次护理单，请勿重复添加")
			return
		}

		// 插入数据库并返回数据对应的ID
		newNrlId, err := nrl2.InsertToDatabase()
		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
			return
		}

		// 在护理单记录总表中插入一条记录
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		nr := model.NursingRecords{
			Updated:     timeNow,
			NursType:    2,
			NursingId:   BCE01A,
			NursingName: BCE03A,
			ClassId:     BCK01,
			PatientId:   VAA01,
			RecordId:    int64(newNrlId),
			Comment:     "新增",
		}
		_, err = model.InsertNRecords(nr)

		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
		} else {
			c.RenderingJsonAutomatically(0, "添加成功")
		}
	} else {
		recordId_i, err_nrlId := strconv.ParseInt(recordId, 10, 64)
		if err_nrlId != nil {
			c.RenderingJsonAutomatically(2, "参数错误 nrl_id")
			return
		}
		// 修改数据库
		err := nrl2.UpdateRecordDatas(recordId_i)

		// 在护理单记录总表更新修改记录
		_, err = model.UpadteNRecords(recordId_i, NRL38)

		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
		} else {
			c.RenderingJsonAutomatically(0, "修改成功")
		}
	}
}

func (c *FirstNursingRecordController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make(map[string]interface{}, 0))
}

func (c *FirstNursingRecordController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}

/*加载首次护理单数据的Html网页，type=1：插入添加，type=2：修改编辑*/
func (c FirstNursingRecordController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	typeStr := r.FormValue("type")
	// 病人ID
	patient_id := r.FormValue("pid")
	// 护士ID
	nurse_id := r.FormValue("nid")
	// 文书ID（修改时需要）
	nrl_id := r.FormValue("nrlid")

	if "" == typeStr || patient_id == "" || nurse_id == "" {
		fmt.Fprintln(w, "参数不完整！")
		return
	} else if typeStr == "2" {
		if nrl_id == "" {
			fmt.Fprintln(w, "参数不完整！")
			return
		}
	}

	// 根据病人ID查询病人基本信息
	patintInfos, err_pt := model.GetPatientInfo(patient_id)
	if err_pt != nil {
		fmt.Fprintln(w, "Database "+err_pt.Error())
		return
	}
	if length := len(patintInfos); length == 0 || patintInfos == nil {
		fmt.Fprintln(w, "病人ID错误！"+patient_id)
		return
	}
	patient := patintInfos[0]
	// 住院时间
	VAA73 := time.Time(patient.VAA73).Format("2006-01-02 15:04")

	_, err_nur := strconv.Atoi(nurse_id)
	if err_nur != nil {
		fmt.Fprintln(w, "护士ID错误！"+nurse_id)
		return
	}
	// 查询护士的信息
	nurse, _ := model.FetchAccountWithUid(nurse_id)

	switch typeStr {
	case "1":
		//插入
		// 查询是否存在该病人的首次护理记录单
		//if isExist, err_exist := model.IsExistNRL(patient_id); isExist == true || err_exist != nil {
		//	fmt.Fprintln(w, "<h1>已存在该病人的首次护理单,无需再次添加</h1>")
		//	return
		//}
		NRL := model.NRL2{}

		c.Data = fit.Data{
			"Patient":      patient,        // 病人的基本信息
			"VAA73":        VAA73,          // 病人的入院时间
			"NurseId":      nurse_id,       // 护士ID
			"NurseName":    nurse.Username, // 护士姓名
			"Type":         typeStr,        // 类型 1
			"DepartmentId": patient.BCK01B, // 科室ID
			"NRL":          NRL,            // 护理单内容
			"NRL06A":       "",             // 过敏源的index
			"NRL06B":       "",             // 过敏源的补充内容
			"NRL18A":       "",             // n次/天
			"NRL18B":       "",             // 1次/m天
			"NRL38A":       "",             // 录入护理单的年月日
			"NRL38B":       "",             // 录入护理单的时分
		}
		c.LoadView(w, "v_nrl2_update.html")
	case "2": //更新

		// 查询护理单内容
		NRL, err_db := model.QueryNRL2(nrl_id)
		if err_db != nil {
			fmt.Fprintln(w, "无法查询到相应护理单记录"+err_db.Error())
			return
		}

		// 过敏史-过敏源-NRL06-Json字符串，key：对应index，value：对应内容
		NRL06A := ""
		NRL06B := ""
		if NRL.NRL06 == "2" {
			by := []byte(NRL.NRL06A)
			anyObj := make(map[string]string, 0)
			json.Unmarshal(by, &anyObj)
			for k, v := range anyObj {
				NRL06A = k
				NRL06B = v
			}
		}

		// 排便次数- "n,m"，n,m代表2个空的数值，即n次/天，1次/m天
		NRL18A := ""
		NRL18B := ""
		if NRL.NRL18 != "" {
			slice := strings.Split(NRL.NRL18, ",")
			length := len(slice)
			if length == 1 {
				NRL18A = slice[0]
			} else if length == 2 {
				NRL18A = slice[0]
				NRL18B = slice[1]
			}
		}
		// 拆分护理单录入时间
		tempTime, _ := time.ParseInLocation("2006-01-02 15:04:05", NRL.NRL38, time.Local)
		NRL38A := tempTime.Format("2006-01-02")
		NRL38B := tempTime.Format("15:04")

		c.Data = fit.Data{
			"Patient":      patient,        // 病人的基本信息
			"VAA73":        VAA73,          // 病人的入院时间
			"NurseId":      nurse_id,       // 护士ID
			"NurseName":    nurse.Username, // 护士姓名
			"Type":         typeStr,        // 类型 1
			"DepartmentId": patient.BCK01B, // 科室ID
			"NRL":          NRL,            // 护理单内容
			"NRLID":        nrl_id,         // 护理单ID
			"NRL06A":       NRL06A,         // 过敏源的index
			"NRL06B":       NRL06B,         // 过敏源的补充内容
			"NRL18A":       NRL18A,         // n次/天
			"NRL18B":       NRL18B,         // 1次/m天
			"NRL38A":       NRL38A,         // 录入护理单的年月日
			"NRL38B":       NRL38B,         // 录入护理单的时分
		}
		c.LoadView(w, "v_nrl2_update.html")
	default:
		fmt.Fprintln(w, "参数错误！ type")
		return
	}

}

type QueryFirstNursingRecordController struct {
	fit.Controller
}

func (c QueryFirstNursingRecordController) Exist(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	pid := r.FormValue("pid")
	if pid == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	patient_id, err_p := strconv.Atoi(pid)
	if err_p != nil || patient_id == 0 {
		c.RenderingJsonAutomatically(2, "参数错误 pid")
		return
	}
	isExist, err_exs := model.IsExistNRL2(pid)
	if err_exs != nil {
		c.RenderingJsonAutomatically(3, "Database "+err_exs.Error())
		return
	}

	if isExist == true {
		c.RenderingJsonAutomatically(0, "已存在该病人的首次护理单")
	} else {
		c.RenderingJsonAutomatically(4, "不存在该病人的首次护理单")
	}
}

func (c *QueryFirstNursingRecordController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make(map[string]interface{}, 0))
}

func (c *QueryFirstNursingRecordController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}

/*查看病人的首次护理单数据*/
func (c QueryFirstNursingRecordController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl_id := r.FormValue("nrlid")
	patient_id := r.FormValue("pid")
	if "" == nrl_id || "" == patient_id {
		fmt.Fprintln(w, "参数不完整！")
		return
	}
	// 查询相应护理单数据
	NRL, err_db := model.QueryNRL2(nrl_id)

	if err_db != nil {
		fmt.Fprintln(w, "无法查询到相应护理单记录"+err_db.Error())
	}
	// 查询对应病人信息
	patients, err_pt := model.GetPatientInfo(patient_id)
	if length := len(patients); length == 0 {
		fmt.Fprintln(w, "无法查询到相关病人的信息")
		return
	} else if err_pt != nil {
		fmt.Fprintln(w, "无法查询到相关病人的信息，"+err_pt.Error())
		return
	}
	patient := patients[0]

	// 过敏史-过敏源-NRL06-Json字符串，key：对应index，value：对应内容
	NRL06A := ""
	NRL06B := ""
	if NRL.NRL06 == "2" {
		by := []byte(NRL.NRL06A)
		anyObj := make(map[string]string, 0)
		json.Unmarshal(by, &anyObj)
		for k, v := range anyObj {
			NRL06A = k
			NRL06B = v
		}
	}

	// 排便次数- "n,m"，n,m代表2个空的数值，即n次/天，1次/m天
	NRL18A := ""
	NRL18B := ""
	if NRL.NRL18 != "" {
		slice := strings.Split(NRL.NRL18, ",")
		length := len(slice)
		if length == 1 {
			NRL18A = slice[0]
		} else if length == 2 {
			NRL18A = slice[0]
			NRL18B = slice[1]
		}
	}

	// 入院时间
	VAA73 := time.Time(patient.VAA73).Format("2006-01-02 15:04")
	// 拆分护理单录入时间
	tempTime, _ := time.ParseInLocation("2006-01-02 15:04:05", NRL.NRL38, time.Local)
	NRL38A := tempTime.Format("2006-01-02")
	NRL38B := tempTime.Format("15:04")

	c.Data = fit.Data{
		"Patient": patient, // 病人数据
		"VAA73":   VAA73,   // 入院时间
		"NRL":     NRL,     // 护理单数据
		"NRL06A":  NRL06A,  // 过敏源的index
		"NRL06B":  NRL06B,  // 过敏源的补充内容
		"NRL18A":  NRL18A,  // n次/天
		"NRL18B":  NRL18B,  // 1次/m天
		"NRL38A":  NRL38A,  // 录入护理单的年月日
		"NRL38B":  NRL38B,  // 录入护理单的时分
	}

	c.LoadView(w, "v_nrl2.html")
}
