package heap

import (
	"jvmgo/classpath"
	"fmt"
	"jvmgo/classfile"
)

//类加载器
type ClassLoader struct {
	cp *classpath.Classpath //classpath
	classMap map[string]*Class //已加载的类
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{cp :cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	return self.loadNonArrayClass(name)
}


//类的加载大致可以分为三个步骤：首先找到 class 文件并把数据读取到内存；
//然后解析class 文件，生成虚拟机可以使用的类数据，并放入方法区；最后进行链接。
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}


func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte)  *Class {
	//把 class 文件数据转换成 Class 结构体
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class

}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}



func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//todo 暂时忽略
}

func prepare(class *Class) {

	//计算成员变量的个数
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)

}

//计算成员变量的个数
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		//如果有父类，那么加上父类的成员属性的个数（父类在此之前，已经走过完整的loadClass流程）
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}

	class.instanceSlotCount = slotId
}

//计算静态变量的个数
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

//初始化静态变量
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)

	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			//初始化话静态常量
			//如果静态变量属于基本类型或 String 类型，有 final 修饰符，且它的值在编译期已知，
			//则该值存储在 class 文件常量池中
			initStaticFinalVar(class, field)
		}
	}

}


//从常量池中加载常量值，然后给静态变量赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)

		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}


