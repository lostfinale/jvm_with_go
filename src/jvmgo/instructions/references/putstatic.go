package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)


//两个操作数。第一个操作数是 uint16 索引，来自字节码。通过这个索引可以从当前类的运行时常量池中找到一个字段符号引用，解析这个符号引用就可以知
//道要给类的哪个静态变量赋值。第二个操作数是要赋给静态变量的值，从操作数栈中弹出
type PUT_STATIC struct {
	base.Index16Instruction
}



func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()//当前方法
	currentClass := currentMethod.Class()//当前类
	cp := currentClass.ConstantPool()//当前常量池
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef) //字段符号应用
	field := fieldRef.ResolveField()//字段
	class := field.Class()

	if !field.IsStatic() {
		//只能处理静态字段
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.IsFinal() {
		//常量必须在字段所在的类初始化，并且必须在类初始化方法中赋值
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars() //类变量
	stack := frame.OperandStack() //操作数栈
	switch descriptor[0] {
		case 'Z', 'B', 'C', 'S', 'I':
			slots.SetInt(slotId, stack.PopInt())
		case 'F':
			slots.SetFloat(slotId, stack.PopFloat())
		case 'J':
			slots.SetLong(slotId, stack.PopLong())
		case 'D':
			slots.SetDouble(slotId, stack.PopDouble())
		case 'L', '[':
			slots.SetRef(slotId, stack.PopRef())
	}

}