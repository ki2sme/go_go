package ki2sme

import "fmt"

func Maptest()  {

	//在Go Map中，如果我们获取一个不存在的键的值，也是可以的，返回的是值类型的零值

	//声明一个map
	var mapVar map[string]int

	//初始化
	mapVar = make(map[string]int)

	//获取一个key
	i := mapVar["no"]

	fmt.Printf("map value : %d \n", i)
}
