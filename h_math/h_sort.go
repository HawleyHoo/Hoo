package main

import (
	"fmt"
)

func main()  {
	arr := []int{23,45,16,67,89,32,96,64,10,78,23,45,16,67,89,32,96,64,10,78}
	//arr := []int{23,45,16,67,89,32,96,64,10,78}
	//arr := []int{9,8,7,6,5}
	//arr := []int{1,2,3,4,5,6}
	fmt.Println("org", arr)
	//fmt.Println(quickSort(arr))
	fmt.Println(hooQuickSort2(arr, 0, len(arr) - 1))
	fmt.Println(bubbleSort(arr))
}

// 堆排序
func heapSort(arr []int) []int  {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)

	for i := arrLen - 1; i >= 0 ; i--  {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr

}

func buildMaxHeap(arr []int, arrLen int)  {
	for i := arrLen / 2; i >= 0; i-- {
		fmt.Println(arr, i, arrLen)
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i , arrLen int)  {
	left := 2 * i + 1
	right := 2 * i + 2
	largest := i
	//fmt.Println(left, right, largest)
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)

	}
}

func swap(arr []int, i,j int)  {
	//fmt.Println(i, j)
	arr[i], arr[j] = arr[j], arr[i]
}

// hoo
func buildh(arr []int, i, l int)  {
	left := 2 * i + 1
	right := 2 * i + 2
	largest := i
	if left < l && arr[left] > arr[largest] {
		largest = left
	}
	if right < l && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		buildh(arr, largest, l)
	}
}

func heapsort(arr []int) []int {
	l := len(arr)
	// 初始化大顶堆
	for i := l / 2; i >= 0;i-- {
		//fmt.Println("i:", i,l)
		buildh(arr, i, l)
	}

	fmt.Println("init:", arr)
	for i := l - 1; i >= 0;i--  {
		swap(arr, 0, i)
		l -= 1
		buildh(arr, 0, l)
	}
	return arr
}

// 冒泡排序
func bubbleSort(arr []int) []int  {
	length := len(arr)
	for i := 0; i < length; i++  {
		for j := 0; j < length - 1 - i; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
	return arr
}

// 选择排序
func selectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length - 1; i++  {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
	return arr
}


// 插入排序
func insertionSort(arr []int) []int  {
	for i := range arr {
		preIndex := i - 1
		cur := arr[i]
		for preIndex >= 0 && arr[preIndex] > cur  {
			arr[preIndex + 1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex + 1] = cur
	}
	return arr
}

func insertionSort2(arr []int) []int  {
	l := len(arr)
	for i:= 1; i < l; i++ {
		temp := arr[i]
		j := i
		for j > 0 && arr[j - 1] > temp  {
			arr[j] = arr[j - 1]
			j--
		}
		arr[j] = temp
	}
	return arr
}



// 希尔排序
func shellSort(arr []int) []int  {
	l := len(arr)
	gap := 1
	for gap < l / 3  {
		gap = gap * 3 + 1
	}
	//fmt.Println("gap", gap)
	for gap > 0 {
		for i := gap; i < l ; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp  {
				arr[j + gap] = arr[j]
				j -= gap
			}
			arr[j + gap] = temp
		}
		fmt.Println("gap", gap)
		gap = gap / 3
	}

	return arr
}

// 归并排序
func mergeSort(arr []int) ([]int)  {
	l := len(arr)
	if l < 2 {
		return arr
	}

	middle := l / 2
	left := arr[0:middle]
	right := arr[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (res []int) {
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		res = append(res, left[0])
		left = left[1:]
	}
	for len(right) != 0  {
		res = append(res, right[0])
		right = right[1:]
	}
	fmt.Println("res", res)
	return res
}

// 快速排序
func quickSort(arr []int) []int  {
	return _quicksort(arr, 0, len(arr) - 1)
}

func _quicksort(arr []int, left, right int) []int  {
	if left < right {
		partitionIndex := partition(arr, left, right)
		fmt.Println("partition", partitionIndex, arr)
		_quicksort(arr, left, partitionIndex - 1)
		_quicksort(arr, partitionIndex + 1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int  {
	//pivot := left
	index := left + 1

	for i := index;i <= right ;i++  {
		if arr[i] < arr[left] {
			swap(arr, i, index)
			index++
		}
	}
	swap(arr, left, index - 1)
	return index - 1
}

func hooQuickSort(arr []int, left, right int) ([]int)  {
	i := left
	j := right
	temp := arr[left]
	for i < j {
		for arr[j] >= temp && j > i  {
			j--
		}

		if j > i {
			arr[i] = arr[j]
			i++

			for arr[i] <= temp && j > i  {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}


	}

	arr[i] = temp


	if left < i - 1 {
		hooQuickSort(arr, left, i - 1)
	}
	if j + 1 < right {
		hooQuickSort(arr,j + 1, right)
	}

	return arr
}


func hooQuickSort2(arr []int, left, right int) ([]int)  {
	i := left
	j := right
	temp := arr[left]
	for i < j {
		for arr[j] >= temp && j > i  {
			j--
		}

		for arr[i] <= temp && j > i {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[left] = arr[i]
	arr[i] = temp


	if left < i - 1 {
		hooQuickSort2(arr, left, i - 1)
	}
	if j + 1 < right {
		hooQuickSort2(arr,j + 1, right)
	}

	return arr
}

