package ki2sme

import "fmt"

//这种定义 称为函数
//函数的定义声明没有接收者，所以我们直接在go文件里，go包之下定义声明即可
func Functest()  {

	p := Person{age:12, name:"ss"}
	array := [...]int{1, 2}

	funcArray(array)
	funcStruct(p)

	fmt.Printf("func : %v \n", p)
	fmt.Printf("p ptr : %p \n", &p)
	fmt.Printf("func : %v \n", array)

	funcStructPtr(&p)
	fmt.Printf("func * : %v \n", p)
}

func funcArray(array [2]int)  {
	array[0] = 123
}

func funcStruct(p Person)  {
	p.name = "changed"
	fmt.Printf("p ptr in func : %p \n", &p)
	fmt.Printf("p v in func : %v \n", p)
}

func funcStructPtr(p *Person)  {
	p.name = "changed *"
}

type Person struct {
	age int
	name string
}

//定义一个方法
//方法在定义的时候，会在func和方法名之间增加一个参数，这个参数就是接收者，这样我们定义的这个方法就和接收者绑定在了一起，称之为这个接收者的方法。
//p Person即为接受者
//Go语言里有两种类型的接收者：值接收者和指针接收者
func (p Person) String() {
	fmt.Printf("p : %s, %d \n", p.name, p.age)
}

func (p *Person) update(name string) {
	p.name = name
}
//不管是使用值接收者，还是指针接收者，一定要搞清楚类型的本质：对类型进行操作的时候，是要改变当前值，还是要创建一个新值进行返回？这些就可以决定我们是采用值传递，还是指针传递。

//可变参数
func sum(args ...int) int {

	all := 0
	for _, v := range args{
		all += v
	}
	return all
}