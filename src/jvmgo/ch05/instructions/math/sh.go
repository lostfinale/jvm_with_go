package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//int 左位移
type ISHL struct {
	base.NoOperandsInstruction
}

func (sh * ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	//int 变量只有 32 位，所以只取 v2 的前 5 个比特就足够表示位移位数
	//Go 语言位移操作符右侧必须是无符号整数
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

//int 右算术位移（有符号位移，高位用符号位填充）
type ISHR struct {
	base.NoOperandsInstruction
}


func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

//int 右逻辑位移（无符号位移，高位用0填充）
type IUSHR struct {
	base.NoOperandsInstruction
}

func (sh * IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	//Go 语言并没有 Java 语言中的 >>> 运算符，为了达到无符号位移的目的，
	//需要先把 v1 转成无符号整数，位移操作之后，再转回有符号整数
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)

}

//long 左位移
type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

//long 右算术位移（有符号位移，高位用符号位填充）
type LSHR struct {
	base.NoOperandsInstruction
}


func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

//long 右逻辑位移（无符号位移，高位用0填充）
type LUSHR struct {
	base.NoOperandsInstruction
}
func (sh * LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}

