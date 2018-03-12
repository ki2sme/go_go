package main

import (
	"fmt"
	"unsafe"
	"bytes"
	"strconv"
)

func main() {
	//常量
	const SUM int32 = 12

	fmt.Println("hello word")
	fmt.Println(Unknown, Female, Male)
	fmt.Println(unsafe.Sizeof(SUM))

	//变量
	i := 2

	//指针 指向 int类型
	var ptr *int
	//指针数组
	var ptrArray [1] *int
	ptrArray[0] = &i

	//指向指针的指针
	var pptr **int
	ptr = &i
	pptr = &ptr

	fmt.Printf("** pptr ： %x \n", pptr)
	fmt.Printf("** pptr value ： %d \n", **pptr)

	if ptr == nil{
		ptr = &i
	}

	fmt.Println(ptr)
	fmt.Println(*ptr)	// 表示取指针指向变量的值

	if i > 1 {
		println(i)
	}
	//声明一个数组
	numbers := [6]int{13,23,44}

	for i,x := range numbers {
		fmt.Printf("%d, %d \n", i, x)
	}

	nextNumber := getSquence()
	fmt.Printf("%d seq \n", nextNumber())
	fmt.Printf("%d seq \n", nextNumber())
	fmt.Printf("%d seq \n", nextNumber())
	fmt.Printf("%d seq \n", nextNumber())

	//多维数组
	multiArray := [3][4]int{
		{1,2,3,4} ,
		{1,2,3,4} ,
		{1,2,3,4} ,
	}
	fmt.Printf("%d .....\n", multiArray[1][1])

	book := Book{"Go", "Go", 123}
	book.title = "java"

	var ptrBook *Book = &book

	fmt.Printf("BOOK: %s\n", book.title)

	//一下效果一样
	fmt.Printf("BOOK: %s\n", (*ptrBook).title)
	fmt.Printf("BOOK: %s\n", ptrBook.title)

	//未定义大小的数组 即为 切片
	var slice []int
	fmt.Printf("sliece len : %d \n", len(slice))
	fmt.Printf("sliece cap : %d \n", cap(slice))

	slice2 := append(slice, 1, 2)
	fmt.Printf("sliece cap : %d \n", cap(slice))
	fmt.Printf("sliece cap : %d \n", cap(slice2))
	
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("sliece len : %d cap : %d \n", len(slice2), cap(slice2))

	slice3 := make([]int, 2)
	//类型，长度，容量
	slice3 = make([]int, 2, 8)

	fmt.Printf("slice %d %d \n", len(slice3), cap(slice3))

	fmt.Printf("slice v %v \n", slice3)

	//切片初始化
	slice4 := []int {1, 2, 3}
	fmt.Printf("slice 4 : %v \n", slice4)
	fmt.Printf("slice 4 : %v \n", slice4[1:2])

	// range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值
	for _, v := range slice4{
		fmt.Printf("range slice %d \n", v)
	}
	// 第一个string 为 key类型，第二个为 value 类型
	//var kvMap map[string]string  //只是声明，未初始化
	var kvMap = make(map[string]string)
	kvMap["k"] = "v"
	for k := range kvMap{
		fmt.Printf(" kv v is : %s \n", kvMap[k])
	}
	//存在
	_, exists := kvMap["k"]
	fmt.Printf(" exists : %t \n", exists)

	//删除
	delete(kvMap, "k")

	_, exists = kvMap["k"]
	fmt.Printf(" exists : %t \n", exists)

	//类型转换
	var someInt = 12
	var someInt2 = 7;

	mean := float64(someInt)/float64(someInt2)

	fmt.Printf("mean : %f \n", mean)

	b := bytes.Buffer{}
	b.WriteString("add ")
	b.WriteString(strconv.Itoa(someInt))

	fmt.Println(b.String())

	phone := new(IPhone)
	phone.call()
}

//闭包
func getSquence() func() int {

	i := 0
	return func() int {
		i += 1
		return i
	}
}
// 枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)
//结构体
type Book struct {
	title string
	author string
	id int
}

//接口
type Phone interface {
	call()
}

type IPhone struct {

}

func (iphone IPhone) call() {
	fmt.Println("iphone call")
}