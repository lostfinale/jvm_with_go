package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)


/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/


//Java 语言中的 switch-case 语句有两种实现方式：
// 如果 case 值可以编码成一个索引表，则实现成tableswitch 指令；
// 否则实现成 lookupswitch 指令。

//Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32 //defaultOffset 对应默认情况下执行跳转所需的字节码偏移量
	low int32 // low 和 high 记录 case 的取值范围
	high int32
	jumpOffsets []int32//jumpOffsets 是一个索引表，里面存放 high-low+1 个 int 值，对应各种 case 情况下，执行跳转所需的字节码偏移量
}

func (t * TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// tableswitch 指令操作码的后面有 0~3 字节的 padding ，
	// 以保证 defaultOffset 在字节码中的地址是 4 的倍数
	reader.SkipPadding()
	t.defaultOffset = reader.ReadInt32()
	t.low = reader.ReadInt32()
	t.high = reader.ReadInt32()
	jumpOffsetsCount := t.high - t.low + 1
	t.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

//先从操作数栈中弹出一个 int 变量，然后看它是否在 low 和 high 给定的范围之内。
//如果在，则从 jumpOffsets 表中查出偏移量进行跳转，否则按照 defaultOffset 跳转
func (t * TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= t.low && index <= t.high {
		offset = int(t.jumpOffsets[index-t.low])
	} else {
		offset = int(t.defaultOffset)
	}
	base.Branch(frame, offset)
}