package golang

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

// struct
// 结构体的内存对齐，见上
func TestStruct(t *testing.T) {
	// 匿名结构体，_字段
	m := struct {
		_    int
		Name string
		Age  int
	}{Name: "song", Age: 18}
	fmt.Printf("%#v\n", m)
	println()

	// 在所有字段都支持==、!=时，struct可比较
	type Foo struct {
		Name string
		Age  int
		// Child []int
	}
	f1 := Foo{Name: "song", Age: 12}
	f2 := Foo{Name: "li", Age: 12}
	println(f1 == f2)
	println()

	// 空结构体的长度为0，作为数组元素也是0
	// 一个不需要占用实际内存的 堆上 的变量，都会指向runtime.zerobase，例如一个空结构体的切片
	b := struct{}{}
	bs := [100]struct{}{}
	println(unsafe.Sizeof(b))
	println(unsafe.Sizeof(bs))

	// 使用空结构体在管道中做通信
	c := make(chan struct{})
	go func() {
		<-time.After(3 * time.Second)
		c <- struct{}{}
	}()
	select {
	case <-c:
		println("get struct, exit")
	}
	println()

	// 结构体的组合
	// 匿名字段，只有类型，没有名称的字段，
	// 默认以类型名作为字段名，但可以直接引用匿名字段成员
	// 可以是任何类型、类型指针。
	type Bar struct {
		Height int `高度:"完美"` // Tag，字段标签
		Foo
		int     // 名称为int
		*string // 名称为string
		// string 类型及其指针字段名相同，不能同时包含
	}

	bb := Bar{Height: 12}
	v := reflect.ValueOf(bb)
	println(v.Type().Field(0).Tag)
}

// 方法, 面向对象的部分实践
func TestMethod(t *testing.T) {
	// 一个对象可以调用其方法，以及其字段的方法
	// 对象方法会覆盖匿名字段的方法，除非显示指定用该字段调用
	m := Manager{}
	m.toString()
	m.User.toString()
	println("————————————————————————")

	// 实例和指针的方法集不同，但是它们均可以调用所有方法，不论其接收者是实例还是指针
	m.toString2()
	(&m).toString()
	println("————————————————————————")

	// 方法集：
	// 实例：所有 receiver T 的方法
	// 指针：所有 receiver T, *T 的方法
	// 嵌入S，T包含所有 receiver S 方法
	// 嵌入*S，T包含所有 receiver S, *S 方法
	// 嵌入S或者*S，*T包含所有 receiver S, *S 方法
	ty := reflect.TypeOf(m)
	for i, n := 0, ty.NumMethod(); i < n; i++ {
		me := ty.Method(i)
		fmt.Println(me.Name, me.Type)
	}
}

// interface
//	type iface struct {
//		tab  *itab // 保存interface类型、对象类型、对象方法地址
//		data unsafe.Pointer // 实际对象指针
//	}
// 最常用于 对包外提供访问 预留拓展空间
// 根据实例的方法集判断对象是否实现接口
// 接口可组合，方法不能重名
func TestInterface(t *testing.T) {

	//  实例的方法集没有实现接口，指针实现了
	// m := Manager{}
	// multi1 := MultiString(m)
	// multi2 := MultiString(&m)

	// 空接口没有方法，所有被任何类型实现
	// 如果实现接口的类型支持，那么接口可比较
	var t1, t2 interface{}
	println(t1 == t2, t1 == nil)
	t1, t2 = 100, 100
	println(reflect.TypeOf(t1).String())
	println(t1 == t2)
	// t1, t2 = []int{}, []int{}
	// println(t1 == t2)

	// 接口组合
	var mm MMMMMultiString
	var m MultiString = mm
	println(m, "\n")

	// 匿名接口
	var tt interface {
		toString()
	} = User{}
	println(tt)
	println()

	// 把变量赋值给接口时，会发生复制
	// unaddressable的变量不可赋值

	// 两个方法集相同的接口，可以作比较。
	// 先比较类型，再比较方法。接口默认值是nil。

	// 接口变量的两部分都为nil, 接口才为nil。
	var a interface{} = nil
	var b interface{} = (*int)(nil) // b是有类型的
	println(a == nil, b == nil)
	println()

	// 接口的类型转换。
	// 接口和接口, 不使用ok模式会panic
	if x, ok := mm.(MultiString); ok {
		println(x)
	}
	// 接口和具体类型，同样可以ok模式，或者 switch a.(type) case int ...
	println(b.(*int))

	// 通过编译器检查是否实现某个接口
	// var x string
	// var _ MultiString = x // 提示错误，因为x并没有实现该接口
}

// 嵌入匿名变量，TB继承了该变量的所有方法，当然实际调用方法的仍然是匿名变量
type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

func TestInterfaceWrapper(t *testing.T) {
	// 隐式转换，因为TB实现了testing.TB的所有方法
	// 这样就跳过了私有方法，而在外部实现了testing.TB接口。
	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")
}
