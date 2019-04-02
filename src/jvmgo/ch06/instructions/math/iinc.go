package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

//iinc 指令给局部变量表中的 int 变量增加常量值，局部变量表索引和常量值都由指令的操作数提供。

type IINC struct {
	Index uint
	Const int32
}

func (i *IINC) FetchOperands(reader *base.BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
	i.Const = int32(reader.ReadUint8())
}


func (i *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(i.Index)
	//加上常量值
	val += i.Const
	localVars.SetInt(i.Index, val)
}