package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs int32
	matchOffsets []int32
}



func (l *LOOKUP_SWITCH)FetchOperands(reader *base.BytecodeReader) {

	//跳过 padding
	reader.SkipPadding()

	l.defaultOffset = reader.ReadInt32()
	l.npairs = reader.ReadInt32()
	l.matchOffsets = reader.ReadInt32s(l.npairs * 2)
}

//先从操作数栈中
//弹出一个 int 变量，然后用它查找 matchOffsets ，看是否能找到匹配的 key 。如果能，则按照 value 给出的偏
//移量跳转，否则按照 defaultOffset 跳转
func (l *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i:=int32(0); i < l.npairs*2; i+=2 {
		if l.matchOffsets[i] == key {
			offset := l.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			break
		}
	}
	base.Branch(frame, int(l.defaultOffset))
}