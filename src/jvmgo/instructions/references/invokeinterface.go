package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

//在字节码中， invokeinterface 指令的操作码后面跟着 4 字
//节而非 2 字节。前两字节的含义和其他指令相同，是个 uint16 运行时常量池索引。第 3 字节的值是给方法传
//递参数需要的 slot 数，其含义和给 Method 结构体定义的 argSlotCount 字段相同。正如我们所知，这个数是
//可以根据方法描述符计算出来的，它的存在仅仅是因为历史原因。第 4 字节是留给 Oracle 的某些 Java 虚拟
//机实现用的，它的值必须是 0 。该字节的存在是为了保证 Java 虚拟机可以向后兼容。
func (self *INVOKE_INTERFACE) FetchOperands(reader * base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8()
	reader.ReadUint8()
}

func(self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {

	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}


	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}