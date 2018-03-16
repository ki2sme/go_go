package ki2sme

import "fmt"

func Arraytest()  {

	array1 := [...]int{1,2,3}		//自动推到数组大小
	fmt.Printf("array1 %v \n", array1)

	array2 := [5]int{1:1,3:4} 		//初始化索引 为 1 和 3 的位置
	fmt.Printf("array1 %v \n", array2)

	//同样类型的数组是可以相互赋值的，不同类型的不行，会编译错误。那么什么是同样类型的数组呢？Go语言规定，必须是长度一样，并且每个元素的类型也一样的数组，才是同样类型的数组。
	//array1 = array2

	//指针数组
	//创建一个指针数组，并未第一个单元分配空间
	//但是第二个还未分配空间，不能使用
	arrayPtr := [2]*int{0:new(int)}
	//为第二个分配空间
	arrayPtr[1] = new(int)
	//赋值
	i := 1
	arrayPtr[0] = &i
	*arrayPtr[1] = 2

	fmt.Printf("array ptr : %v \n", arrayPtr)

	//声明一个数组的指针
	//不同大小的数组  指针不一样 比如 *[2]*int 必须为2
	//指向 指针数组 的 指针
	var ptr *[2]*int = &arrayPtr
	fmt.Printf("%p \n", ptr)
	fmt.Printf("%v \n", *ptr)

	//nil slice
	var nilSlice []int

	//空 slice 类型 Collections.emptyList()
	emptySlice := []int{}

	//使用一下防止报错
	nilSlice = emptySlice
	emptySlice = nilSlice

	//数组传值， slice 传指针
}