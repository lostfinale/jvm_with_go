package heap

import "jvmgo/classfile"

//类的属性
type Field struct {
	ClassMember
	constValueIndex uint
	slotId uint //在属性列表中的序号
}



func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {

	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		//读取描述符、name、和访问权限标记
		fields[i].copyMemberInfo(cfField)
		//读取属性
		fields[i].copyAttributes(cfField)
	}
	return fields
}

//从字段属性表中读取 constValueIndex
func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}
func (self *Field) SlotId() uint {
	return self.slotId
}
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) Descriptor() string {
	return self.descriptor
}

