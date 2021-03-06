package main

import (
	"jvmgo/rtda"
	"fmt"
	"jvmgo/instructions/base"
	"jvmgo/instructions"
)

func interpret(thread *rtda.Thread, logInst bool) {
	defer catchErr(thread)
	loop(thread, logInst)

}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func loop(t *rtda.Thread,  logInst bool){

	reader := &base.BytecodeReader{}
	for {
		frame := t.CurrentFrame()
		pc := frame.NextPC()
		t.SetPC(pc)
		//decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if logInst {
			logInstruction(frame, inst)
		}

		//执行
		inst.Execute(frame)
		if t.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

