package heap

import "jvmgo/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}


func (self *FieldRef) ResolveField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}

	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.field = field
}

func lookupField(c *Class, name string,  descriptor string) *Field {

	//首先在 C 的字段中查
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	//在 C 的直接接口递归应用这个查找过程
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	//在 C 的超类中递归应用这个查找过程
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}