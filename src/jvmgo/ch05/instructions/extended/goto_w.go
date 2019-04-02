package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//goto_w 指令和 goto 指令的唯一区别就是索引从 2 字节变成了 4 字节

type GOTO_W struct {
	offset int
}

func (g *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	g.offset = int(reader.ReadInt32())
}

func (g *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.offset)
}