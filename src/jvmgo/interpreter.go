package main

import (
	"jvmgo/rtda"
	"fmt"
	"jvmgo/instructions/base"
	"jvmgo/instructions"
	"jvmgo/rtda/heap"
)

func interpret(method *heap.Method) {

	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())


}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		//fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		//fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		//panic(r)
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
