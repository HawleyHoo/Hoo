package h_tool

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	opt1 := WriteA(int64(1))
	opt2 := WriteB("test")
	//opt3 := WriteC(make(map[int]string,0))

	//c1 := make(map[string]string, 2)
	c1 := map[string]string{"c": "hehe", "d": "haha"}
	opt3 := WriteC(c1)

	op := NewOption(opt1, opt2, opt3)

	fmt.Println(op.a, op.b, op.c)
}
