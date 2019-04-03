package classfile

/**
	CONSTANT_Fieldref_info {
		u1 tag;
		u2 class_index;
		u2 name_and_type_index;
	}
*/

//基类
type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}


//由于一模一样，所以属性
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct{
	ConstantMemberRefInfo
}
type ConstantInterfaceMethodRefInfo struct{
	ConstantMemberRefInfo
}

func (ci *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	ci.classIndex = reader.readUint16()
	ci.nameAndTypeIndex = reader.readUint16()
}

func (ci *ConstantMemberRefInfo) ClassName() string {
	return ci.cp.getClassName(ci.classIndex)
}

func (ci *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return ci.cp.getNameAndType(ci.nameAndTypeIndex)
}
