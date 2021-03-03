package main

import (
	"fmt"
	"os"
)

type bar interface {
	getVal() int
	updateVal(v int)
}

type foo struct {
	val int
}

func (f foo) getVal() int {
	return f.val
}

func (f *foo) updateVal(v int) {
	f.val = v
}

func main() {
	println(1)
	println(2)
	println(3)
	println(4)

	fmt.Println("数组***********************************")
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	fmt.Println(arr1, arr2)
	arr2[0] = 10002
	fmt.Println(arr1, arr2)

	fmt.Println("切片***********************************")
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	fmt.Println(slice1, slice2)
	slice2[0] = 10002
	fmt.Println(slice1, slice2)

	os.Exit(1)

	arr := make([]int, 5)
	for {
		arr = append(arr, )
	}
	//
	//// map只能为slice, map, channel分配内存，并返回一个初始化的值
	//
	//t1 := make(map[int]string)
	//// 第一种用法，即缺少长度的参数，只传类型
	//// 这种用法只能用在类型为map或chan的场景，例如make([]int)是会报错的
	//
	//t2 := make([]int, 2)
	//// 第二种用法，指定了长度，例如make([]int, 2)返回的是一个长度为2的slice
	//
	//t3 := make([]int, 2, 4)
	//// 第三种用法，第二参数指定的是切片的长度，第三个参数是用来指定预留的空间长度，
	//// 例如a := make([]int, 2, 4), 这里值得注意的是返回的切片a的总长度是4，
	//// 预留的意思并不是另外多出来4的长度，
	//// 其实是包含了前面2个已经切片的个数的。
	//// 所以举个例子当你这样用的时候 a := make([]int, 4, 2)，就会报语法错误。
	//
	//t3[0] = 1
	//t3 = append(t3, 1, 2, 7)
	//
	//println(t3)
	//println(cap(t3))
	//println(len(t3))
	//
	//println(t1, t2, t3)
	//
	//
	//f := make([]bar, 0, 100)
	//
	//for i := 0; i < 3; i++ {
	//	f = append(f, &foo{1})
	//}
	//
	//for _, fo := range f {
	//
	//	println(fo.getVal())
	//
	//	fo.updateVal(15)
	//
	//	println(fo.getVal())
	//	println("----------------------------")
	//
	//}
	//println("++++++++++++++++++++++++++++++++")
	//for _, fo := range f {
	//	println(fo.getVal())
	//}

}
