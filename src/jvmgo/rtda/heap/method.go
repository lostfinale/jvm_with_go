package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxStack uint //操作数栈大小   值是由 Java 编译器计算好的
	maxLocals uint //局部变量表大小 值是由 Java 编译器计算好的
	code []byte  // code 字段存放方法字节码
	argSlotCount uint //
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		//提取操作数栈大小，局部变量表大小，方法字节码大小
		methods[i].copyAttributes(cfMethod)
		//计算参数个数
		methods[i].calcArgSlotCount()
	}
	return methods
}


//从 method_info 结构中提取属性
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeArr := cfMethod.CodeAttribute(); codeArr != nil {
		self.maxStack = codeArr.MaxStack()
		self.maxLocals = codeArr.MaxLocals()
		self.code = codeArr.Code()
	}
}

func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)

}

func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}