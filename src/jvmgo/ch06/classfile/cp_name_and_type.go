package classfile


/*
	CONSTANT_NameAndType_info {
		u1 tag;
		u2 name_index;
		u2 descriptor_index;
	}
 */


type ConstantNameAndTypeInfo struct {
	nameIndex uint16
	descriptorIndex uint16
}

func (ci *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	ci.nameIndex = reader.readUint16()
	ci.descriptorIndex = reader.readUint16()
}

