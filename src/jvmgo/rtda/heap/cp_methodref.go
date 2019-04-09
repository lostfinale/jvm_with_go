package heap

import "jvmgo/classfile"


//方法引用常量（非接口方法）
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodRefInfo) * MethodRef {
	ref := &MethodRef{}
	ref.cp =cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolveMethodRef()  {

	//找到调用者的class
	d := self.cp.class
	//找到方法所属的class
	c := self.ResolvedClass()

	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)

	if method == nil {
		panic("java.lng.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}


func lookupMethod(class *Class, name string, descriptor string) *Method{
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = LookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}

func LookupMethodInClass(class *Class, name string, descriptor string) *Method {
	for c:=class; c != nil; c = c.superClass {
		for _,method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func LookupMethodInInterfaces(ifaces []*Class, name string, descriptor string) *Method {
	for _, iface := range ifaces {
		for _,method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}

		method := LookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}









