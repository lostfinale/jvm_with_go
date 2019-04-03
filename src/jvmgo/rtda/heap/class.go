package heap

import (
	"jvmgo/classfile"
	"strings"
)

const (
	ACC_PUBLIC = 0x0001 // class field method
	ACC_PRIVATE = 0x0002 // field method
	ACC_PROTECTED = 0x0004 // field method
	ACC_STATIC = 0x0008 // field method
	ACC_FINAL = 0x0010 // class field method
	ACC_SUPER = 0x0020 // class
	ACC_SYNCHRONIZED = 0x0020 // method
	ACC_VOLATILE = 0x0040 // field
	ACC_BRIDGE = 0x0040 // method
	ACC_TRANSIENT = 0x0080 // field
	ACC_VARARGS = 0x0080 // method
	ACC_NATIVE = 0x0100 // method
	ACC_INTERFACE = 0x0200 // class
	ACC_ABSTRACT = 0x0400 // class method
	ACC_STRICT = 0x0800 // method
	ACC_SYNTHETIC = 0x1000 // class field method
	ACC_ANNOTATION = 0x2000 // class
	ACC_ENUM = 0x4000 // class field
)

//代表java的class对象
type Class struct {
	//类的访问标志，总共 16 比特。字段和方法也有访问标志，但具体标志位的含义可能有所不同
	accessFlags uint16
	name string
	superClassName string
	interfaceNames []string
	constantPool *ConstantPool //类的运行时常量池
	fields []*Field
	methods []*Method
	loader *ClassLoader //类加载器指针
	superClass *Class
	interfaces []*Class
	instanceSlotCount uint //实例变量占据的空间大小
	staticSlotCount uint //类变量占据的空间大小

	staticVars Slots //静态变量列表
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceName()
	class.constantPool = newConstantPool(class, cf.ConstantPool())//todo
	class.fields = newFields(class, cf.Fields())//todo
	class.methods = newMethods(class, cf.Methods())//todo
	return class
}


func (self *Class) NewObject() *Object {
	return newObject(self)
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}


func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

