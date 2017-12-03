package h_pkg

type CommonFunc struct{}

var commonFunc CommonFunc

func (c *CommonFunc) Merge(s ...[]interface{}) (slice []interface{}) {
	switch len(s) {
	case 0:
		break
	case 1:
		slice = s[0]
		break
	default:
		s1 := s[0]
		s2 := commonFunc.Merge(s[1:]...) //...将数组元素打散
		slice = make([]interface{}, len(s1)+len(s2))
		copy(slice, s1)
		copy(slice[len(s1):], s2)
		break
	}

	return
}

// x数组，i：x数组的起始下标
// y数组，j：y数组的起始下标
// 说明：x和y数组都是有序的数组
func merge(x []int, i int, y []int, j int) []int {
	xlen := len(x)              //x数组的长度
	ylen := len(y)              //y数组的长度
	z := make([]int, xlen+ylen) //创建一个大小为xlen+ylen的数组切片
	k := 0                      //数组切片z的下标
	for k !=  xlen + ylen {
		if x[i] < y[j] { //把小数放在数组切片z里
			z[k] = x[i]
			i++
		} else {
			z[k] = y[j]
			j++
		}
		k++
	}
	for i != xlen { //把x到xlen-1的数据也存入z中
		z[k] = x[i]
		k++
		i++
	}
	for j != ylen {
		z[k] = y[j]
		k++
		j++
	}
	return z
}
