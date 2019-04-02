package main

import (
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/rtda"
	"fmt"
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/instructions"
)

func interpret(methodInfo *classfile.MemberInfo) {

	codeAttr := methodInfo.CodeAttribute()

	maxLoals := codeAttr.MaxLocals()

	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLoals,maxStack)

	thread.PushFrame(frame)

	defer catchErr(frame)

	loop(thread, bytecode)


}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(t *rtda.Thread, bytecode []byte){
	frame := t.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		t.SetPC(pc)
		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		//执行
		inst.Execute(frame)
	}
}
