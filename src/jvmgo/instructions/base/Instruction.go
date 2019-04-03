package base

import "jvmgo/rtda"

/*
	JVM规范2.11定义
	do {
		atomically calculate pc and fetch opcode at pc;
		if (operands) fetch operands;
		execute the action for the opcode;
	} while (there is more to do);



	go中伪代码
	for {
		pc := calculatePC()
		opcode := bytecode[pc]
		inst := createInst(opcode)
		inst.fetchOperands(bytecode)
		inst.execute()
	}
*/
type Instruction interface {
	//从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	//执行指令逻辑
	Execute(frame * rtda.Frame)
}


//NoOperandsInstruction 表示没有操作数的指令，所以没有定义任何字段
type NoOperandsInstruction struct {

}

func (ins *NoOperandsInstruction) FetchOperands(reader * BytecodeReader) {
	// noting to to
}

//跳转指令 Offset字段存放跳转偏移量
type BranchInstruction struct {
	Offset int
}

//读取操作数
func (ins *BranchInstruction) FetchOperands(reader * BytecodeReader) {
	ins.Offset = int(reader.ReadInt16())
}

//存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出。
//把这类指令抽象成Index8Instruction 结构体，用 Index 字段表示局部变量表索引
type Index8Instruction struct {
	Index uint
}

//读取一个 int8 整数，转成 uint 后赋给 Index 字段
func (ins *Index8Instruction) FetchOperands(reader * BytecodeReader) {
	ins.Index = uint(reader.ReadUint8())
}
//有一些指令需要访问运行时常量池，常量池索引由两字节操作数给出。
//把这类指令抽象成Index16Instruction 结构体，用 Index 字段表示常量池索引。
type Index16Instruction struct {
	Index uint
}

func (ins *Index16Instruction) FetchOperands(reader * BytecodeReader) {
	ins.Index = uint(reader.ReadUint16())
}