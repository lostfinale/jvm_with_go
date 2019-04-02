package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/loads"
	"jvmgo/ch05/rtda"
	"jvmgo/ch05/instructions/stores"
	"jvmgo/ch05/instructions/math"
)


//加载类指令、存储类指令、 ret 指令和 iinc 指令需要按索引访问局部变量表，索引以 uint8 的形式存在
//字节码中。对于大部分方法来说，局部变量表大小都不会超过 256 ，所以用一字节来表示索引就够了。
//但是如果有方法的局部变量表超过这限制呢？ Java 虚拟机规范定义了 wide 指令来扩展前述指令。


type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode{
	case 0x15:
		inst := &loads.ILOAD{}
		//扩展成了2字节，本身是1字节的
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
		break
	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x36:
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x37:
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x38:
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x39:
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x3a:
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		//iinc 指令有两个操作数，都需要扩展成 2 字节
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0xa9://ret
	panic("Unsupported opcode: 0xa9!")
	}
}

//wide 指令只是增加了索引宽度，并不改变子指令操作，所以其 Execute （）方法只要调用子指令的
//Execute （）方法即可
func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}