package rtda


import "jvmgo/rtda/heap"

//存放局部变量的插槽，int字段存放整数，ref存放引用
type Slot struct {
	num int32
	ref *heap.Object
}