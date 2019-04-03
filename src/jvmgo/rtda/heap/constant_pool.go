package heap

import (
	"fmt"
	"jvmgo/classfile"
)

//运行时常量
type Constant interface {}

//运行时常量池
type ConstantPool struct {
	class *Class
	consts []Constant
}


func (self *ConstantPool) GetConstant(index uint) Constant {
	if c:= self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants a5 index %d", index))
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool{
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value() // float32
		case *classfile.ConstantLongInfo:

			intInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = intInfo.Value()
			// long 或 double 型常量占据两个位置
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value() // float64
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String() // string

			//以下开始是符号常量，需要用对象来表示，以上是字面量，所以直接可以存储
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.ConstantMethodRefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		}
	}
	return rtCp

}