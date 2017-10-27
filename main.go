package main

import (
	//"Hoo/h_math"
	//"fmt"
	//"math"
	//"Hoo/h_pkg"
	"time"
	"fmt"
)

func main() {
	/*y0 := math.Sqrt(2)
	y1 := h_math.SqrtByNewton(2)
	y2 := h_math.SqrtNewton(2)
	y3 := h_math.InvSqrt(2)

	fmt.Println(y0)
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)


	h_pkg.Interfacetest()*/
	//h_pkg.ReflectTest()
	//h_pkg.ReflectTest2()
	//h_pkg.ReflectTest3()
	//h_pkg.ReflectTest5()
	timetest()

}




func timetest() {
	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.ISOWeek())
	fmt.Println(t.Weekday())

	t1, err := time.ParseInLocation("2006-01-02 15:04", "2017-10-10 00:00", t.Location() )


	fmt.Println(time.Since(t1).Hours() / 24, time.Since(t1))
	fmt.Println(int(time.Since(t1).Hours() / 24) / 7)
	fmt.Println(int(t.Sub(t1).Hours() / 24) / 7, int(t.Sub(t1).Hours() / 24) % 7)
	fmt.Println("t1:", t1.String())
	if err != nil {
		fmt.Println(err.Error())
	}
	index := t.Weekday()

	weeknum := int(time.Since(t1).Hours() / 24) / 7
	//weekoffset := int(time.Since(t) / 24) / 7
	fmt.Println("weeknum", weeknum)
	t3 := time.Date(t1.Year(), t1.Month(), t1.Day() + 7 * weeknum, 0, 0, 0, 0, t.Location())
	var weeks []time.Time
	for i := 0; i < 7; i++ {
		t2 := t3.AddDate(0, 0, -int(index)+i)
		fmt.Println("time:", t2.String())
		weeks = append(weeks, t2)
	}
	fmt.Println(weeks)
}
