package h_pkg

import (
	"time"
	"fmt"
)

func test(datestr string)  {

	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", datestr, time.Now().Location())
	t2, err2 := time.ParseInLocation("2006-01-02 15:04:05", datestr, loc)
	t3, _ := time.Parse("2006-01-02 15:04:05", datestr)
	if err != nil || err2 != nil {
		fmt.Println("error", "temp chart", err, err2)
	}
	fmt.Println(t)
	fmt.Println(t2)
	fmt.Println(t3)
}

func GetWeeks() []time.Time {
	t := time.Now()
	index := t.Weekday()
	fmt.Println(time.Now())
	loc, _ := time.LoadLocation("Local")

	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	var weeks []time.Time
	for i := 0; i < 7; i++ {
		t2 := t1.AddDate(0, 0, -int(index)+i)
		//fmt.Println("time:", t2.String())
		weeks = append(weeks, t2)
	}
	return weeks
}


