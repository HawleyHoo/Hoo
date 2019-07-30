package main

// 判断一个点是不是在多边形内
// 从这个点随便发出一条射线，看这条射线与多边形相交的次数，如果为奇数则点在多边形内，为偶数则在多边形外

func pnpoly(nvert, x0, y0 int, xlist, ylist []int) {
	for i, j := 0, nvert-1; i < nvert; i++ {
		if ((ylist[i] > y0) != (ylist[j] > y0)) &&
			(x0 < (xlist[j]-xlist[i])*(y0-ylist[i])/(ylist[j]-ylist[i])+xlist[i]) {

		}
	}
}
